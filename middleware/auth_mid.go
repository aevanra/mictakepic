package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/aevanra/mictakepic/pkg/session"
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
