package blacklist

import (
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// Repository blacklist repo interface
type Repository interface {
	Blacklist(key string, expiration time.Duration) error
	Blacklisted(key string) (bool, error)
}

// BlacklistRepository blacklist repo state
type BlacklistRepository struct {
	redisClient *redis.Client
}

// NewBlacklistRepository Create new instance of repo
func NewBlacklistRepository(redisClient *redis.Client) *BlacklistRepository {
	return &BlacklistRepository{redisClient: redisClient}
}

// Blacklist insert into blacklist
func (repo *BlacklistRepository) Blacklist(key string, expiration time.Duration) error {
	if err := repo.redisClient.Set(key, key, expiration).Err(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"key": key,
		}).Error("Error while inserting into blacklist")
		return err
	}
	return nil
}

// Blacklisted check if blacklisted
func (repo *BlacklistRepository) Blacklisted(key string) (bool, error) {
	if err := repo.redisClient.Get(key).Err(); err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"key": key,
		}).Error("Error while checking blacklist")
		return true, err
	}
	return true, nil
}
