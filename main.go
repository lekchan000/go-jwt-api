package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	AuthControler "github.com/lekchan000/go-jwt-api/controler/auth"
	"github.com/lekchan000/go-jwt-api/orm"
)

func main() {

	orm.InitDB()

	if orm.Db == nil {
		log.Fatal("Database connection is nil")
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/Register", AuthControler.Register)
	r.Run(":8080")
}
