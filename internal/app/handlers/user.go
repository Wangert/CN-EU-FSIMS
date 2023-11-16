package handlers

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/service"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func Register(c *gin.Context) {
	glog.Info("################## FSIMS User Register ##################")

	var reqUser request.ReqUser
	if err := c.ShouldBind(&reqUser); err != nil || !checkRegisterParams(&reqUser) {
		glog.Errorln("User registration error")
		response.MakeFail(c, http.StatusNotAcceptable, "user registration failure!")
		return
	}

	// insert database
	glog.Info("request user parameters:")
	glog.Info(reqUser)
	err := service.AddFsimsUser(&reqUser)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("add new fsims user successful!")
	response.MakeSuccess(c, http.StatusOK, "successfully register the user!")
	return
}

func checkRegisterParams(reqUser *request.ReqUser) bool {
	if reqUser.Name == "" || reqUser.Account == "" || reqUser.Type == 0 || reqUser.Role == "" {
		glog.Errorln("Missing user registration parameters")
		return false
	}
	return true
}

func Login(c *gin.Context) {
	glog.Info("################## FSIMS User Login ##################")

	var reqLogin request.ReqLogin
	if err := c.ShouldBind(&reqLogin); err != nil || !checkLoginParams(&reqLogin) {
		response.MakeFail(c, http.StatusBadRequest, "login parameters error!")
		return
	}

	reqPwdHash := crypto.CalculateSHA256(reqLogin.Password, service.PASSWORD_SALT)
	uuid, pwdHash, err := service.QueryFsimsUserUuidAndPwdHash(reqLogin.Account)
	if err != nil {
		glog.Errorln("query fsims password hash error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	// check password
	if reqPwdHash != pwdHash {
		glog.Errorln("password incorrect!")
		response.MakeFail(c, http.StatusBadRequest, "password incorrecrt!")
		return
	}

	// create jwt token with uuid
	token, err := service.CreateJwtToken(uuid)
	if err != nil {
		glog.Errorln("create jwt token error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Infoln(reqLogin.Account, " login successful!")
	glog.Infoln("token:", token)

	resLogin := response.ResLogin{
		UUID:  uuid,
		Token: token,
	}
	response.MakeSuccess(c, http.StatusOK, resLogin)
}

func checkLoginParams(reqLogin *request.ReqLogin) bool {
	if reqLogin.Account == "" || reqLogin.Password == "" || reqLogin.Type == 0 {
		glog.Errorln("Missing login parameters")
		return false
	}
	return true
}

func GetAllUsers(c *gin.Context) {
	glog.Info("################## Get All FSIMS User ##################")

	users, err := query.FSIMSUser.WithContext(context.Background()).Find()
	if err != nil {
		glog.Errorln("query all fsims users error!")
		response.MakeFail(c, http.StatusBadRequest, "query all fsims user error")
		return
	}

	// user array to ReqUser array
	var resUsers = make([]response.ResUser, len(users))
	for index, user := range users {
		//convert
		resUsers[index] = models.FsimsUserToResUser(user)
	}

	glog.Info("query all fsims users successful")
	response.MakeSuccess(c, http.StatusOK, resUsers)
	return

}

//func Login(c *gin.Context) {
//	log.Println("******************* UserLogin *********************")
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	usertype, _ := strconv.Atoi(c.PostForm("usertype"))
//	res := mysql.QueryUser(username, password, usertype)
//	if res.Login {
//		log.Println("login success")
//	} else {
//		log.Println("login failed")
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"result":  res.Login,
//		"message": "登录请求发送成功！",
//		"data":    res,
//	})
//}

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

//func PlaceRisk(c *gin.Context) {
//	fmt.Println("******************* PlaceRisk *********************")
//	place := c.PostForm("Place")
//	ProductInfo := c.PostForm("ProductInfo")
//	Risk := c.PostForm("Risk")
//	Result, _ := strconv.Atoi(c.PostForm("Result"))
//	time := c.PostForm("Time")
//	mysql.InsertPlaceRiskNotification(Result, place, ProductInfo, Risk, time)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "请求发送成功！",
//		"result":  true,
//	})
//}
