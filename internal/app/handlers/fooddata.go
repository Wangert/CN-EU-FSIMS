package handlers

import (
	"CN-EU-FSIMS/database/mongo"
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/service"
	"CN-EU-FSIMS/internal/service/analysis"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMongoDBData(c *gin.Context) {
	fmt.Println("***************** GetAllMongoDBData *****************")
	res := mongo.SelectAllResult()
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   res,
	})
}

func GetHazardStatus(c *gin.Context) {
	fmt.Println("***************** GetHazardStatus *****************")
	TCount, ACount, BCount := mongo.CountHazard()
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"TCount": TCount + 13,
		"ACount": ACount + 61,
		"BCount": BCount + 44,
	})
}

func GetCompanyFoodData(c *gin.Context) {
	fmt.Println("***************** GetCompanyFoodData *****************")
	company := c.Query("Company")
	res := mongo.SelectChainByCompany(company)
	fmt.Println("GetCompanyFoodData res: ", res)
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   res,
	})
}

func GetStepStatus(c *gin.Context) {
	fmt.Println("***************** GetStepStatus *****************")
	ProduceSum, ProcessSum, PackageSum, RetailSum, TCount, ACount, BCount := mongo.SelectAllStepAndCountHazard()
	c.JSON(http.StatusOK, gin.H{
		"result":     true,
		"ProduceSum": ProduceSum,
		"ProcessSum": ProcessSum,
		"PackageSum": PackageSum,
		"RetailSum":  RetailSum,
		"TCount":     TCount,
		"ACount":     ACount,
		"BCount":     BCount,
	})
}

// FoodDataUpload 将食品数据整理后调用mongodb接口函数进行上传
func FoodDataUpload(c *gin.Context) {
	log.Println("***************** FoodDataUpload *****************")
	FoodChainProcess := c.PostForm("FoodChainProcess")
	ProductInfo := c.PostForm("ProductInfo")
	Description := c.PostForm("Description")
	UploadCompany := c.PostForm("UploadCompany")
	UploadPerson := c.PostForm("UploadPerson")
	UploadTime := c.PostForm("UploadTime")
	UploadFile, _ := c.FormFile("UploadFile")

	// 计算文件hash
	filetemp, _ := UploadFile.Open()
	m := md5.New()
	_, _ = io.Copy(m, filetemp)
	hash := hex.EncodeToString(m.Sum(nil))
	fmt.Println("hash:", hash)

	FoodSafetyResult, ResultReason := analysis.FoodFileCheck(filetemp, FoodChainProcess)

	// 存储上传数据文件
	basePath := "./UploadFile/"
	filename := basePath + hash + ".csv"
	if err := c.SaveUploadedFile(UploadFile, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	if FoodChainProcess != "生产" {
		fmt.Println("非生产环节")
		LastSignature := c.PostForm("Signature")
		res, resID, signature := mongo.UpdateFoodChain(FoodChainProcess, ProductInfo, hash, UploadCompany, UploadPerson, UploadTime, Description, LastSignature, FoodSafetyResult, ResultReason)
		fmt.Println("Update target ID: ", resID)

		node := models.FoodChainNode{
			FoodChainProcess: FoodChainProcess,
			ProductInfo:      ProductInfo,
			Description:      Description,
			UploadCompany:    UploadCompany,
			UploadPerson:     UploadPerson,
			UploadTime:       UploadTime,
			FileHash:         hash,
			SafetyResult:     FoodSafetyResult,
		}
		sendNotRes := service.SafetyResultNotification(FoodSafetyResult, ResultReason, UploadTime, resID, UploadCompany, UploadPerson, node)
		sendToCompany := service.SafetyResultNotification(FoodSafetyResult, ResultReason, UploadTime, resID, UploadCompany, UploadCompany, node)
		if sendNotRes && sendToCompany {
			fmt.Println("通知发送成功")
		} else {
			fmt.Println("通知发送失败")
		}

		fabricres, err := fabric.UpdateFoodData(resID, signature, FoodChainProcess, hash, UploadCompany)
		fmt.Println("Update target ID: ", fabricres)
		if err == nil {
			fmt.Println("更新数据成功")
			c.JSON(http.StatusOK, gin.H{
				"result":    true,
				"data":      res,
				"Signature": GetSignature(node),
			})
		}
	} else {
		temp := "{UploadCompany:" + UploadCompany + " " + "FileHash:" + hash + " ProductInfo:" + ProductInfo + "}"
		ret := md5.Sum([]byte(temp))
		FoodChainID := hex.EncodeToString(ret[:])
		fmt.Println("FoodchainID: ", FoodChainID)

		res := mongo.UploadFoodChainNode(FoodChainProcess, ProductInfo, hash, UploadCompany, UploadPerson, UploadTime, Description, FoodChainID, FoodSafetyResult, ResultReason)
		resID := res.InsertedID.(primitive.ObjectID).Hex()

		node := models.FoodChainNode{
			FoodChainProcess: FoodChainProcess,
			ProductInfo:      ProductInfo,
			Description:      Description,
			UploadCompany:    UploadCompany,
			UploadPerson:     UploadPerson,
			UploadTime:       UploadTime,
			FileHash:         hash,
			SafetyResult:     FoodSafetyResult,
		}
		sendNotRes := service.SafetyResultNotification(FoodSafetyResult, ResultReason, UploadTime, resID, UploadCompany, UploadPerson, node)
		sendToCompany := service.SafetyResultNotification(FoodSafetyResult, ResultReason, UploadTime, resID, UploadCompany, UploadCompany, node)
		if sendNotRes && sendToCompany {
			fmt.Println("通知发送成功")
		} else {
			fmt.Println("通知发送失败")
		}

		fabricresult, err := fabric.UploadFoodData(resID, FoodChainID, FoodChainProcess, hash, UploadCompany)
		fmt.Println("fabric: ", fabricresult)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result":    true,
				"data":      res,
				"Signature": FoodChainID,
			})
		}
	}

}

func FoodDataQuery(c *gin.Context) {
	log.Println("***************** FoodDataQuery *****************")

	FoodID := c.Query("FoodID")
	log.Println("FoodId:", FoodID)

	res := mongo.SelectChainByHash(FoodID)
	fmt.Println("FoodDataQuery res: ", res)

	fabricres := fabric.GetFoodData(res.ID)
	fmt.Println("FabricFoodDataQuery", fabricres)

	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":     true,
			"Signature":  res.Signature,
			"FoodChain":  res.FoodChain,
			"BlockChain": fabricres,
		})
	}
}

func DeleteDB(c *gin.Context) {
	fmt.Println("***************** DeleteDB *****************")
	Signature := c.PostForm("Signature")
	mongo.DeleteBySingature(Signature)
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "删除数据库数据成功",
	})
}

func GetAllBlockchainData(c *gin.Context) {
	log.Println("***************** GetAllBlockchainData *****************")
	res := fabric.GetAllFabricResult()
	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": true,
			"data":   res,
		})
	}
}
func DeleteTest(c *gin.Context) {
	log.Println("***************** DeleteTest *****************")
	AssetID := c.PostForm("id")
	res := fabric.DeleteTest(AssetID)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": "删除数据成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "删除数据失败",
		})
	}
}

func FoodDataFileDownload(c *gin.Context) {
	fmt.Println("***************** FoodDataFileDownload *****************")
	FileHash := c.Query("FileHash")
	basePath := "./UploadFile/"
	FileName := basePath + FileHash + ".csv"
	fmt.Println("FileName:", FileName)

	c.File(FileName)
}

func GetSignature(node models.FoodChainNode) string {
	temp := "{UploadCompany:" + node.UploadCompany + " " + "FileHash:" + node.FileHash + " ProductInfo:" + node.ProductInfo + "}"
	ret := md5.Sum([]byte(temp))
	FoodChainID := hex.EncodeToString(ret[:])
	return FoodChainID
}
