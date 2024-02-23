# Go Template Server
A template server for go projects using Gin

## Usage
* Click the `Use this template` button to create a new repository using this template
* Clone the new repository
* Delete the `go.sum` file
* Use `go mod init <project name>` to create a new `go.mod` file
* Run `go mod tidy` to download the required dependencies

## Aproach 
* All the API code is in the `internal` package
    * `handlers` contains the API handlers which will define the routes of the API
    * `dto` contains all the model structs for the API
    * `controllers` contains the business logic for the API
    * `app` contains the main function to start the server

* The API test code will be written inside the `internal/controllers` dir only and will follow table driven tests approach

* `pkg` contains the common code which can be used across multiple projects. We can add directories in like mongo, repository, etc.

* `cmd` contains the main function to start the server

* `util` contains the utility functions which can be used across the project.

* Controllers and Handlers will not be responsible for sending of data to the client. They will only be responsible for processing the data and sending it to the client. 

* The response will be sent by the `util.HandlerWrapper` function. This makes the code more testable and maintainable.
