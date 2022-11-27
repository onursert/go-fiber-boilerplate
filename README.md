# Go Fiber Boilerplate
Golang Rest API boilerplate built with Go-Fiber and MongoDB database with tests.

# File structure
```
configs/
    env.go
db/
    db.go
mocks/
    repositories/
        mock_todoRepository.go
    services/
        mock_todoService.go
models/
    todoModel.go
routes/
    todoRoute.go
controllers/
    todoController.go
    todoController_test.go
services/
    todoService.go
    todoService_test.go
repositories/
    todoRepository.go
utils/
    errors.go
main.go
.env
```

# Dependencies
Install dependencies with go
`go mod tidy`