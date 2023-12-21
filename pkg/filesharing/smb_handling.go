package smb

import (
    "net"
    "os"
    "slices"
    iofs "io/fs"
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/hirochachacha/go-smb2"
    "github.com/aevanra/mictakepic/pkg/session"
    "github.com/aevanra/mictakepic/pkg/objects"
)


func listDataShares() []obj.DataShare {
    datashares := make([]obj.DataShare, 0)

    _ = godotenv.Load()

    conn, err := net.Dial("tcp", os.Getenv("SMB_SHARE_HOST"))
    if err != nil {
        panic("SMB Share Connection Failed")
    }
    defer conn.Close()

    d := &smb2.Dialer{
        Initiator: &smb2.NTLMInitiator{
            User: os.Getenv("SMB_USER"),
            Password: os.Getenv("SMB_PASS"),
        },
    }

    s, err := d.Dial(conn)
    if err != nil {
        panic(err)
    }
    defer s.Logoff()

    names, err := s.ListSharenames()
    if err != nil {
        panic(err)
    }

    for _, name := range names { 
        if name != "General" && name != "IPC$"{
            datashares = append(datashares, obj.DataShare{ShareName: name})
        }
    }

    return datashares
}

func ValidateUserDatashare(user obj.User) bool {
    dataShares := listDataShares()
    return slices.Contains(dataShares, obj.DataShare{ShareName: user.DataShare})

}

func listImagesFromShare(user obj.User) []string {
    _ = godotenv.Load()

    conn, err := net.Dial("tcp", os.Getenv("SMB_SHARE_HOST"))
    if err != nil {
        panic("SMB Share Connection Failed")
    }
    defer conn.Close()

    d := &smb2.Dialer{
        Initiator: &smb2.NTLMInitiator{
            User: os.Getenv("SMB_USER"),
            Password: os.Getenv("SMB_PASS"),
        },
    }

    s, err := d.Dial(conn)
    if err != nil {
        panic(err)
    }
    defer s.Logoff()
    
    if !ValidateUserDatashare(user) {
        return nil
    }
    
    foundFiles := make([]string,0)
    fs, err := s.Mount(user.DataShare)
    if err != nil {
        panic(err)
    }
    defer fs.Umount()

    matches, err := iofs.Glob(fs.DirFS("."), "*")
    if err != nil {
        panic(err)
    }

    for _, match := range matches {
        if string(match[0]) != "." {
            foundFiles = append(foundFiles, match)
        }
    }

    return foundFiles

}

func GetImagesFromShare(user obj.User) []*smb2.File {
    _ = godotenv.Load()

    conn, err := net.Dial("tcp", os.Getenv("SMB_SHARE_HOST"))
    if err != nil {
        panic("SMB Share Connection Failed")
    }
    defer conn.Close()

    d := &smb2.Dialer{
        Initiator: &smb2.NTLMInitiator{
            User: os.Getenv("SMB_USER"),
            Password: os.Getenv("SMB_PASS"),
        },
    }

    s, err := d.Dial(conn)
    if err != nil {
        panic(err)
    }
    defer s.Logoff()
    
    if !ValidateUserDatashare(user) {
        return nil
    }
    
    foundFiles := make([]*smb2.File,0)
    fs, err := s.Mount(user.DataShare)
    if err != nil {
        panic(err)
    }
    defer fs.Umount()

    matches, err := iofs.Glob(fs.DirFS("."), "*")
    if err != nil {
        panic(err)
    }

    for _, match := range matches {
        if string(match[0]) != "." {
            file, _ := fs.Open(match)
            foundFiles = append(foundFiles, file) 
        }
    }

    return foundFiles
}

func LoadImageGETHandler(c *gin.Context) {
    val := session.GetSessionValue(c.Request, "User")
    user := val.(*obj.User)
    fmt.Println(user)
    filesList := listImagesFromShare(*user)

    c.HTML(http.StatusOK, "filelist.html", gin.H{"file_list": filesList})

}
