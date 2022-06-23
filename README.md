# jumia-task
This task is a simple API for validating numbers using 
regex, manipulating given database and rendering required data.

This task has two parts
- backend (golang)
- frontend (react)

##Technologies Used for backend
- Golang
- SQLite
- Gin-Gonic Router

##Scope of the Task
- OOP skills
- Standard/Clean code
- Regex

## Running the application
### Backend
- sync all dependencies with the command `go mod tidy`
- run the backend with the command `make run`

If .env file with port is not available, the backend will run on port 8081

### Frontend
- Open the jumia-frontend directory on a new window/terminal
- install all dependencies with command `npm i`
- run the frontend with the command `npm start`
- check your axios file to configure your port. Default port at axios.js is 8081


##Testing Backend
- Mock the Database with the command `make mock`
- Handler Test with mocked DB can be found in /handler/test directory
- Unit test for services can be found in /services directory