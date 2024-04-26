package auth

import (
	"context"
    "io"
    "encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aevanra/mictakepic/pkg/objects"
	"github.com/aevanra/mictakepic/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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
    // Get the username and password from the request body
    postBodyBytes, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading request body"})
    }

    postBody := make(map[string]interface{})
    json.Unmarshal(postBodyBytes, &postBody)

    username := postBody["username"].(string)
    pass := postBody["password"].(string)

    if username == "" ||  pass == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Please enter a Username and Password"})
        return
    }

    foundUser, err := getUserByUsername(username)
    
    if err != nil{
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Username and/or Password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(foundUser.PassHash), []byte(pass))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Username and/or Password"})
        return
    }
    
    err = session.SetSessionValue(c.Request, c.Writer, "User", foundUser)
    if err != nil {
        panic(err)
    }

    c.Header("Access-Control-Allow-Origin", "*")
    c.JSON(http.StatusOK, foundUser)
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

