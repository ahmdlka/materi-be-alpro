package entities

type User struct {
    Common
    Name     string `gorm:"not null" json:"name"`
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `gorm:"not null" json:"-"`
    Role     string `gorm:"default:'user'" json:"role"`
}
