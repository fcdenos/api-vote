package middlware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

func NewJWT() *jwt.GinJWTMiddleware {
	var identityKey string = "uuid"
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{

		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "identityKey",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"uuid": v.UUID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			var u model.User
			u.UUID = claims[identityKey].(string)
			return &u
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var u model.User
			u.UUID = c.GetString("user_uuid")
			u.Pass = c.GetString("user_pass")
			if u.PassIsValid(c.GetString("payload_pass")) {
				return &model.User{
					ModelDB: model.ModelDB{
						UUID: u.UUID,
					},
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			if int64(claims["exp"].(float64)) < time.Now().Unix() {
				return false
			}
			c.Set("jwt_user_uuid", claims["uuid"].(string))
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		panic(err)
	}
	return authMiddleware
}
