package common

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func UseJWTAuth(engine *gin.Engine) {
	engine.Use(jwtAuthMiddleware())
}

// jwtAuthMiddleware 验证JWT Token
// @Summary 验证JWT Token
// @Description 验证JWT Token
// @Tags 通用
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} GetListRes
// @Router /role/list [get]
// @Security BearerToken
func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if skipAuth(c) {
			c.Next()
			return
		}
		// 从Header提取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供Token", "code": 100})
			return
		}

		// 校验格式: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误", "code": 100})
			return
		}
		tokenString := parts[1]

		// 解析并验证Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效Token"})
			return
		}

		// 提取Claims并存入上下文
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("username", claims["username"])
		}
		c.Next()
	}
}

func skipAuth(c *gin.Context) bool {
	if c.Request.URL.Path == "/user/regist" {
		return true
	}
	if c.Request.URL.Path == "/user/login" {
		return true
	}
	return false
}
