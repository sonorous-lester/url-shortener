package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"url-shortener/internal/service/config"
)

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS url(id int auto_increment, origin_url text,  
        url_id varchar(32), expired_at varchar(32), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP, primary key (id))`
	insertQuery = `INSERT INTO url (origin_url, url_id, expired_at) VALUES (?, ?, ?)`
)

type Sql struct {
	server *sql.DB
}

func NewSql(c config.Config) Sql {
	fmt.Println(c)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Db.UserName, c.Db.Password, c.Db.Addr, c.Db.Port, c.Db.Database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		panic(err.Error())
	}

	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(c.Db.MaxOpenConns)
	db.SetMaxIdleConns(c.Db.MaxIdleConns)

	_, err = db.Exec(createTableQuery)

	if err != nil {
		fmt.Println("create table failed:", err)
		panic(err.Error())
	}
	return Sql{server: db}
}

func (db Sql) Store(longUrl, urlId, expireAt string) error {
	_, err := db.server.Exec(insertQuery, longUrl, urlId, expireAt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db Sql) Close() error {
	return db.server.Close()
}
