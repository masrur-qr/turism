package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)
var myKey = []byte("sekKey")
func Genertetoken()(string) {
	token := jwt.New(jwt.SigningMethodHS256)
	
	claim := token.Claims.(jwt.MapClaims)
	claim["authrized"] = true
	claim["usre"] = "mas,d"
	claim["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenStrng, err := token.SignedString(myKey)

	if err != nil{
		fmt.Printf(err.Error())
	}

	return tokenStrng
}
 