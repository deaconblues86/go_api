package main


import (
    "github.com/gin-gonic/gin"
    "github.com/go_api/controllers"
)


func setupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/posts", controllers.GetPosts)
    router.POST("/posts", controllers.PostPosts)

    router.GET("/posts/:id", controllers.GetPost)
    router.PUT("/posts/:id", controllers.PutPost)
    router.DELETE("/posts/:id", controllers.DeletePost)

    router.POST("/comments", controllers.PostComments)

    router.GET("/comments/:id", controllers.GetComment)
    router.PUT("/comments/:id", controllers.PutComment)
    router.DELETE("/comments/:id", controllers.DeleteComment)

    return router
}


func main() {
    r := setupRouter()
    r.Run("localhost:8080")
}