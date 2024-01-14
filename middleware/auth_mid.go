package middleware

import (
    "slices"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/aevanra/mictakepic/pkg/session"
    "github.com/aevanra/mictakepic/pkg/objects"
)    

func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        val := session.GetSessionValue(c.Request, "User")

        if val == nil {
            c.HTML(http.StatusOK, "login.html", gin.H{"login_required": "Requested page requires login"})
            c.Abort()
            return
        }

        c.Next()
    }
}

func ConfirmShareAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        val := session.GetSessionValue(c.Request, "User")

        if val == nil {
            c.Status(401)
            c.Abort()
        }
        foundUser := val.(*obj.User)

        if !slices.Contains(foundUser.AllDatashares, c.Query("share")) {
            c.Status(401)
            c.Abort()
        }

        c.Next()
    }
}
