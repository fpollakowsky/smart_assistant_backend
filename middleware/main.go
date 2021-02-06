package middleware

import (
	"github.com/gin-gonic/gin"
	"shome-backend/mysql"
)

type APIKeyMiddleware struct {
	APIToken map[string]string
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"err": message})
}

func (nethcon *APIKeyMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := c.Request.Header.Get("X-Session-Token")

		if _token == "" {
			respondWithError(c, 401, "API _token required")
			return
		}

		if _, _found := nethcon.APIToken[_token]; _found {
			c.Next()
		} else {
			respondWithError(c, 401, "Invalid API _token")
			return
		}
	}
}

func (nethcon *APIKeyMiddleware) Populate() {
	_keys := mysql.GetApiKeys()

	for i := range _keys {
		nethcon.APIToken[_keys[0][i]] = _keys[1][i]
	}
}
