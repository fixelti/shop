package user

import "shop/internal/common/models"

func exist(user models.UserEntity) bool {
	if user.ID == 0 {
		return false
	}
	return true
}
