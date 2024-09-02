package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tomvodi/limepipes/internal/config"
	"github.com/tomvodi/limepipes/internal/interfaces"
	"github.com/tomvodi/limepipes/internal/pluginloader"
	"github.com/tomvodi/limepipes/internal/utils"
	"os"
	"path/filepath"
)

const DefaultOutputDir = "./parser_success"

var (
	OutputDir string
)

// parserCmd represents the parser command
var parserCmd = &cobra.Command{
	Use:   "parser [paths...]",
	Short: "Tests parsing given bww files",
	Long: `A bww parser testing command. Parses bww files and moves successful parses into an output directory.

When given directory paths, it will import all files of that directory. It will also include 
subdirectories when given the recursive flag.
If a given file that has an extension which is not in the import-file-types, it will be ignored.`,
	RunE: func(_ *cobra.Command, args []string) error {
		utils.SetupConsoleLogger()
		fs := afero.NewOsFs()

		err := checkForInvalidImportTypes()
		if err != nil {
			return err
		}

		cfg, err := config.Init()
		if err != nil {
			return fmt.Errorf("failed init configuration: %s", err.Error())
		}

		// TODO: Load plugins from config
		LoadPlugins := []string{
			"bww",
		}

		var pluginProcHandler interfaces.PluginProcessHandler = pluginloader.NewProcessHandler(LoadPlugins)
		var pluginLoader interfaces.PluginLoader = pluginloader.NewPluginLoader(
			fs,
			pluginProcHandler,
			LoadPlugins,
		)
		err = pluginLoader.LoadPluginsFromDir(cfg.PluginsDirectoryPath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed loading plugins")
		}
		defer func(pluginLoader interfaces.PluginLoader) {
			err := pluginLoader.UnloadPlugins()
			if err != nil {
				log.Fatal().Err(err).Msg("failed unloading plugins")
			}
		}(pluginLoader)

		bwwPlugin, err := pluginLoader.PluginForFileExtension(".bww")
		if err != nil {
			log.Fatal().Err(err).Msgf("failed getting plugin for extension .bww")
		}

		allFiles, err := getAllFilesFromPaths(args)
		if err != nil {
			return fmt.Errorf("failed getting files: %s", err.Error())
		}
		if argVerbose {
			log.Info().Msg("Processing files: ")
			for _, file := range allFiles {
				log.Info().Msg(file)
			}
		}

		if !filepath.IsAbs(OutputDir) {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			OutputDir = filepath.Join(wd, OutputDir)
		}
		err = os.MkdirAll(OutputDir, 0755)
		if err != nil {
			return fmt.Errorf("failed creating output directory: %s", err.Error())
		}

		if argVerbose {
			log.Info().Msgf("successful parsed files will be moved to: %s", OutputDir)
		}

		allFileCnt := len(allFiles)
		for i, file := range allFiles {
			fileData, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			if argVerbose {
				log.Info().Msgf("parsing file %d/%d %s", i+1, allFileCnt, file)
			}
			tunesImport, err := bwwPlugin.Import(fileData)
			if err != nil {
				if argSkipFailedFiles {
					log.Error().Err(err).Msgf("failed parsing file %s", file)
					continue
				}

				return fmt.Errorf("failed parsing file %s: %v", file, err)
			}

			if err == nil {
				fileName := filepath.Base(file)
				newPath := filepath.Join(OutputDir, fileName)
				err = os.Rename(file, newPath)
				if err != nil {
					log.Error().Err(err).Msgf("failed moving file %s to dir %s", file, OutputDir)
				}
			}

			if argVerbose {
				log.Info().Msgf("(%d/%d) successfully parsed %d tunes from file %s",
					i+1,
					allFileCnt,
					len(tunesImport.ImportedTunes),
					file,
				)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(parserCmd)

	addVerbose(parserCmd)
	addDryRun(parserCmd)
	addRecursive(parserCmd)
	addImportFileTypes(parserCmd)
	addSkipFailedFiles(parserCmd)
	parserCmd.Flags().StringVarP(
		&OutputDir,
		"output-dir",
		"o",
		DefaultOutputDir,
		"Output directory where to move the successful parsed files into",
	)
}
