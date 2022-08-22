# gocrud
###### Crud operataions with go.

## Features

- go 1.18
- gin-gonic for routing
- Postgre sql
- gorm for PostgreSQL database access
- swagger for api documentation
- go modules for package management
- dockerized environment

## Installation & Run

##### # docker-compose
- Setup docker environment use "docker-compose.yml" to docker-compose up.
- Go to project folder directory and execute
```sh
# start docker environment
$ docker-compose up -d --build
# list running services
$ docker-compose ps
# stop all containers
$ docker-compose stop
# remove all containers
$ docker-compose rm
```

##### # Run Locally

- Download posgresql from https://www.postgresql.org/download/ and Setup
- On macos download from https://postgresapp.com 

- If download the server find pg_hba.conf file and change IPv4 and IPv6 connections trust (not scram-sha-256 or md5)
- Setup postgresql connection variables with local.env file or just setup yours and change parameters inside local.env
```sh
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_DB=DBNAME
POSTGRES_PASSWORD=STRONGPASSWORD
```

- After installation create DB on postgre with your given "DBNAME" then just go to project directory and execute
```sh
$ go run main.go
```

> ✨-postgre sql runs automatically on 5432 port
> ✨ server run on 8080 port
## Test

> Create mocks : Go to project folder directory and execute
```sh
$ go generate ./...
```
> Run Tests :  Go to project folder directory and execute
```sh
$ go test -v ./pkg/service/...
```
## Usage

####  - API

##### # GetAllUsers 
- GET http://localhost:8080/users
```sh
curl --location --request GET 'http://localhost:8080/users'
```

##### #GetUserById 
- GET http://localhost:8080/users/1
```sh
curl --location --request GET 'http://localhost:8080/users/1'
```

##### # CreateUser
- POST http://localhost:8080/users
- body:  { "name": "NurcanNew", "email": "NurcanNew@mail.com", "password": "NurcanNew" }
```sh
curl --location --request PATCH 'http://localhost:8080/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{"name":"NurcanNew", "email":"NurcanNew@mail.com", "password":"NurcanNew"}'
```

##### # UpdateUser
- PATCH http://localhost:8080/users/1
- body:  { "name": "NurcanNew", "email": "NurcanNew@mail.com", "password": "NurcanNew" }
```sh
curl --location --request PATCH 'http://localhost:8080/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{"name":"NurcanNew", "email":"NurcanNew@mail.com", "password":"NurcanNew"}'
```

##### # DeleteUser
- DELETE http://localhost:8080/users/1
```sh
curl --location --request DELETE 'http://localhost:8080/users/1' \
```

>  Note: Response model showen on below
```sh
type RestErr struct {
  Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
```
