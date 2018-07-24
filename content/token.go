package content

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vexth1/handler"
)

const PrivateTokenKey = "123your_token_key"

type Token struct {
	Token string
}

// 获取token值
func GetToken(ctx *handler.Context) error {
	res := ctx.Response()
	res.WriteHeader(http.StatusOK)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &jwt.StandardClaims{}
	// private key generated with http://kjur.github.io/jsjws/tool_jwt.html
	s, e := token.SignedString([]byte(PrivateTokenKey))
	if e != nil {
		fmt.Println(e)
	}
	// res.Write([]byte(s))

	response := Token{s}
	json, _ := json.Marshal(response)

	res.Write(json)

	return nil
}
