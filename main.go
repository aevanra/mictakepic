package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aevanra/mictakepic/middleware"
	"github.com/aevanra/mictakepic/pkg/auth"
	"github.com/aevanra/mictakepic/pkg/filesharing"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    r := gin.Default()

    // Load HTML and Static files
    r.LoadHTMLGlob("./assets/*")
    r.Static("/Shares", "./Shares")
    r.Static("/static", "./static")

    // Routes

    // Landing Page
    r.GET("/", func(c *gin.Context) {
        homeImages := filesharing.ListImagesFromShare(os.Getenv("HOME_IMAGES_DIR_NAME"))
        c.HTML(http.StatusOK, "index.html",
        gin.H{"images": homeImages})
    })
    
    //Login page
    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    })
    r.POST("/auth", auth.LoginPOSTHandler)
    r.GET("/createUser", func(c *gin.Context) {
        c.HTML(http.StatusOK, "createUser.html", gin.H{"Shares": filesharing.ListDataShareNames() })
    })
    r.POST("/createUser", auth.CreateNewUser)

    // Auth-Handling Routes
    users := r.Group("/users", middleware.RequireAuth())
    users.GET("/homepage", auth.UserHomeHandler)
    users.GET("/createUser", func(c *gin.Context) {
        c.HTML(http.StatusOK, "createUser.html", gin.H{"Shares": filesharing.ListDataShareNames() })
    })
    users.POST("/createUser", auth.CreateNewUser)
    users.GET("/getImages", filesharing.LoadImageGETHandler)

    // File Service Routes
    fileservice := r.Group("/files", middleware.RequireAuth(), middleware.ConfirmShareAccess())
    fileservice.GET("/download", filesharing.DownloadFileGETHandler)
    err = r.Run(":8082") // listen and serve on
    if err != nil {
        log.Fatal(err)
    }

}

