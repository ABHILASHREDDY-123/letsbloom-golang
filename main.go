package main

import (
	"database/sql"
	"io"
	"os"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
// disable coloring in logging with default logger
       gin.DisableConsoleColor()
       
// logging into file
    file,_ :=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/library")
       if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	router := gin.Default()
	router.GET("/api/books", func(ctx *gin.Context) {
		jsonObject := map[string]interface{}{
			"audi": 4.8,
			"benz": 4.9,
		}
		ctx.AsciiJSON(200, jsonObject)
	})

	router.POST("/api/books", func(ctx *gin.Context) {
		jsonObject := map[string]interface{}{
			"audi": 4.8,
			"benz": 4.9,
		}
		ctx.AsciiJSON(200, jsonObject)
	})

	router.PUT("/api/books/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		jsonObject := map[string]interface{}{
			"id":id,
		}
		ctx.AsciiJSON(200, jsonObject)
	})
	



	// serve router
	router.Run(":8080")
}
