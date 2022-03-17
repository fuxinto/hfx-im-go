package jwtx

import "github.com/golang-jwt/jwt"

/*
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
iat: jwt的签发时间
iss: jwt签发者
nbf: 定义在什么时间之前，该jwt都是不可用的.就是这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
sub: jwt所面向的用户
*/
func GetToken(secretKey, uid string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
