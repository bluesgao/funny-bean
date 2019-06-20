package main

import "funny-bean/db"
import "funny-bean/router"

func main() {
	println("funny-bean")

	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":8888")
}
