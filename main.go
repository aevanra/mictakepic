package main

import (
    "net/http"
    "log"

    "github.com/aevanra/mictakepic/pkg/auth"
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
    
    //Login page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Main website",
        })
    })

    // Auth-Handling Routes
    r.POST("/auth", auth.LoginPOSTHandler)

    r.GET("/createUser", func(c *gin.Context) {
        c.HTML(http.StatusOK, "createUser.html", gin.H{
            "title": "New User Registration",
        })
    })
    r.POST("/createUser", auth.CreateNewUser)
    r.GET("/listFiles", smb.LoadImageGETHandler)

    err = r.Run(":8082") // listen and serve on
    if err != nil {
        log.Fatal(err)
    }

}

