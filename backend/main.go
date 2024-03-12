package main

import (
	"log"
	"net/http"
	"os"
    "sync"

	"github.com/aevanra/mictakepic/middleware"
	"github.com/aevanra/mictakepic/pkg/auth"
	"github.com/aevanra/mictakepic/pkg/objects"
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
        homeShare := os.Getenv("HOME_IMAGES_DIR_NAME")
        homeImages := filesharing.ListImagesFromShare(homeShare)
        c.HTML(http.StatusOK, "index.html",
        gin.H{"images": homeImages})
    })
    r.GET("/listHomeImages", func(c *gin.Context) {
        homeShare := os.Getenv("HOME_IMAGES_DIR_NAME")
        homeImages := filesharing.ListImagesFromShare(homeShare)

        //Make list to hold Image Objects
        returnImages := make([]obj.Image, 0, len(homeImages))

        //Make Channel to get metadata concurrently
        imageCH := make(chan obj.Image, len(homeImages)+1)

        //Make WaitGroup to force all images to finish
        wg := sync.WaitGroup{}

        for _, imageName := range(homeImages) {
            // Add count of concurrent to wait for
            wg.Add(1)

            go filesharing.GetImageDimensions(imageName, homeShare, imageCH, &wg)
        }

        //Wait for channel and close
        wg.Wait()
        close(imageCH)

        for image := range(imageCH) {
            returnImages = append(returnImages, image)
        }

        c.JSON(http.StatusOK, obj.ImageList{Images: returnImages}.Sort())
    })
    r.GET("/listShareImages", func(c *gin.Context) {
        shareImages := filesharing.ListImagesFromShare(c.Query("shareName"))
        if len(shareImages) > 0 {
            c.JSON(http.StatusOK, shareImages)
            return
        }

        c.JSON(http.StatusNoContent, "No images found")
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

