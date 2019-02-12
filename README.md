# Lebaros Backend
Golang REST API with MySQL integration using Gin and GORM.

## Requirements
| Require | Minimum Version | Recomended Version |
| ------ | ------ | ------ |
| Go | v1.7 | v1.11 |
| Dep | - | devel |
| Docker Compose | 1.18 | 1.23 |

## Project Installation
- Clone the repository `$ git clone https://github.com/abaron/lebaros-backend.git
`
- Go to project dir `$ cd lebaros-backend`
- Intall dependencies `$ dep ensure`
- Install gorm and gin `go get github.com/jinzhu/gorm && go get github.com/codegangsta/gin`
- Install gin `go get -u github.com/gin-gonic/gin`

GORM should be installed via `go get` since installation via `dep` is imperfect (it does not download dialects directory).

[codegangsta/gin](https://github.com/codegangsta/gin) is an optional package to install if you want to make usage of live reloading feature of server (just like nodemon in Node.js environment).


GORM should be installed via `go get` since installation via `dep` is imperfect (it does not download dialects directory).

[codegangsta/gin](https://github.com/codegangsta/gin) is an optional package to install if you want to make usage of live reloading feature of server (just like nodemon in Node.js environment). 

### Configure Environment Variables
Open .env file and edit the values if you need to. This project uses [godotenv](https://github.com/joho/godotenv) to read and use .env file. 

## Start Project
Run docker first `$ docker-compose up` or `$ sudo docker-compose up`
Then
```
$ go run main.go
```

To explicitly compile the code before you run the server:

```
$ go build main.go
$ ./main
```

To use live-reloading in development environment, 

```
$ ./scripts/start-dev
```

MySQL Port `33060`
MySQL Connection `root:root@172.17.0.1:33060/lebaros` or `lebaros:lebaros@127.0.0.1:3306/lebaros`

## TODO
- [x] Import from existing data
- [x] Actual Stock
- [x] Stock in
- [x] Stock out
- [ ] Import from file upload
- [ ] Request validation
- [ ] Report sales report
- [ ] Report stock report
- [ ] Swagger Docs
- [ ] GUI version

Thank you