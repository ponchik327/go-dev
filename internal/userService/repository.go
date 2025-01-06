package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, newUser User) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}

	user.Email = newUser.Email
	user.PasswordHash = newUser.PasswordHash

	err = r.db.Save(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}

	return r.db.Unscoped().Delete(&user).Error
}
