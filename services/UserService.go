package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

type UserService struct {
}

var (
	UserDAO dao.UserDAO
)

func HashPassword(password string) (string, error) {
	saltStr := os.Getenv("SALT")
	salt, _ := strconv.Atoi(saltStr)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (UserService *UserService) Register(user models.UserModel) (int, string, error) {
	hashed, err := HashPassword(user.Password)
	if err != nil {
		return -1, "Error when hashing password", err
	}
	user.Password = hashed
	user.RoleId = 1
	id := UserDAO.Create(user)
	return id, "", nil
}

func (UserService *UserService) Login(user models.UserModel) (models.UserModel, string, error) {
	var userModel []models.UserModel
	condition := "WHERE email = '" + user.Email + "'"
	userModel, err := UserDAO.FindByCondition(condition)
	if err != nil {
		return user, "error when getting user", err
	}
	if len(userModel) == 0 {
		return user, "Email not find", err
	}
	if !CheckPasswordHash(user.Password, userModel[0].Password) {
		return user, "Password not match", err
	}
	return userModel[0], "", nil
}
