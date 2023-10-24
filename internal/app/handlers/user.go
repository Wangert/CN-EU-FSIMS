package handlers

import (
	"CN-EU-FSIMS/database/mysql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	log.Println("***************UserRegister******************")
	username := c.PostForm("username")
	password := c.PostForm("password")
	company := c.PostForm("company")
	telephone := c.PostForm("telephone")
	place := c.PostForm("place")
	usertype, _ := strconv.Atoi(c.PostForm("usertype"))
	res := mysql.InsertUser(username, password, usertype, company, telephone, place)
	if res {
		log.Println("insert success")
	} else {
		log.Println("insert failed")
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  res,
		"message": "注册请求发送成功！",
	})
}

func Login(c *gin.Context) {
	log.Println("******************* UserLogin *********************")
	username := c.PostForm("username")
	password := c.PostForm("password")
	usertype, _ := strconv.Atoi(c.PostForm("usertype"))
	res := mysql.QueryUser(username, password, usertype)
	if res.Login {
		log.Println("login success")
	} else {
		log.Println("login failed")
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  res.Login,
		"message": "登录请求发送成功！",
		"data":    res,
	})
}

func UserNotification(c *gin.Context) {
	fmt.Println("******************* UserNotification *********************")
	username := c.Query("UserName")
	fmt.Println("User username", username)
	res := mysql.GetNotification(username)
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "请求发送成功！",
		"result":  true,
	})
}
func ReadNotification(c *gin.Context) {
	fmt.Println("******************* ReadNotification *********************")
	user := c.Query("user")
	mysql.ClearWarnNum(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求发送成功！",
		"result":  true,
	})
}

func PlaceRisk(c *gin.Context) {
	fmt.Println("******************* PlaceRisk *********************")
	place := c.PostForm("Place")
	ProductInfo := c.PostForm("ProductInfo")
	Risk := c.PostForm("Risk")
	Result, _ := strconv.Atoi(c.PostForm("Result"))
	time := c.PostForm("Time")
	mysql.InsertPlaceRiskNotification(Result, place, ProductInfo, Risk, time)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求发送成功！",
		"result":  true,
	})
}
