package main

import "github.com/liu-houliang/goencrypt/gin"

func main() {

	ginServer := gin.NewRouter()
	err := ginServer.Run(":8082")
	if err != nil {
		panic(err)
		return
	}

}
