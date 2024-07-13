package main

import (
    //"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)
func main() {
    // MySQL データベースへの接続文字列
    db, err := sql.Open("mysql", "root@tcp(myapp-mysql-1:3306)/chat_develop")
    
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    // データベースへの接続をテストする
    err = db.Ping()
    if err != nil {
      log.Fatal(err)
    }
    log.Println("Connected to MySQL database!")
}