package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ABHILASHREDDY-123/letsbloom-golang/models"
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

    // creating database 
	createDataBase := `CREATE TABLE IF NOT EXISTS Book (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50),
		author_name VARCHAR(50),
		publish_date VARCHAR(50),
		image_url VARCHAR(50),
		location VARCHAR(50)
	)`

	 db.Exec(createDataBase)
	 


	router := gin.Default()
	router.GET("/api/books", func(ctx *gin.Context) {
		getQuery := `SELECT * from Book`
		rows,err := db.Query(getQuery)
		if(err!=nil){
			ctx.JSON(401,gin.H{"error": err.Error()})
			return
		}
		var books []models.Book;
		for rows.Next(){
            var book models.Book;

			err := rows.Scan(&book.Id,&book.Name,&book.Author,&book.PublishedDate,&book.Image,&book.Location)
			if(err!=nil){
				ctx.JSON(401,gin.H{"error": err.Error()})
				return
			}	
			books = append(books, book)
		}

		ctx.JSON(200,gin.H{"books":books})
	})

	router.POST("/api/books", func(ctx *gin.Context) {
		var book models.Book;
		if err := ctx.BindJSON(&book); err != nil {
			ctx.JSON(401,gin.H{"error": err.Error()})
			return
		}
		insertQuery := `INSERT INTO Book (name,author_name,publish_date,image_url,location) VALUES (?,?,?,?,?)`
		_, err := db.Exec(insertQuery,book.Name,book.Author,book.PublishedDate,book.Image,book.Location)
		if(err!=nil){
			ctx.JSON(401,gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200,gin.H{"message":"Successfully Added Book"})

	})

	router.PUT("/api/books/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		var updatedBook models.Book;
		if err := ctx.BindJSON(&updatedBook); err != nil {
			ctx.JSON(401,gin.H{"error": err.Error()})
			return
		}
		var fields []interface{}
		var values []interface{}
		if(len(updatedBook.Location)>0){
			fields = append(fields, "location")
			values = append(values,updatedBook.Location)
		}
		if(len(updatedBook.Name)>0){
			fields = append(fields, "name")
			values = append(values,updatedBook.Name)
		}
		if(len(updatedBook.Image)>0){
			fields = append(fields, "image_url")
			values = append(values,updatedBook.Image)
		}
		if(len(updatedBook.PublishedDate)>0){
			fields = append(fields, "publish_date")
			values = append(values,updatedBook.PublishedDate)
		}
		if(len(updatedBook.Author)>0){
			fields = append(fields, "author_name")
			values = append(values,updatedBook.Author)
		}
		query := "UPDATE Book SET "
		for r:= 0;r<len(fields);r++ {
			if r == len(fields)-1 {
				query+=fmt.Sprintf("%s = ? ",fields[r])
			} else {
             query+=fmt.Sprintf("%s = ?, ",fields[r])
			}
		}
		query+=fmt.Sprintf(" WHERE %s = ?","id")
		values = append(values, id)
		fmt.Println(query)
		_, err := db.Exec(query,values...)
		if err != nil {
			ctx.JSON(401,gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"message":"Sucessfully updated"});
		
	})
	// serve router
	router.Run(":8080")
}
