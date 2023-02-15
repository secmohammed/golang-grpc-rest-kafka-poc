package entities

type User struct {
	BaseModel
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"-" `
	Name     string `gorm:"not null" json:"name"`
}

type RegisterUserInput struct {
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	Name                 string `json:"name" binding:"required"`
}
type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserLoginResponse struct {
	BaseModel
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
