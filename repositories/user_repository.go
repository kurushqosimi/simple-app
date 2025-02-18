package repositories

import (
	"context"
	"fmt"
	"simple-app/cache"
	"simple-app/config"
	"simple-app/models"
	"time"
)

func GetUserByID(ctx context.Context, id uint) (interface{}, error) {
	cacheKey := fmt.Sprintf("user:%d", id)

	// Try to get user from cache
	cachedUser, err := cache.GetCache(ctx, cacheKey)
	if err == nil {
		fmt.Println("Cache hit")
		return cachedUser, nil
	}
	fmt.Println("Cache miss")

	// If not in cache, fetch from DB
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Store in cache for 10 minutes
	err = cache.SetCache(ctx, cacheKey, user, 10*time.Minute)
	if err != nil {
		fmt.Println("Failed to cache")
	}
	return &user, nil
}
