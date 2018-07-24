package content

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vexth1/handler"
)

const PrivateTokenKey = "123your_token_key"

// Token 对象
type Token struct {
	Token string
}

// GetToken 获取token值
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

// Skipper 过滤一些不需要token验证的接口
func Skipper(path string) bool {
	if path == "/skipper" {
		return true
	}
	return false
}

// JwtErrHandler 返回是否带有token或者token失效的信息
func JwtErrHandler(w http.ResponseWriter, r *http.Request, err string) {
	fmt.Println(err)
	http.Error(w, err, 401)
}
