package filesharing

import (
	"image"
	_ "image/png"
    _ "image/jpeg"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"github.com/aevanra/mictakepic/pkg/objects"
	"github.com/aevanra/mictakepic/pkg/session"
	"github.com/gin-gonic/gin"
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
    VALID_FILETYPES := []string{".jpg", ".jpeg", ".png", ".gif"}
    images, _ := os.ReadDir("./Shares/" + share)
    names := make([]string, 0)

    for _, image := range(images) {
        if !image.IsDir() && slices.Contains(VALID_FILETYPES, filepath.Ext(image.Name())) {
           names = append(names, image.Name())
        }
    }

    return names
}

func getImageDimensions(filename string, share string, ch chan obj.Image, wg *sync.WaitGroup) {
    file, err := os.Open("./Shares/" + share + "/" + filename)

    if err != nil {
        wg.Done()
        return
    }

    defer file.Close()

    image, _, err := image.DecodeConfig(file)

    if err != nil {
        wg.Done()
        return
    }

    ch <- obj.Image{
        Filename: filename,
        DataShare: share,
        Height: image.Height,
        Width: image.Width,
    }

    wg.Done()
}

func GetMetadataFromImageList(images []string, share string, sorted bool) obj.ImageList {
        //Make list to hold Image Objects
        returnImages := make([]obj.Image, 0, len(images))

        //Make Channel to get metadata concurrently
        imageCH := make(chan obj.Image, len(images)+1)

        //Make WaitGroup to force all images to finish
        wg := sync.WaitGroup{}

        for _, imageName := range(images) {
            // Add count of concurrent to wait for
            wg.Add(1)

            go getImageDimensions(imageName, share, imageCH, &wg)
        }

        //Wait for channel and close
        wg.Wait()
        close(imageCH)

        for image := range(imageCH) {
            returnImages = append(returnImages, image)
        }

        imageList := obj.ImageList{Images: returnImages}

        if sorted {
            return imageList.Sort()
        } else {
            return imageList
        }
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
