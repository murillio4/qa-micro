package jwt

import (
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

const (
	// RefreshExpireTime time to
	refreshExpireTime = time.Hour * time.Duration(24)
	// AuthExpireTime
	authExpireTime = time.Minute * time.Duration(4)
)

// Repository interface for token repo
type Repository interface {
	GetAuthClaims(UserInfo) *BaseClaims
	GetRefreshClaims(id string) *BaseClaims
	GenerateAuthToken(*BaseClaims) (string, error)
	GenerateRefreshToken(*BaseClaims) (string, error)
	ValidateAuthToken(string) (*BaseClaims, bool)
	ValidateRefreshToken(string) (*BaseClaims, bool)
}

// TokenRepository state holder for tokenrepo
type TokenRepository struct {
	authKey, refreshKey []byte
}

// NewTokenRepository create a new token repo
func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		[]byte("q4t7w!z%C*F-JaNdRgUkXn2r5u8x/A?D"),
		[]byte("E(H+MbQeThWmZq4t7w!z%C*F-J@NcRfU"),
	}
}

//GetAuthClaims generate new auth claims
func (*TokenRepository) GetAuthClaims(user UserInfo) *BaseClaims {
	expire := time.Now().Local().Add(authExpireTime).Unix()

	return &BaseClaims{
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "www.progsys.no",
		},
		UserInfo: user,
	}
}

//GetRefreshClaims generate new refresh claims
func (*TokenRepository) GetRefreshClaims(id string) *BaseClaims {
	expire := time.Now().Local().Add(refreshExpireTime).Unix()

	return &BaseClaims{
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "www.progsys.no",
			Id:        id,
		},
	}
}

//GenerateAuthToken generate a token string from auth claims
func (repo *TokenRepository) GenerateAuthToken(claims *BaseClaims) (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(repo.authKey)

	return tokenString, err
}

//GenerateRefreshToken generate a token string from refresh claims
func (repo *TokenRepository) GenerateRefreshToken(claims *BaseClaims) (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(repo.refreshKey)

	return tokenString, err
}

//ValidateAuthToken validates auth token and returns claims if ok
func (repo *TokenRepository) ValidateAuthToken(tokenString string) (*BaseClaims, bool) {
	token, err := jwtGo.ParseWithClaims(tokenString, &BaseClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		return repo.authKey, nil
	})

	if claims, ok := token.Claims.(*BaseClaims); token.Valid && ok && err != nil {
		return claims, true
	}

	return nil, false
}

//ValidateRefreshToken validates refresh token and returns claims if ok
func (repo *TokenRepository) ValidateRefreshToken(tokenString string) (*BaseClaims, bool) {
	token, err := jwtGo.ParseWithClaims(tokenString, &BaseClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		return repo.refreshKey, nil
	})

	if claims, ok := token.Claims.(*BaseClaims); token.Valid && ok && err != nil {
		return claims, true
	}

	return nil, false
}
