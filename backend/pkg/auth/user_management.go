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
    
    availableShares := filesharing.ListDataShareNames()
    accessShares := make([]string, 0)
    for _, share := range availableShares {
        access := c.PostForm("access" + share)
        if access == "on" {
            accessShares = append(accessShares,  share)
        }
    }

    if !filesharing.ValidateUserDatashare(obj.User{DefaultDataShare: newDatashare}) {
        c.HTML(http.StatusBadRequest, "createUser.html", 
        gin.H{"message": "Provided DataShare does not exist.", "Shares": availableShares})
        return
    }
    
    if newUsername == "" || newPassword == "" || newDatashare == "" {
        c.HTML(http.StatusBadRequest, "createUser.html", 
        gin.H{"message": "New Users must have a username, a password, and a default datashare", "Shares": availableShares})
        return
    }

    err := addUser(newUsername, newPassword, newDatashare, accessShares, isAdmin)
    if err != nil {
        c.HTML(http.StatusNotFound, "createUser.html", 
        gin.H{"message": "Something went wrong -- user was not created or updated", "Shares": availableShares})
        return
    }

    c.HTML(http.StatusOK, "createUser.html", gin.H{"message": "New user " + newUsername + " has been created or updated", "Shares": availableShares})

}

func addUser(username string, password string, datashare string, accessShares []string, isAdmin string) (err error) {
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
        {Key: "DefaultDataShare", Value: datashare},
        {Key: "AllDataShares", Value: accessShares},
        {Key: "Admin", Value: admin},
    }}}
    opts := options.Update().SetUpsert(true)

    _, err = userCollection.UpdateOne(ctx, filter, update, opts)

    if err != nil {
        return err
    }

    return nil
}
