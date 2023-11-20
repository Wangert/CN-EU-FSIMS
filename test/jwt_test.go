package test

import (
	"CN-EU-FSIMS/internal/service"
	"fmt"
	"testing"
)

func TestCreateAndVerifyJwtToken(t *testing.T) {

	token, err := service.CreateJwtToken("FSIMSU-0001-50e0a1cbe932da438f2f8fc4689c1bc0df264560d969ad9a12265b73b3dab865")
	if err != nil {
		panic("create jwt token error")
	}

	fmt.Println("JWT Token:")
	fmt.Println(token)

	userClaims, err := service.ParseToken(token)
	if err != nil {
		panic("parse jwt token error")
	}

	fmt.Println("user claims:")
	fmt.Println(*userClaims)
}

func TestVerifyJwtToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiRlNJTVNVLTAwMDEtNTBlMGExY2JlOTMyZGE0MzhmMmY4ZmM0Njg5YzFiYzBkZjI2NDU2MGQ5NjlhZDlhMTIyNjViNzNiM2RhYjg2NSIsImV4cCI6MTcwMDAzMDA1OSwiaWF0IjoxNzAwMDMwMDU5LCJpc3MiOiJGU0lNUy1TWVNURU0ifQ.TUey7efIprQOZjLAUew1U_GLybRnbJOSM-ARAVMb2mE"

	fmt.Println(token)
	userClaims, err := service.ParseToken(token)
	if err != nil {
		panic("parse jwt token error")
	}

	fmt.Println("user claims:")
	fmt.Println(*userClaims)
}
