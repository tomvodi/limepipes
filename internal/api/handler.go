package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tomvodi/limepipes/internal/api_gen/apimodel"
	api_interfaces "github.com/tomvodi/limepipes/internal/api_gen/interfaces"
	"github.com/tomvodi/limepipes/internal/common"
	"github.com/tomvodi/limepipes/internal/interfaces"
	"io"
	"net/http"
	"path/filepath"
)

type apiHandler struct {
	service       interfaces.DataService
	pluginLoader  interfaces.PluginLoader
	healthChecker interfaces.HealthChecker
}

func (a *apiHandler) Home(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (a *apiHandler) Health(c *gin.Context) {
	handler, err := a.healthChecker.GetCheckHandler()
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleFunc := gin.WrapH(handler)
	handleFunc(c)
}

func (a *apiHandler) ImportFile(c *gin.Context) {
	iFile, err := c.FormFile("file")
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	fExt := filepath.Ext(iFile.Filename)
	if fExt == "" {
		c.JSON(http.StatusBadRequest,
			apimodel.Error{
				Message: "import file does not have an extension",
			},
		)
		return
	}

	fType, err := a.pluginLoader.FileTypeForFileExtension(fExt)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apimodel.Error{
				Message: fmt.Sprintf("file extension %s is currently not supported: %s", fExt, err.Error()),
			},
		)
		return
	}

	fileReader, err := iFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apimodel.Error{
				Message: fmt.Sprintf("failed open file %s for reading", iFile.Filename),
			},
		)
		return
	}
	defer fileReader.Close()

	fileData, err := io.ReadAll(fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			fmt.Sprintf("failed reading file %s: %s", iFile.Filename, err.Error()))
		return
	}

	fInfo, err := common.NewImportFileInfo(iFile.Filename, fType, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			fmt.Sprintf("failed creating import file info for file %s: %s", iFile.Filename, err.Error()))
		return
	}
	_, err = a.service.GetImportFileByHash(fInfo.Hash)

	if !errors.Is(err, common.NotFound) {
		c.JSON(http.StatusConflict,
			fmt.Sprintf("file %s was already imported", iFile.Filename))
		return
	}

	importTunes, importSet, err := a.importFile(fInfo, fExt)
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	importResponse := &apimodel.ImportFile{
		Name:  iFile.Filename,
		Set:   *importSet,
		Tunes: importTunes,
	}

	c.JSON(http.StatusOK, importResponse)
}

func (a *apiHandler) importFile(
	fInfo *common.ImportFileInfo,
	fileExt string,
) ([]*apimodel.ImportTune, *apimodel.BasicMusicSet, error) {
	filePlugin, err := a.pluginLoader.PluginForFileExtension(fileExt)
	if err != nil {
		return nil, nil, fmt.Errorf("fileData extension %s is currently not supported (no plugin): %s", fileExt, err.Error())
	}

	parsedTunes, err := filePlugin.Import(fInfo.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("failed parsing fileData %s: %s", fInfo.Name, err.Error())
	}

	return a.service.ImportTunes(parsedTunes.ImportedTunes, fInfo)
}

func httpErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(code, apimodel.Error{
		Message: err.Error(),
	})
}

func handleResponseForError(c *gin.Context, err error) {
	code := http.StatusInternalServerError
	if errors.Is(err, common.NotFound) {
		code = http.StatusNotFound
	}

	c.JSON(code, apimodel.Error{
		Message: err.Error(),
	})
}

func (a *apiHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (a *apiHandler) CreateTune(c *gin.Context) {
	var createTune apimodel.CreateTune
	if err := c.ShouldBindJSON(&createTune); err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tune, err := a.service.CreateTune(createTune, nil)
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tune)
}

func (a *apiHandler) GetTune(c *gin.Context) {
	tuneId, err := uuid.Parse(c.Param("tuneId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tune, err := a.service.GetTune(tuneId)
	if err != nil {
		handleResponseForError(c, err)
		return
	}

	c.JSON(http.StatusOK, tune)
}

func (a *apiHandler) ListTunes(c *gin.Context) {
	tunes, err := a.service.Tunes()
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, tunes)
}

func (a *apiHandler) UpdateTune(c *gin.Context) {
	var updateTune apimodel.UpdateTune
	if err := c.ShouldBindJSON(&updateTune); err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tuneId, err := uuid.Parse(c.Param("tuneId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tune, err := a.service.UpdateTune(tuneId, updateTune)
	if err != nil {
		handleResponseForError(c, err)
		return
	}

	c.JSON(http.StatusOK, tune)
}

func (a *apiHandler) DeleteTune(c *gin.Context) {
	tuneId, err := uuid.Parse(c.Param("tuneId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := a.service.DeleteTune(tuneId); err != nil {
		handleResponseForError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *apiHandler) CreateSet(c *gin.Context) {
	var createSet apimodel.CreateSet

	if err := c.ShouldBindJSON(&createSet); err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	set, err := a.service.CreateMusicSet(createSet, nil)
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (a *apiHandler) GetSet(c *gin.Context) {
	setId, err := uuid.Parse(c.Param("setId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	set, err := a.service.GetMusicSet(setId)
	if err != nil {
		handleResponseForError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (a *apiHandler) ListSets(c *gin.Context) {
	sets, err := a.service.MusicSets()
	if err != nil {
		httpErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, sets)
}

func (a *apiHandler) UpdateSet(c *gin.Context) {
	var updateSet apimodel.UpdateSet
	if err := c.ShouldBindJSON(&updateSet); err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	setId, err := uuid.Parse(c.Param("setId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	set, err := a.service.UpdateMusicSet(setId, updateSet)
	if err != nil {
		handleResponseForError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (a *apiHandler) DeleteSet(c *gin.Context) {
	setId, err := uuid.Parse(c.Param("setId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := a.service.DeleteMusicSet(setId); err != nil {
		handleResponseForError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *apiHandler) AssignTunesToSet(c *gin.Context) {
	var tuneIds []uuid.UUID
	if err := c.ShouldBindJSON(&tuneIds); err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	setId, err := uuid.Parse(c.Param("setId"))
	if err != nil {
		httpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	set, err := a.service.AssignTunesToMusicSet(setId, tuneIds)
	if err != nil {
		handleResponseForError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func NewApiHandler(
	service interfaces.DataService,
	pluginLoader interfaces.PluginLoader,
	healthChecker interfaces.HealthChecker,
) api_interfaces.ApiHandler {
	return &apiHandler{
		service:       service,
		pluginLoader:  pluginLoader,
		healthChecker: healthChecker,
	}
}
