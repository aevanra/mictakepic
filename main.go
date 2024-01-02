package main

import (
    "net/http"
    "log"

    "github.com/aevanra/mictakepic/pkg/auth"
    "github.com/aevanra/mictakepic/middleware"
    "github.com/aevanra/mictakepic/pkg/filesharing"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
    r := gin.Default()

    // Load HTML and Static files
    r.LoadHTMLGlob("./assets/*")
    r.Static("/Shares", "./Shares")
    r.Static("/static", "./static")

    // Routes

    // Landing Page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html",
        gin.H{})
    })
    
    //Login page
    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    })
    r.POST("/auth", auth.LoginPOSTHandler)

    // Auth-Handling Routes
    users := r.Group("/users", middleware.RequireAuth())
    users.GET("/homepage", auth.UserHomeHandler)
    users.GET("/createUser", func(c *gin.Context) {
        c.HTML(http.StatusOK, "createUser.html", gin.H{})
    })
    users.POST("/createUser", auth.CreateNewUser)
    users.GET("/getImages", smb.LoadImageGETHandler)

    err = r.Run(":8082") // listen and serve on
    if err != nil {
        log.Fatal(err)
    }

}

