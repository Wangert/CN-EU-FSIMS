package models

type FoodChainNode struct {
	FoodChainProcess string
	ProductInfo      string
	Description      string
	UploadCompany    string
	UploadPerson     string
	UploadTime       string
	FileHash         string
	SafetyResult     string
	SafetyReason     string
}

func (s *FoodChainNode) String() string {
	return "{ \"FoodChainProcess\": \"" + s.FoodChainProcess +
		" \", \"ProductInfo\": \"" + s.ProductInfo +
		" \", \"Description\": \"" + s.Description +
		" \", \"UploadCompany\": \"" + s.UploadCompany +
		" \", \"UploadPerson\": \"" + s.UploadPerson +
		" \", \"UploadTime\": \"" + s.UploadTime +
		" \", \"FileHash\": \"" + s.FileHash +
		" \", \"SafetyResult\": \"" + s.SafetyResult + "\" }"
}

type FoodChainBody struct {
	ID        string          `bson:"_id"`
	Signature string          `bson:"Signature"`
	FoodChain []FoodChainNode `bson:"FoodChain"`
	NowStep   string          `bson:"NowStep"`
}

func (s *FoodChainBody) String() string {
	return "{hash:" + s.Signature + " " + "FoodChainNode:" + s.FoodChain[0].String() + "}"
}
