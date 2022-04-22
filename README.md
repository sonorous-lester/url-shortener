# url-shortene

## 如何使用 

將 `config-template.yml` 並更名為 `config.yml` 並填上其相關資訊。

範例
```ymal
server:
  host: localhost:6666
  port: :6666
db:
  userName: your-name
  password: your-password
  addr: 127.0.0.1
  port: 9898
  database: database-name
  maxOpenConns: 10
  maxIdleConns: 10
nosql:
  network: tcp
  addr: :6379
  maxIdle: 80
  maxActive: 12000
```

開啟 terminal 移動到該資料夾 `go run main.go` 

## 使用 Library

- [Gin](https://github.com/gin-gonic/gin)
- [Mysql-driver](https://github.com/go-sql-driver/mysql)
- [Redigo](https://github.com/gomodule/redigo)
- [Shortuuid](https://github.com/lithammer/shortuuid)
- [Yaml](https://github.com/go-yaml/yaml)

## 待辦清單
- [X] Logging
- [ ] Rate Limiter
- [X] Improve valid url mechanism
- [ ] Testing
