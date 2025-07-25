package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"task-week-four/handlers/common"
	"task-week-four/models"
	"task-week-four/utils"
	"time"
)

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 模拟数据库验证（实际替换为DB查询）
	user, err := models.UserLogin(req.UserName, req.Password)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "用户登录失败，用户名和密码错误！",
		})
		return
	}

	// 生成JWT
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(common.JwtSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"token":   tokenString,
	})
}

func Regist(c *gin.Context) {
	req := AddUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  100,
		})
		return
	}
	row := req.ToUser()
	if err := models.UserRegister(row); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "注册用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户注册成功",
	})
}
