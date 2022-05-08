package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
	"github.com/pathak107/cloudesk/pkg/api/service"
)

// AuthorizeUser validates the token from the http request, returning a 401 if it's not valid
func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		//Check whether there is a token or not
		if authHeader == "" || len(BEARER_SCHEMA) > len(authHeader) {
			c.Error(apierrors.NewUnauthorizedError(errors.New("token not provided"), "jwt_middleware"))
			c.Abort()
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.NewDefaultAuthService().ValidateToken(tokenString)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			c.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("user_data", claims["name"])
			fmt.Println(claims["name"])
			fmt.Println(claims["user_type"])
			fmt.Println(claims["user_id"])
			fmt.Println(claims["email"])
		} else {
			c.Error(apierrors.NewUnauthorizedError(err, "jwt_middleware"))
			c.Abort()
			return
		}
	}
}
