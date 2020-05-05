package repositories

import (
	"github.com/AndersonOA/makersweb-plataforma-desafios/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return nil, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persis user: %v", err)
		return nil, err
	}

	return user, nil
}
