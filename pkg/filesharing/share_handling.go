package filesharing

import (
    "os"
    "slices"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/aevanra/mictakepic/pkg/session"
    "github.com/aevanra/mictakepic/pkg/objects"
)


func ListDataShares() []obj.DataShare {
    datashares := make([]obj.DataShare, 0)
    entries, _ := os.ReadDir("./Shares")

    for _, dir := range(entries) {
        if dir.IsDir() {
            datashares = append(datashares, obj.DataShare{ShareName: dir.Name()})
        }
    }

    return datashares
}

func ListDataShareNames() []string {
    shares := ListDataShares()
    names := make([]string, 0)

    for _, share := range(shares) {
        names = append(names, share.ShareName)
    }

    return names
}

func ValidateUserDatashare(user obj.User) bool {
    dataShares := ListDataShares()
    return slices.Contains(dataShares, obj.DataShare{ShareName: user.DefaultDataShare})

}

func ListImagesFromShare(share string) []string {
    images, _ := os.ReadDir("./Shares/" + share)
    names := make([]string, 0)

    for _, image := range(images) {
        names = append(names, image.Name())
    }

    return names
}

func LoadImageGETHandler(c *gin.Context) {
    val := session.GetSessionValue(c.Request, "User")
    user := val.(*obj.User)
    share := c.Query("share")

    if !slices.Contains(user.AllDatashares, share) {
        c.AbortWithStatus(http.StatusUnauthorized)
    }
    filesList := ListImagesFromShare(share)

    c.HTML(http.StatusOK, "filelist.html", gin.H{"images": filesList, "share": share})

}
