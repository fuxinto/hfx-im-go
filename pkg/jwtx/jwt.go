package jwtx

import (
	"HIMGo/pkg/response"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

const (
	Success                    = 200
	Error                      = 500
	Invalid_params             = 400
	Error_auth_token_invalid   = 1000
	Error_auth_token_Malformed = 1001
	Error_account_none         = 1002
	Error_account_exist        = 1003
	Error_password_wrong       = 1004
	Error_updateImageFail      = 1005
)

var msgFlags = map[int]string{
	Success:                    "成功",
	Error:                      "失败",
	Invalid_params:             "请求参数有误",
	Error_account_none:         "账号不存在",
	Error_account_exist:        "账号已存在",
	Error_password_wrong:       "密码错误",
	Error_auth_token_invalid:   "Token已失效,请重新登录",
	Error_auth_token_Malformed: "token效验失败",
	Error_updateImageFail:      "图片上传失败",
}

/*
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
iat: jwt的签发时间
iss: jwt签发者
nbf: 定义在什么时间之前，该jwt都是不可用的.就是这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
sub: jwt所面向的用户
*/
func GetToken(secretKey, uid string, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

//验证token
func GetUidWithToken(tokenString, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["uid"].(string), nil
	}
	return "", fmt.Errorf("token解析失败")
}
func SwitchTokenError(w http.ResponseWriter, err error) {
	var code = Success
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				code = Error_auth_token_Malformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				code = Error_auth_token_invalid
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				code = Error_auth_token_invalid
			} else {
				code = Error_auth_token_Malformed
			}
		}
		code = Error_auth_token_Malformed
	}
	var body response.Body
	body.Code = code
	body.Msg = msgFlags[code]
	httpx.OkJson(w, body)
}
