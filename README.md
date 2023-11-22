# gitplugin

`gitplugin` is a ChatGPT plugin written in Go, providing a simple yet effective way to interact with GitHub repositories. This plugin is designed to be integrated into ChatGPT applications to enhance their capabilities in fetching and displaying content from GitHub repositories.

## Features

1. **Repository Structure Retrieval:**
   - Fetch the structure of a specified GitHub repository, including file paths and types.

2. **File Content Retrieval:**
   - Retrieve the content of a specific file within a GitHub repository.

## Components

The plugin consists of the following main components:

- **Main Application (`main.go`):**
  - Sets up an Echo server.
  - Defines routes to handle POST requests for fetching repository structures and file contents.

- **Service Layer (`service` directory):**
  - `gitService.go` contains the logic for handling the API requests to GitHub.
  - Functions `GetStructure` and `GetFile` process the requests, interacting with the GitHub API to fetch the desired data.

- **Data Models (`model` directory):**
  - `structures.go` defines the data structures used in the application.
  - `Repository`, `File`, `GithubResponse`, and `TreeItem` models are used to parse and structure the data retrieved from GitHub.

- **Additional Configuration Files:**
  - `.gitignore`, `Dockerfile`, `go.mod`, `go.sum`, and `go.work` for setup and dependencies management.

## Usage

The service can be accessed via POST requests to the Echo server:

1. **Get Repository Structure:**
   - Endpoint: `/repo`
   - Payload: JSON object specifying the `owner`, `repo`, and `branch` of the GitHub repository.

2. **Get File Content:**
   - Endpoint: `/file`
   - Payload: JSON object specifying the `owner`, `repo`, `branch`, and `path` of the file in the GitHub repository.

## Setup

The application requires a Go environment for running and can be deployed in two ways:

1. **As a Standalone Service:**
   - Install the Go environment.
   - Manage dependencies as listed in `go.mod`.

2. **Using Docker:**
   - A `Dockerfile` is provided for building a Docker image of the application.
   - Build the Docker image and run the application inside a container for an isolated and consistent deployment environment.

