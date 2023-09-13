package services

import (
	"computer_shop/dao"
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/utils"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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
	condition := fmt.Sprintf("WHERE email = '%s'", user.Email)
	result, err := UserDAO.FindByCondition(condition)
	if err != nil {
		return -1, "Error when getting user", err
	}
	if len(result) > 0 {
		return -1, "Email already registered", errors.New("Email already registered")
	}
	hashed, err := HashPassword(user.Password)
	if err != nil {
		return -1, "Error when hashing password", err
	}
	user.Password = hashed
	user.RoleId = utils.User
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
	} else if userModel[0].Password == "" {
		return user, "This account is blocked", err
	}
	if !CheckPasswordHash(user.Password, userModel[0].Password) {
		return user, "Password not match", err
	}
	return userModel[0], "", nil
}

func (UserService *UserService) SetToken(user models.UserModel, token string) error {
	user.Token = token
	return UserDAO.Update(user)
}

func (UserService *UserService) ResetPassword(password, newPassword string, user models.UserModel) error {
	var userModel []models.UserModel
	condition := fmt.Sprintf("WHERE email = '%s' AND user_id = %d", user.Email, user.UserId)
	userModel, err := UserDAO.FindByCondition(condition)
	if err != nil {
		return errors.New("internal Server")
	}
	if len(userModel) == 0 {
		return errors.New("can not find user")
	}
	if !CheckPasswordHash(password, userModel[0].Password) {
		return errors.New("incorrect password")
	}
	hashed, _ := HashPassword(newPassword)
	user.Password = hashed
	return UserDAO.Update(user)
}

func (UserService *UserService) ChangeAvatar(avatar []*multipart.FileHeader, user models.UserModel) error {
	if len(avatar) == 0 {
		return errors.New("must upload an image")
	}
	urls, errStr, err := helpers.UploadFiles(avatar)
	if err != nil {
		return errors.New(errStr)
	}
	user.Avatar = urls[0]
	return UserDAO.Update(user)
}

func (UserService *UserService) UpdateInformation(phone, address string, user models.UserModel) error {
	user.Phone = phone
	user.Address = address
	return UserDAO.Update(user)
}

func (UserService *UserService) GetAllUsers() ([]models.UserModel, error) {
	userCondition := fmt.Sprintf("WHERE users.role_id != %d", utils.SuperAdmin)
	return UserDAO.FindByCondition(userCondition)
}

func (UserService *UserService) BlockUser(userId int) error {
	user, err := UserDAO.FindById(userId)
	if err != nil {
		return errors.New("can not get user")
	}
	user.Password = ""
	user.Token = ""
	return UserDAO.Update(user)
}
