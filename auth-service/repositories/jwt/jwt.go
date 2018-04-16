package jwt

import (
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

const (
	refreshTokenExpireTime = time.Hour * time.Duration(24)
	authTokenExpireTime    = time.Minute * time.Duration(4)
)

//AuthClaims claims for auth jwt token
type AuthClaims struct {
	jwtGo.StandardClaims
	UserInfo
}

//RefreshClaims claims for refresh jwt token
type RefreshClaims struct {
	jwtGo.StandardClaims
	UUID string
}

var authKey, refreshKey []byte

func init() {
	authKey = []byte("q4t7w!z%C*F-JaNdRgUkXn2r5u8x/A?D")
	refreshKey = []byte("E(H+MbQeThWmZq4t7w!z%C*F-J@NcRfU")
}

//NewRefreshClaims generate new refresh claims
func NewRefreshClaims(id string) *RefreshClaims {
	expire := time.Now().Local().Add(refreshTokenExpireTime).Unix()

	return &RefreshClaims{
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "www.progsys.no",
		},
		UUID: id,
	}
}

//NewAuthClaims generate new refresh claims
func NewAuthClaims(userInfo *UserInfo) *AuthClaims {
	expire := time.Now().Local().Add(authTokenExpireTime).Unix()

	return &AuthClaims{
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "www.progsys.no",
		},
		UserInfo: *userInfo,
	}
}

//GetExpireAt get expire date of refresh claims
func (jc *RefreshClaims) GetExpireAt() time.Time {
	return time.Unix(jc.ExpiresAt, 0)
}

//GenerateTokenString generate a token string from refresh claims
func (jc *RefreshClaims) GenerateTokenString() (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jc)
	tokenString, err := token.SignedString(refreshKey)

	return tokenString, err
}

//GenerateTokenString generate a token string from auth claims
func (jc *AuthClaims) GenerateTokenString() (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jc)
	tokenString, err := token.SignedString(authKey)

	return tokenString, err
}

//ValidateRefreshToken validates refresh token and returns claims if ok
func ValidateRefreshToken(tokenString string) (*RefreshClaims, bool) {
	token, err := jwtGo.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if claims, ok := token.Claims.(*RefreshClaims); token.Valid && ok && err != nil {
		return claims, true
	}

	return nil, false
}

//ValidateAuthToken validates auth token and returns claims if ok
func ValidateAuthToken(tokenString string) (*AuthClaims, bool) {
	token, err := jwtGo.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		return authKey, nil
	})

	if claims, ok := token.Claims.(*AuthClaims); token.Valid && ok && err != nil {
		return claims, true
	}

	return nil, false
}
