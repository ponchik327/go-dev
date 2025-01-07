package userService

import (
	"project/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUserByID(id uint, user models.User) (models.User, error)
	DeleteUserByID(id uint) error
	GetTasksByUserID(id uint) ([]models.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, newUser models.User) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	user.Email = newUser.Email
	user.PasswordHash = newUser.PasswordHash

	err = r.db.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}

	return r.db.Unscoped().Delete(&user).Error
}

func (r *userRepository) GetTasksByUserID(id uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Joins("JOIN users ON users.id = tasks.user_id").Where("users.id = ?", id).Find(&tasks).Error

	return tasks, err
}
