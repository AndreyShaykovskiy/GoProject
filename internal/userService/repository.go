package userService

import "gorm.io/gorm"

type UserRepository interface {
	// CreateUser - Передаем в функцию user типа User из orm.go
	// возвращаем созданный User и ошибку
	CreateUser(user User) (User, error)
	// GetAllUser - Возвращаем массив из всех пользователей в БД и ошибку
	GetAllUsers() ([]User, error)
	// UpdateUserByID - Передаем id и User, возвращаем обновленный User
	// и ошибку
	UpdateUserByID(id uint, user User) (User, error)
	// DeleteUserByID - Передаем id для удаления, возвращаем только ошибку
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	result := r.db.Model(&User{}).Where("id = ?", id).Select("email", "password").Updates(user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
