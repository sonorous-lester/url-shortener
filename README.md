# 短網址服務

## 功能介紹

這是一個簡單的短網址服務。

主要提供服務有以下兩項
1. 使用者傳送原網址與到期時間，server 回傳短網址給使用者使用。
2. 使用者傳送短網址至 server ；server 會回傳 302 給使用者，轉址回原網址。
   1. 如果該短網址已過期或沒有該短網址，回傳 404 告知使用者查無資訊。

## API

### 短網址

`POST - http://{host}/api/v1/urls`

#### Request

```json
{
  "url": "https://tech.lester.tw",
  "expireAt": "2022-04-24T09:20:41Z"
}
```

#### Response

**成功** `200`
```json
{
  "id": "D7RQy9729MQR68ZGJfuxug",
  "shortUrL": "http://{host}/D7RQy9729MQR68ZGJfuxug"
}
```

**失敗** `400`
```json
{
  "message": "time format is invalid."
}
```

### 轉址

`GET - http://{host}/{urlId}`

#### Response

**成功** `302` `empty response body`

**失敗** `404`
```json
{
  "message": "not found any match url."
}
```

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
完成上方動作，你可以選擇兩種方式來啟動服務
1. 開啟 terminal 移動到該資料夾 `go run main.go` 
2. 使用 `go build`
   1. in mac `./url-shortener`
   2. in windows `url-shortener.exe`

## 使用 Library

- [Gin](https://github.com/gin-gonic/gin)
- [Mysql-driver](https://github.com/go-sql-driver/mysql)
- [Redigo](https://github.com/gomodule/redigo)
- [Shortuuid](https://github.com/lithammer/shortuuid)
- [Yaml](https://github.com/go-yaml/yaml)
- [Logrus](https://github.com/sirupsen/logrus)

## 待辦清單
- [X] Improve valid url mechanism
- [X] Logging
- [ ] Testing
- [ ] Rate Limiter
