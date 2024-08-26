package repositories

import (
	"first_api/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	FindAll() ([]entities.APIUser, error)
	FindOneByEmail(email string) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindAll() ([]entities.APIUser, error) {
	var users entities.User
	var usersAPI []entities.APIUser

	// Corrigido: Passa uma referência para o slice de resultados
	if err := u.db.Model(&users).Find(&usersAPI).Error; err != nil {
		return nil, err
	}

	return usersAPI, nil
}

func (u *userRepository) Create(user entities.User) (entities.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (u *userRepository) FindOneByEmail(email string) (entities.User, error) {
	var user entities.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, nil // Retorna um User vazio se não encontrar o registro
		}
		return entities.User{}, err
	}

	return user, nil
}
