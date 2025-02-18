package services

import (
	"context"
	"fmt"
	"simple-app/models"
	"simple-app/repositories"
)

func GetUser(ctx context.Context, id uint) (*models.User, error) {
	userI, err := repositories.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user, ok := userI.(*models.User)
	if !ok {
		return nil, fmt.Errorf("got an unexpexted result: %v", user)
	}

	return user, nil
}
