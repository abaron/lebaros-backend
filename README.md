# Lebaros Backend
Golang REST API with MySQL integration using Gin and GORM.

## Requirements
| Require | Minimum Version | Recomended Version |
| ------ | ------ | ------ |
| Go | v1.7 | v1.11 |
| Dep | x | x |
| Docker Compose | x | x |

## Project Installation
- Clone the repository `$ git clone https://github.com/abaron/lebaros-backend.git
`
- Go to project dir `$ cd lebaros-backend`
- Intall dependencies `$ dep ensure`
- Install gorm and gin `go get github.com/jinzhu/gorm && go get github.com/codegangsta/gin`
- Install gin `go get -u github.com/gin-gonic/gin`
- Run docker `$ docker-compose up` or `$ sudo docker-compose up`

GORM should be installed via `go get` since installation via `dep` is imperfect (it does not download dialects directory).

[codegangsta/gin](https://github.com/codegangsta/gin) is an optional package to install if you want to make usage of live reloading feature of server (just like nodemon in Node.js environment).
