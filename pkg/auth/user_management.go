package auth

import (
	"context"
	"net/http"
	"os"
	"time"


	"github.com/aevanra/mictakepic/pkg/filesharing"
	"github.com/aevanra/mictakepic/pkg/objects"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)    

func CreateNewUser(c *gin.Context) {
    newUsername := c.PostForm("username")
    newPassword := c.PostForm("password")
    newDatashare := c.PostForm("datashare")
    isAdmin := c.PostForm("isAdmin")

    if !smb.ValidateUserDatashare(obj.User{DataShare: newDatashare}) {
        c.HTML(http.StatusBadRequest, "createUser.html", 
        gin.H{"message": "Provided DataShare does not exist."})
        return
    }
    
    if newUsername == "" || newPassword == "" || newDatashare == "" {
        c.HTML(http.StatusBadRequest, "createUser.html", 
        gin.H{"message": "New Users must have a username, a password, and a default datashare"})
        return
    }

    err := addUser(newUsername, newPassword, newDatashare, isAdmin)
    if err != nil {
        c.HTML(http.StatusNotFound, "createUser.html", 
        gin.H{"message": "Something went wrong -- user was not created or updated"})
        return
    }

    c.HTML(http.StatusOK, "createUser.html", gin.H{"message": "New user " + newUsername + " has been created or updated"})

}

func addUser(username string, password string, datashare string, isAdmin string) (err error) {
    //Load env file
    err = godotenv.Load()
    if err != nil {
        return err
    }
    
    admin := false  
    if isAdmin == "on" {
        admin = true
    }

    // Creating Mongo Connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    if err != nil {
     return err 
    }
    userCollection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USERS_COLLECTION"))

    ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    filter := bson.D{{Key: "Username", Value: username}}
    update := bson.D{{Key: "$set", 
    Value: bson.D{
        {Key: "ID", Value: uuid.New()}, 
        {Key: "Username", Value: username}, 
        {Key: "PassHash", Value: hash},
        {Key: "DataShare", Value: datashare},
        {Key: "Admin", Value: admin},
    }}}
    opts := options.Update().SetUpsert(true)

    _, err = userCollection.UpdateOne(ctx, filter, update, opts)

    if err != nil {
        return err
    }

    return nil
}
