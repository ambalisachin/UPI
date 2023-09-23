package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"strconv"
// )

// // Models

// type User struct {
// 	ID   int
// 	Name string
// }

// type Database interface {
// 	GetUser(id int) (*User, error)
// }

// // Database Implementations

// type realDatabase struct {
// 	// ... connection info, etc.
// }

// func (db *realDatabase) GetUser(id int) (*User, error) {
// 	// Here would be the real database logic.
// 	return &User{ID: id, Name: "Real User"}, nil
// }

// type mockDatabase struct{}

// func (db *mockDatabase) GetUser(id int) (*User, error) {
// 	return &User{ID: id, Name: "Mocked User"}, nil
// }

// // Handlers

// func GetUserHandler(db Database) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		idStr := c.Param("id")
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 			return
// 		}

// 		user, err := db.GetUser(id)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func main() {
// 	r := gin.Default()
// 	// Using mock database for demonstration.
// 	// In a real-world scenario, you'd use something like: db := &realDatabase{}
// 	db := &mockDatabase{}
// 	r.GET("/user/:id", GetUserHandler(db))

// 	r.Run(":8080")
// }
