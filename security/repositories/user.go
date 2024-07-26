package repositories

import (
	"github.com/ad3n/microservices/models"
)

var email = "surya.iksanudin@ionmobility.com"

type UserRepository struct {
}

func (r *UserRepository) FindByEmail(user *models.User) {
	if user.Email == email {
		user.ID = 1
		user.Password = "123456"

		return
	}

	user = nil
}
