package main

import (
    "fmt"
    "log"
    "net/http"
    "blog-management/database"
    "blog-management/handlers"
)

func main() {
    // Initialize the database
    database.InitDB()

    // Define routes
    http.HandleFunc("/blog", handlers.CreateBlog)      // POST
    http.HandleFunc("/blog/", handlers.GetBlogByID)    // GET by ID, PUT, DELETE
    http.HandleFunc("/blogs", handlers.GetAllBlogs)    // GET all

    // Start the server
    fmt.Println("Starting server at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
