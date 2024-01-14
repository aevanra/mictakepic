package filesharing

import (
    "path/filepath"
    "fmt"

    "github.com/gin-gonic/gin"
)

func DownloadFileGETHandler(c *gin.Context) {
    share := c.Query("share")
    filename := c.Query("filename")
    fullPath, err := filepath.Abs("Shares/" + share + "/" + filename)
    fmt.Println(fullPath)
    fmt.Println(err)

    if err != nil {
        c.AbortWithStatus(404)
        return
    }

    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; filename="+filename)
    c.Header("Content-Type", "application/octet-stream")

    c.FileAttachment(fullPath, filename)
}
