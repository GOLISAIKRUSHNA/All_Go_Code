package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type User struct{
	Id int
	FirstName  string
	LastName string
	Email string

}


func main() {

	db, err := gorm.Open(mysql.Open("root:root@/Test_MS"), &gorm.Config{})

	if err != nil {
		panic("not connect to database connection")
	}

	fmt.Println("database connected",db)

	db.AutoMigrate(&User{})

	


	// u:=User{
	// 	FirstName: "Tejas",
	// 	LastName: "Surve",
	// 	Email:"Tejas.surve@gmail.com",
	// }


	// up:=User{
	// 	Id :1,
	// 	FirstName: "goli",
	// 	LastName: "saikrushna",
	// 	Email: "golisaikrushna@gmail.com",
	// }

	del:=User{
		Id:2,
	}

	// db.Create(&u)
	fmt.Println("created successfully")

	// db.Updates(&up)

	db.Delete(&del)

	

	







}

