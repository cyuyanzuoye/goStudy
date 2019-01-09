package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("mysql-学习使用-https://www.cnblogs.com/tsiangleo/p/4483657.html")
}

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "user:password@/dbname")
	//创建查
	fmt.Println(db)
	fmt.Println(err)
}
