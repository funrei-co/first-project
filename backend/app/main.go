package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main()  {
    // MySQL データベースへの接続文字列
    db, _ := sql.Open("mysql", "root@tcp(myapp-mysql-1:3306)/chat_develop")
    var user_Id, user_Name string
    
    db.QueryRow("SELECT * FROM user where user_Id = 'ksasaki';",).Scan(&user_Id, &user_Name)
    fmt.Println(user_Id, user_Name)
    }


