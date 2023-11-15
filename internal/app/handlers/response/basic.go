package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeSuccess return success response
func MakeSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "data": data})
}

//func MakeSuccessPage(c *gin.Context, code int, data interface{}, count int, pageIndex int, pageSize int, sl *model.Syslog) {
//	//put log into database
//	if sl != nil {
//		sl.Result = "成功"
//		err := model.PutSyslog(*sl)
//		if err != nil {
//			glog.Error(err, sl)
//		}
//	}
//	c.JSON(http.StatusOK, gin.H{"statusCode": code, "data": data, "count": count, "index": pageIndex, "size": pageSize})
//}

// MakeFail return fail response
func MakeFail(c *gin.Context, code int, message interface{}) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "message": message})
}
