package repository

import (
	"coba_dulu/internal/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	VerifyCredential(email string, password string) (interface{}, error)
	FindByEmail(email string) (entity.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	err := db.connection.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *userConnection) VerifyCredential(email string, password string) (interface{}, error) {
	var user entity.User
	err := db.connection.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, comparePassword(user.Password, []byte(password))
}

func (db *userConnection) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := db.connection.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash a password")
	}
	return string(hash)
}

func comparePassword(hashedPwd string, plainPassword []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return err
	}
	return nil
}
