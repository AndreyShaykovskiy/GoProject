package userService

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Определяем структуру User, которая будет представлять таблицу в базе данных
type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt и DeletedAt
	Email      string `json:"email" validate:"required,email"` // Обязательное поле, должно быть корректным email
	Password   string `json:"password" validate:"required,min=8"` // Обязательное поле, минимальная длина 8 символов
}

// Создаем валидатор
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Метод для валидации пользователя
func (u *User) Validate() error {
	return validate.Struct(u)
}
