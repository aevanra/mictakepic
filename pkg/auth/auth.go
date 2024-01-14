package auth

import (
    "context"
    "net/http"
    "os"
    "time"

    "github.com/aevanra/mictakepic/pkg/objects"
    "github.com/aevanra/mictakepic/pkg/session"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func getUserByUsername(username string) (obj.User, error) {
    //Load env file
    err := godotenv.Load()
    if err != nil{
        return obj.User{}, err
    }

    // Creating Mongo Connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    if err != nil {
       panic(err) 
    }
    userCollection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USERS_COLLECTION"))

    filter := bson.D{{Key: "Username", Value: username}}
    var foundUser obj.User
    err = userCollection.FindOne(ctx, filter).Decode(&foundUser)

    if err != nil {
        return obj.User{}, err
    }

    return foundUser, nil
}


func LoginPOSTHandler(c *gin.Context) {
    username, pass := c.PostForm("username"), c.PostForm("password")

    if username == "" ||  pass == "" {
        c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Please enter a Username and Password"})
        return
    }

    foundUser, err := getUserByUsername(username)
    message := gin.H{"username": foundUser.Username, "adminStatus": foundUser.Admin, "Shares": foundUser.AllDatashares}
    
    if err != nil{
        c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid Username and/or Password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(foundUser.PassHash), []byte(pass))
    if err != nil {
        c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid Username and/or Password"})
        return
    }
    
    err = session.SetSessionValue(c.Request, c.Writer, "User", foundUser)
    if err != nil {
        panic(err)
    }

    c.HTML(http.StatusOK, "userpage.html", message)
}

func UserHomeHandler(c *gin.Context) {
    val := session.GetSessionValue(c.Request, "User")
    foundUser := val.(*obj.User)
    message := gin.H{"username": foundUser.Username, "Shares": foundUser.AllDatashares}
    
    if foundUser.Admin {
        message["adminStatus"] = foundUser.Admin
    } 

    c.HTML(http.StatusOK, "userpage.html", message)
}

