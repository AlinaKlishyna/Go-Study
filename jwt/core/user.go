package core

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"` // здесь хранится ХЭШ
}

func (d *Domain) CreateUser(user *User) (*User, error) {
	if err := d.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Domain) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := d.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
