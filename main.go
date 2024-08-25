package main

import (
    "github.com/gin-gonic/gin"
    "github.com/3x-haust/Go_TodoList/database"
    "github.com/3x-haust/Go_TodoList/handlers"
)

func main() {
    database.InitDB()

    r := gin.Default()

    r.POST("/todos", handlers.CreateTodo)
    r.GET("/todos", handlers.ListTodos)
    r.PUT("/todos/:id", handlers.UpdateTodo)
    r.DELETE("/todos/:id", handlers.DeleteTodo)

    r.Run(":3000")
}