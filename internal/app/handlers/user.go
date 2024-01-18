package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"CN-EU-FSIMS/utils/crypto"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func QueryPastureTrashFifteenDays(c *gin.Context) {
	var req request.ReqTrashFifteenDays
	if err := c.ShouldBind(&req); err != nil {
		glog.Errorln("Query trash error")
		response.MakeFail(c, http.StatusNotAcceptable, "query trash failure!")
		return
	}

	infos, count, err := service.QueryPastureTrashFifteenDays(&req)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query pasture trash disposal fifteen days error!")
		return
	}

	res := response.ResPastureTrashDisposalFifteenDays{
		Count:                  count,
		PastureDisposalRecords: infos,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QueryEndFeedCow(c *gin.Context) {
	glog.Info("###Total number of feed end cows")
	count, err := service.QueryEndFeedCow()
	if err != nil {
		glog.Errorln("Query feed end cows error!")
		response.MakeFail(c, http.StatusNotAcceptable, "query cow failure!")
	}
	response.MakeSuccess(c, http.StatusOK, count)
	return
}

func QueryEndBatches(c *gin.Context) {
	glog.Info("###Total number of end batches###")
	count, err := service.QueryBatchesEnd()
	if err != nil {
		glog.Errorln("Query number of batches error!")
		response.MakeFail(c, http.StatusNotAcceptable, "query batches failure!")
	}
	response.MakeSuccess(c, http.StatusOK, count)
	return
}
func QueryEndSlaCow(c *gin.Context) {
	glog.Info("###Total number of slaughter end cows###")
	count, err := service.QueryEndSlaCow()
	if err != nil {
		glog.Errorln("Query slaughter end cows error!")
		response.MakeFail(c, http.StatusNotAcceptable, "query cow failure!")
	}
	response.MakeSuccess(c, http.StatusOK, count)
	return
}
func QuerySlaughterFifteenDays(c *gin.Context) {
	var req request.ReqTrashFifteenDays
	if err := c.ShouldBind(&req); err != nil {
		glog.Errorln("Query trash error")
		response.MakeFail(c, http.StatusNotAcceptable, "query trash failure!")
		return
	}

	infos, count, err := service.QuerySlaughterTrashFifteenDays(&req)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter trash disposal fifteen days error!")
		return
	}

	res := response.ResSlaughterDisposalFifteenDays{
		Count:                    count,
		SlaughterDisposalRecords: infos,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
func QuerySlaAndPasTrashPerDay(c *gin.Context) {
	var r request.ReqTrashPerDay
	if err := c.ShouldBind(&r); err != nil {
		glog.Errorln("Query trash error")
		response.MakeFail(c, http.StatusNotAcceptable, "query trash failure!")
		return
	}

	fmt.Println(r.TimeStamp)
	info, err := service.QueryTrashPerDay(&r)
	if err != nil {
		glog.Errorln("Query trash error")
		response.MakeFail(c, http.StatusNotAcceptable, "query trash failure!")
		return
	}

	res := response.ResSlaAndPasTrashPeyDay{
		OdorSlaAndPasTrashPeyDay1:    info.TrashDisposalPerDayOdorInfo1,
		OdorSlaAndPasTrashPeyDay2:    info.TrashDisposalPerDayOdorInfo2,
		OdorSlaAndPasTrashPeyDay3:    info.TrashDisposalPerDayOdorInfo3,
		OdorSlaAndPasTrashPeyDay4:    info.TrashDisposalPerDayOdorInfo4,
		ResidueSlaAndPasTrashPeyDay1: info.TrashDisposalPerDayResidueInfo1,
		ResidueSlaAndPasTrashPeyDay2: info.TrashDisposalPerDayResidueInfo2,
		ResidueSlaAndPasTrashPeyDay3: info.TrashDisposalPerDayResidueInfo3,
		ResidueSlaAndPasTrashPeyDay4: info.TrashDisposalPerDayResidueInfo4,
		WaterSlaAndPasTrashPeyDay1:   info.TrashDisposalPerDayWaterInfo1,
		WaterSlaAndPasTrashPeyDay2:   info.TrashDisposalPerDayWaterInfo2,
		WaterSlaAndPasTrashPeyDay3:   info.TrashDisposalPerDayWaterInfo3,
		WaterSlaAndPasTrashPeyDay4:   info.TrashDisposalPerDayWaterInfo4,
	}
	//查询垃圾处理信息
	glog.Info("query trash disposal info successfully")
	response.MakeSuccess(c, http.StatusOK, res)
}

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
	ps := reqUser.Password
	if len(ps) < 9 {
		glog.Errorln("password len is < 9")
		return false
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		glog.Errorln("password need num :%v", err)
		return false
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		glog.Errorln("password need a_z :%v", err)
		return false
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		glog.Errorln("password need A_Z :%v", err)
		return false
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		glog.Errorln("password need symbol :%v", err)
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

	fmt.Println(reqLogin.Account)
	reqPwdHash := crypto.CalculateSHA256(reqLogin.Password, service.PASSWORD_SALT)
	uuid, pwdHash, housenumber, usertype, err := service.QueryFsimsUserUuidAndPwdHash(reqLogin.Account)
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
	glog.Infoln("type", usertype)
	resLogin := response.ResLogin{
		UUID:        uuid,
		Token:       token,
		UserType:    usertype,
		HouseNumber: housenumber,
	}
	response.MakeSuccess(c, http.StatusOK, resLogin)
}

func checkLoginParams(reqLogin *request.ReqLogin) bool {
	if reqLogin.Account == "" || reqLogin.Password == "" {
		glog.Errorln("Missing login parameters")
		return false
	}
	return true
}

func GetAllUsers(c *gin.Context) {
	glog.Info("################## Get All FSIMS User ##################")

	users, count, err := service.GetAllUsers()
	if err != nil {
		glog.Errorln("query all fsims users error!")
		response.MakeFail(c, http.StatusBadRequest, "query all fsims user error")
		return
	}
	glog.Info("query all fsims users successful")

	res := response.ResUsers{
		Users: users,
		Count: count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func ViewLogs(c *gin.Context) {
	fmt.Println("******************* View Logs *********************")

	logs, count, err := service.GetAllLogs()
	if err != nil {
		glog.Errorln("query all logs error!")
		response.MakeFail(c, http.StatusBadRequest, "query all logs error")
		return
	}

	glog.Info("query all fsims logs successful")

	res := response.ResLogs{
		Logs:  logs,
		Count: count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
}

func AddUserByAdmin(c *gin.Context) {
	glog.Info("################## Add A FSIMS User By Admin ##################")
	var u request.ReqAddUser
	if err := c.ShouldBind(&u); err != nil || !checkAddParams(&u) {
		response.MakeFail(c, http.StatusBadRequest, "Add a user parameters error!")
		return
	}
	//add
	glog.Info("request user add parameters:")
	glog.Info(u)
	err := service.AddFsimsUserByAdmin(&u)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "The new user information insert error!")
		return
	}
	glog.Info("fsims user add successful")
	response.MakeSuccess(c, http.StatusOK, "successfully add the user by admin")
}

func checkAddParams(reqAddUser *request.ReqAddUser) bool {
	if reqAddUser.Name == "" || reqAddUser.Account == "" || reqAddUser.Type == 0 || reqAddUser.Role == "" {
		glog.Errorln("Missing user registration parameters")
		return false
	}
	return true
}

func ResetPasswordByAdmin(c *gin.Context) {
	glog.Info("################## Reset User's Password By Admin ##################")
	var a request.ReqAccount
	if err := c.ShouldBind(&a); err != nil {
		response.MakeFail(c, http.StatusBadRequest, "Reset User's Password Error!")
		return
	}
	//reset
	glog.Info("request user account:")
	glog.Info(a)

	err := service.ResetFsimsPassWord(&a)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "reset the user's password error!")
		return
	}
	glog.Info("reset user's password successful!")
	response.MakeSuccess(c, http.StatusOK, "successfully update the user!")
	return
}

func DeleteUser(c *gin.Context) {
	glog.Info("################## Delete A FSIMS User ##################")
	var a request.ReqAccount
	if err := c.ShouldBind(&a); err != nil {
		response.MakeFail(c, http.StatusBadRequest, "Delete a user error")
		return
	}
	fmt.Println(a.Account)
	fmt.Println(a.Type)
	err := service.DeleteFsimUser(&a)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "The user information delete error!")
		return
	}
	glog.Info("delete fsims user successful")
	response.MakeSuccess(c, http.StatusOK, "successfully delete the user!")
}

//func UserNotification(c *gin.Context) {
//	fmt.Println("******************* UserNotification *********************")
//	username := c.Query("UserName")
//	fmt.Println("User username", username)
//	res := mysql.GetNotification(username)
//	c.JSON(http.StatusOK, gin.H{
//		"data":    res,
//		"message": "请求发送成功！",
//		"result":  true,
//	})
//}
//func ReadNotification(c *gin.Context) {
//	fmt.Println("******************* ReadNotification *********************")
//	user := c.Query("user")
//	mysql.ClearWarnNum(user)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "请求发送成功！",
//		"result":  true,
//	})
//}
