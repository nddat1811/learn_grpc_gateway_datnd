package auth

import (
	"gateway/demo"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)
var jwtKey = os.Getenv("API_KEY")

const (
	expRefreshToken = 24
	expToken        = 15
	company         = "Hybrid Technologies Viet Nam"
	hostMail        = "humghuy201280@gmail.com"
	subjectMail     = "Email reset password"
	textContent     = "Struction to reset your password:"
)

type AuthMiddlewareConfig struct {
	svc *demo.UnimplementedDemoGatewayServer
}

func InitAuthMiddleware(svc *demo.UnimplementedDemoGatewayServer) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (m *AuthMiddlewareConfig) GenerateToken(username string) (map[string]string, error) {
	//generate access token
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["id"] = "1111"
	claims["exp"] = time.Now().Add(time.Hour * expToken).Unix()
	token, err := tokenJwt.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": token,
	}, nil
}

// func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
// 	authorization := ctx.Request.Header.Get("authorization")

// 	if authorization == "" {
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	token := strings.Split(authorization, "Bearer ")

// 	if len(token) < 2 {
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
// 		Token: token[1],
// 	})

// 	if err != nil || res.Status != http.StatusOK {
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	ctx.Set("userId", res.UserId)

// 	ctx.Next()
// }
