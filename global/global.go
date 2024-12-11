package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client
var RedisCtx = context.Background()
