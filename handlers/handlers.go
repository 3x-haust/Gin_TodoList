package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/3x-haust/Go_TodoList/database"
    "github.com/3x-haust/Go_TodoList/models"
)

func CreateTodo(c *gin.Context) {
    var newTodo models.Todo
    if err := c.BindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    result := database.DB.Create(&newTodo)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, newTodo)
}

func ListTodos(c *gin.Context) {
    var todos []models.Todo
    result := database.DB.Find(&todos)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    result := database.DB.Delete(&models.Todo{}, id)
    
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}