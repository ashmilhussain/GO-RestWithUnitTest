package models

import (
	"time"

	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

/*****************************************************
	    model and crud operations for user
******************************************************/

type User struct {
	UserId    int       `gorm:"primary_key;auto_increment"`
	UserName  string    `gorm:"size:40;not null;unique"`
	Email     string    `gorm:"size:50;not null;unique"`
	RoleId    int       `gorm:"not null"`
	Password  string    `gorm:"size:100;not null;" json:"_"`
	IsDeleted bool      `gorm:"default:false" json:"_"`
	CreatedBy int       `json:"_"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
	UpdatedBy int       `json:"_"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	hashedPassword, _ := Hash(u.Password)
	u.Password = string(hashedPassword)
}

func (u *User) SaveUser(db *gorm.DB) error {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid int) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid int) (*User, error) {

	// To hash the password
	u.Prepare()

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.UserName,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid int) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"deleted": true,
		},
	)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
