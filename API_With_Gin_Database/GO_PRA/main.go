
package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"database/sql"
)

var db *sql.DB

type User struct{
	Id int
	Name string
	Email string
	Age string
}

func main()  {

	db,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/blogs")

	if err!=nil{
		panic(err.Error())
	}

	defer db.Close()





	
	
}