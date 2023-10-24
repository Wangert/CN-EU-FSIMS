package mongo

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/service"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var mongodb MongoDB

func initAndConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.176.35.60:27017"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("get database client success")
	}
	mongodb.client = client

	collection := client.Database("foodchaindb").Collection("foodchaindata")
	mongodb.collection = collection
	if collection != nil {
		log.Println("database connect success")
	}
}

func disconnect() {
	client := mongodb.client
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("disconnect database success")
	}
}

func SelectAllResult() []models.FoodChainBody {
	fmt.Println("**************SelectAllResult****************")

	initAndConnect()
	collection := mongodb.collection
	findOptions := options.Find()

	// 定义一个切片用来存储查询结果
	var results []models.FoodChainBody
	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem models.FoodChainBody
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	fmt.Println("Found multiple documents (array of pointers):", results)

	return results
}

func SelectChainByCompany(company string) []models.FoodChainBody {
	fmt.Println("***************** SelectChainByCompany *****************")
	var results []models.FoodChainBody

	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	findOptions := options.Find()

	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"Produce", company}},
				bson.D{{"Process", company}},
				bson.D{{"Storage", company}},
				bson.D{{"Retail", company}},
			}},
	}

	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem models.FoodChainBody
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	fmt.Println("Found multiple documents (array of pointers):", results)

	return results

}

func SelectChainByHash(hash string) *models.FoodChainBody {
	fmt.Println("**************SelectChainByHash****************")

	//result := model.FoodChainNode{}
	result := new(models.FoodChainBody)

	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	err := collection.FindOne(context.TODO(), bson.D{{"Signature", hash}}).Decode(result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Println("res: ", result)
	}

	return result
}

func SelectAllStepAndCountHazard() (int, int, int, int, int, int, int) {
	log.Println("***************** SelectAllStep *****************")
	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	filterProduce := bson.D{{"NowStep", "生产"}}
	filterProcess := bson.D{{"NowStep", "加工"}}
	filterPackage := bson.D{{"NowStep", "装储"}}
	filterRetail := bson.D{{"NowStep", "零售"}}
	countProduce, err := collection.CountDocuments(context.TODO(), filterProduce)
	if err != nil {
		panic(err)
	}
	countProcess, err := collection.CountDocuments(context.TODO(), filterProcess)
	if err != nil {
		panic(err)
	}
	countPackage, err := collection.CountDocuments(context.TODO(), filterPackage)
	if err != nil {
		panic(err)
	}
	countRetail, err := collection.CountDocuments(context.TODO(), filterRetail)
	if err != nil {
		panic(err)
	}
	log.Println("处于生产环节的食品链数量：", countProduce)
	log.Println("处于加工环节的食品链数量：", countProcess)
	log.Println("处于装储环节的食品链数量：", countPackage)
	log.Println("处于零售环节的食品链数量：", countRetail)

	allResults := SelectAllResult()
	TemperatureCount := 0
	ACount := 0
	BCount := 0
	for i := range allResults {
		Chain := allResults[i].FoodChain
		for i2 := range Chain {
			if Chain[i2].SafetyResult == "中" {
				ACount++
			} else if Chain[i2].SafetyResult == "高" {
				TemperatureCount++
			} else if Chain[i2].SafetyResult == "低" {
				BCount++
			}
		}
	}
	log.Println("ACount: ", ACount)
	log.Println("TemperatureCount: ", TemperatureCount)
	log.Println("BCount: ", BCount)

	return int(countProduce), int(countProcess), int(countPackage), int(countRetail), TemperatureCount, ACount, BCount
}

func CountHazard() (int, int, int) {
	log.Println("***************** CountHazard *****************")
	initAndConnect()
	defer disconnect()

	allResults := SelectAllResult()
	TemperatureCount := 0
	ACount := 0
	BCount := 0
	for i := range allResults {
		Chain := allResults[i].FoodChain
		for i2 := range Chain {
			if Chain[i2].SafetyResult == "中" {
				ACount++
			} else if Chain[i2].SafetyResult == "高" {
				TemperatureCount++
			} else if Chain[i2].SafetyResult == "低" {
				BCount++
			}
		}
	}
	log.Println("ACount: ", ACount)
	log.Println("TemperatureCount: ", TemperatureCount)
	log.Println("BCount: ", BCount)

	return TemperatureCount, ACount, BCount
}

func UploadFoodChainNode(FoodChainProcess string, ProductInfo string, FileHash string, UploadCompany string, UploadPerson string, UploadTime string, Description string, Signature string, SafetyResult string, SafetyReason string) *mongo.InsertOneResult {
	log.Println("***************** UploadFoodChainNode *****************")

	FoodChainNodeTemp := new(models.FoodChainNode)
	FoodChainNodeTemp.FoodChainProcess = FoodChainProcess
	FoodChainNodeTemp.ProductInfo = ProductInfo
	FoodChainNodeTemp.UploadCompany = UploadCompany
	FoodChainNodeTemp.UploadPerson = UploadPerson
	FoodChainNodeTemp.UploadTime = UploadTime
	FoodChainNodeTemp.Description = Description
	FoodChainNodeTemp.FileHash = FileHash
	FoodChainNodeTemp.SafetyResult = SafetyResult
	FoodChainNodeTemp.SafetyReason = SafetyReason

	FoodChain := []models.FoodChainNode{*FoodChainNodeTemp}
	FoodChainBody := models.FoodChainBody{Signature: Signature, FoodChain: FoodChain}
	fmt.Println("FoodChainBody:", FoodChainBody.FoodChain[0].String())

	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	res, err := collection.InsertOne(context.TODO(), bson.D{{"Signature", FoodChainBody.Signature}, {"FoodChain", FoodChainBody.FoodChain}, {"Produce", UploadCompany}, {"NowStep", FoodChainProcess}})
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Insert document: ", res.InsertedID)

	return res
}

func UpdateFoodChain(FoodChainProcess string, ProductInfo string, FileHash string, UploadCompany string, UploadPerson string, UploadTime string, Description string, LastSignature string, SafetyResult string, ResultReason string) (*mongo.UpdateResult, string, string) {
	log.Println("*****************UpdateFoodChainNode*****************")

	FoodChainNodeTemp := new(models.FoodChainNode)
	FoodChainNodeTemp.FoodChainProcess = FoodChainProcess
	FoodChainNodeTemp.ProductInfo = ProductInfo
	FoodChainNodeTemp.UploadCompany = UploadCompany
	FoodChainNodeTemp.UploadPerson = UploadPerson
	FoodChainNodeTemp.UploadTime = UploadTime
	FoodChainNodeTemp.Description = Description
	FoodChainNodeTemp.FileHash = FileHash
	FoodChainNodeTemp.SafetyResult = SafetyResult

	result := new(models.FoodChainBody)

	temp := "{UploadCompany:" + UploadCompany + " " + "FileHash:" + FileHash + " ProductInfo:" + ProductInfo + "}"
	ret := md5.Sum([]byte(temp))
	FoodChainID := hex.EncodeToString(ret[:])
	fmt.Println("FoodChainID: ", FoodChainID)
	fmt.Println("LastSignature: ", LastSignature)

	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	err := collection.FindOne(context.TODO(), bson.D{{"Signature", LastSignature}}).Decode(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("find one result: ", result)

	FoodChainTest := append(result.FoodChain, *FoodChainNodeTemp)

	//FoodChain := []model.FoodChainNode{*FoodChainNodeTemp}
	FoodChainBody := models.FoodChainBody{Signature: FoodChainID, FoodChain: FoodChainTest}
	fmt.Println("FoodChainBody:", FoodChainBody.FoodChain[0].String())

	selector := bson.M{"Signature": LastSignature}

	ID, _ := primitive.ObjectIDFromHex(result.ID)
	switch FoodChainProcess {
	case "加工":
		fmt.Println("process id: ", ID)
		data := bson.M{"$set": bson.D{{"Process", UploadCompany}}}
		UpsertRes, err := collection.UpdateOne(context.TODO(), selector, data)
		if err != nil {
			fmt.Println("Upsert Err:", err)
		} else {
			fmt.Println("Process UpsertRes: ", UpsertRes)
		}
		break
	case "储藏/包装":
		fmt.Println("storage id: ", ID)
		data := bson.M{"$set": bson.D{{"Storage", UploadCompany}}}
		UpsertRes, err := collection.UpdateOne(context.TODO(), selector, data)
		if err != nil {
			fmt.Println("Upsert Err:", err)
		} else {
			fmt.Println("Storage UpsertRes: ", UpsertRes)
		}
		break
	case "零售":
		fmt.Println("Retail id: ", ID)
		data := bson.M{"$set": bson.D{{"Retail", UploadCompany}}}
		UpsertRes, err := collection.UpdateOne(context.TODO(), selector, data)
		if err != nil {
			fmt.Println("Upsert Err:", err)
		} else {
			fmt.Println("Retail UpsertRes: ", UpsertRes)
		}
		break
	}

	res, err := collection.UpdateOne(context.TODO(), selector, bson.M{"$set": bson.D{{"Signature", FoodChainBody.Signature}, {"FoodChain", FoodChainTest}, {"NowStep", FoodChainProcess}}})
	if err != nil {
		log.Print(err)
	}

	for i, v := range result.FoodChain {
		fmt.Println("result.FoodChain item: ", result.FoodChain[i])
		fmt.Println("item UploadCompany: ", v.UploadCompany)
		sendToOtherCompany := service.SafetyResultNotification(SafetyResult, ResultReason, UploadTime, result.ID, UploadCompany, v.UploadCompany, *FoodChainNodeTemp)
		if sendToOtherCompany {
			fmt.Println("send to other company success")
		}
	}

	return res, result.ID, FoodChainID
}

func DeleteBySingature(signature string) {
	fmt.Println("**************DeleteBySingature****************")

	initAndConnect()
	defer disconnect()

	collection := mongodb.collection
	// 删除名字是小黄的那个
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"Signature", signature}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

}
