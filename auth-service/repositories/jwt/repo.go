package jwt

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// Repository permission repository interface
type Repository interface {
	NewJWTRepository(*redis.Client) error
	Create(*UserInfo) (string, string, error)
	CheckAuthToken(string) (*UserInfo, bool)
	CheckRefreshToken(string) (string, bool)
	DeleteRefreshToken(string) error
}

// TokenRepository Repo state
type TokenRepository struct {
	redisClient *redis.Client
}

// NewTokenRepository new token repo
func NewTokenRepository(redisClient *redis.Client) *TokenRepository {
	return &TokenRepository{redisClient: redisClient}
}

//Create create refresh and auth tokens
func (*TokenRepository) Create(userInfo *UserInfo) (string, string, error) {
	var (
		authTokeString, refreshTokenString string
		err                                error
	)

	refreshClaims := NewRefreshClaims(userInfo.GetID())
	if refreshTokenString, err = refreshClaims.GenerateTokenString(); err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"token": refreshTokenString,
		}).Error("Error while generating refresh token string")

		return "", "", err
	}

	authClaims := NewAuthClaims(userInfo)
	if authTokeString, err = authClaims.GenerateTokenString(); err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"token": authTokeString,
		}).Error("Error while generating auth token sring")

		return "", "", err
	}
	return authTokeString, refreshTokenString, nil
}

//CheckAuthToken check if a auth token is valid and retrieve it
func (*TokenRepository) CheckAuthToken(tokenString string) (*UserInfo, bool) {
	if claims, ok := ValidateAuthToken(tokenString); ok {
		return &claims.UserInfo, true
	}
	return nil, false
}

//CheckRefreshToken checks the refresh token and populates t if valid
func (repo *TokenRepository) CheckRefreshToken(refreshTokenString string) (string, bool) {
	if claims, ok := ValidateRefreshToken(refreshTokenString); ok {
		if err := repo.redisClient.Get(refreshTokenString).Err(); err == redis.Nil {
			return claims.UUID, true
		}
	}

	return "", false
}

//DeleteRefreshToken blacklist refresh token
func (repo *TokenRepository) DeleteRefreshToken(tokenString string) error {
	if _, ok := ValidateRefreshToken(tokenString); ok {
		if err := repo.redisClient.Get(tokenString).Err(); err != redis.Nil && err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Error when retriving refreshtoken from server")

			return err
		}
	}

	return nil
}
