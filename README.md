![limepipes logo](doc/img/logo.png)

## About

LimePipes is an application for handling and storing music for the great highland bagpipes. It parses `.bww` files from the famous BagpipePlayer from Doug Wickstrom and `.bmw` files from the Bagpipe Music Writer by Robert MacNeil Musicworks and stores them into a database. From there, they can be retrieved and converted to other formats like the more common MusicXML format.

Today, many well known music notation programs are using the MusicXML for interchanging musical compositions.
There has been a bww2mxml tool that ships with [MuseScore](https://musescore.org) but it lacks importing `.bww` files that go beyond basic tunes with no complex time lines (e.g. 2 of 4).

## Project Structure

`cmd`

The project uses the API defined in the [limepipes-api](https://github.com/tomvodi/limepipes-api) project and creates the backend service for it in the directory `limepipes`. The `limepipes-cli` directory contains a command line application which makes it easy to import many tunes at once without the need for a GUI.

`internal/api`

The generated code from the OpenAPI spec and the implementation of the REST service.

`internal/bww`

All the lexing and parsing of `.bww` files is taking place here.

`common/music_model`

This directory contains a music model definition which represents the parsed tunes from a `.bww` file.

`internal/database`

Currently, the parsed tunes are stored in a SQlite database but can be changed to any backend for [GORM](https://gorm.io).

`exporter/musicxml`

Everything used to export the internal music model to a MusicXML file.

`limepipes-api`

The git submodule that contains the OpenAPI spec for the REST API.

## Build

`go mod download` downloads all required dependencies

`go build -o ./limepipes github.com/tomvodi/limepipes/cmd/limepipes` builds the executable

## Run

The limepipes application needs a configuration file called `limepipes.env` beside the executable or the environment
variables from this file directly set to the environment of the application. 

## Develop

### Prerequisites

In order to build the server code from the OpenAPI spec, you need to have the [OpenAPI Generator installed](https://openapi-generator.tech/docs/installation/) 
and in your PATH. You also have to run `git submodule init` and `git submodule update` to get the OpenAPI spec.
the `scripts` directory contains a script `generate_server.sh` to generate the server code from the OpenAPI spec.

Mocks are generated with `[vektra/mockery](https://vektra.github.io/mockery/latest/installation)` this also has to 
be in your PATH.

### Configuration and Environment

The application uses a PostgreSQL database to store data. To setup a local database, you can use the
`docker-compose.yml` file in the project directory. This will start a PostgreSQL database in a Docker container and uses
the `db.env` file for configuration. This file also has a template `db.env.default` which can copied and renamed to `db.env`.

The REST API is served over HTTPS and needs a certificate and key file. The `Makefile` has a target `create_test_certificates` 
which generates these files in the `build` directory for development and test purposes and must not be used in production.

The application gets its configuration from an `limepipes.env` file which needs to be in the same directory as the executable.
There is a `limepipes.env.default` which can be used as a template. In this file, you can set variables for database connection,
the paths to the certificate and key files and some other application relevant settings.


### Build

The `Makefile` contains targets for many build and test tasks. Some of the Makefile targets also have a 
corresponding run configuration for Intellij IDEs like GoLand in the `.idea/runConfigurations` directory.
The `test` directory contains `.http` test files used by GoLand from JetBrains to test the REST API manually.

### Source Code

As previously mentioned, the application uses an intermediate music model for storing the parsed `.bww` tunes. 
This model is defined in the `common/music_model` directory and has struct tags to be exported to `.yaml` files. This 
is used by many tests to compare the parsed tunes with the expected output.

All enums that have to be serialized and deserialized to and from yaml, are handled with [Enumer](https://github.com/dmarkham/enumer)
generate the necessary code.

### Used libraries and tools

- [GORM](https://gorm.io) for database access
- [OpenAPI Generator](https://openapi-generator.tech) for the REST API
- [Enumer](https://github.com/dmarkham/enumer) for enum handling like serializing and deserializing
- [Cobra](https://github.com/spf13/cobra) for command line argument parsing
- [Viper](https://github.com/spf13/viper) for configuration handling
- [Gin](https://github.com/gin-gonic/gin) for the REST API server
- [Enumer](https://github.com/dmarkham/enumer) for enum serializing and deserializing
- [Mockery](https://vektra.github.io/mockery/latest/installation) for mocking interfaces
- [Docker and Docker Compose](https://docs.docker.com/compose/) for running the application in a container



