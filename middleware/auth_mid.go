package middleware

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	validator "github.com/wagslane/go-password-validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)    

func require_auth(c *gin.Context) gin.HandlerFunc {
    return func(c *gin.Context) {
        store := sessions.Session
        session, err := store.Get(c.Request, "session")
        if err != nil | session.Values["User"] == nil {
            gin.HTML(http.StatusOK, "index.html", gin.H{"login_required": "Requested page requires login"})
        }
    }

}
