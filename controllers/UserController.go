package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

var (
	UserService services.UserService
)

func (UserController *UserController) Register(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	if userSession != nil {
		return e.JSON(http.StatusOK, map[string]string{
			"message": "Mày đã đăng nhập",
		})
	}

	var user models.UserModel
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var validator helpers.Validator
	validator.Chain = append(validator.Chain, validator.Required(user.Username, "username không được để trống"))
	validator.Chain = append(validator.Chain, validator.Required(user.Email, "email không được để trống"))
	validator.Chain = append(validator.Chain, validator.IsEmail(user.Email, "email không hợp lệ"))
	validator.Chain = append(validator.Chain, validator.Required(user.Password, "mật khẩu không được để trống"))
	validator.Chain = append(validator.Chain, validator.MinLength(user.Password, 6, "mật khẩu phái có ít nhất 6 kí tự"))
	if err := validator.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_, errStr, err := UserService.Register(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Đăng kí thành công",
	})
}

func (UserController *UserController) Login(e echo.Context) error {
	//userSession, _ := helpers.GetSession("user", e)
	//fmt.Println(userSession)
	//if userSession != nil {
	//	return e.JSON(200, map[string]string{
	//		"message": "Mày đã đăng nhập",
	//	})
	//}

	var user models.UserModel
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var validatorLogin helpers.Validator
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.Required(user.Email, "email không được để trống"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.IsEmail(user.Email, "email không hợp lệ"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.Required(user.Password, "mật khẩu không được để trống"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.MinLength(user.Password, 6, "mật khẩu phái có ít nhất 6 kí tự"))
	if err := validatorLogin.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userResult, errStr, err := UserService.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	if errStr != "" {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	errSession := helpers.SetSession("user", userResult, e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errSession)
	}
	token := helpers.GenarateToken()
	errSetToken := UserService.SetToken(userResult, token)
	if errSetToken != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errSetToken)
	}
	helpers.SetCookie("remember", token, e)
	return e.JSON(http.StatusOK, userResult)
}

func (UserController *UserController) Logout(e echo.Context) error {
	err := helpers.RemoveSession("user", e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Bạn chưa đăng nhập")
	}
	errCookie := helpers.RemoveCookie("remember", e)
	if errCookie != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Bạn chưa đăng nhập")
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Đăng xuất thành công",
	})
}

func (UserController *UserController) ResetPassword(e echo.Context) error {
	oldPassword := e.FormValue("old_password")
	newPassword := e.FormValue("new_password")
	confirmNewPassword := e.FormValue("confirm_new_password")
	var validatorChangePassword helpers.Validator
	validatorChangePassword.Chain = append(validatorChangePassword.Chain, validatorChangePassword.Required(oldPassword, "mật khẩu cũ không được để trống"))
	validatorChangePassword.Chain = append(validatorChangePassword.Chain, validatorChangePassword.Required(newPassword, "mật khẩu mới không được để trống"))
	validatorChangePassword.Chain = append(validatorChangePassword.Chain, validatorChangePassword.MinLength(oldPassword, 6, "mật khẩu cũ phái có ít nhất 6 kí tự"))
	validatorChangePassword.Chain = append(validatorChangePassword.Chain, validatorChangePassword.MinLength(newPassword, 6, "mật khẩu mới phái có ít nhất 6 kí tự"))
	validatorChangePassword.Chain = append(validatorChangePassword.Chain, validatorChangePassword.ComfirmPassword(newPassword, confirmNewPassword, "hai mật khẩu không trùng nhau"))
	err := validatorChangePassword.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	errChangePass := UserService.ResetPassword(oldPassword, newPassword, user)
	if errChangePass != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errChangePass)
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Đổi mật khẩu thành công",
	})
}

func (UserController *UserController) ChangeAvatar(e echo.Context) error {
	form, err := e.MultipartForm()
	if err != nil {
		return e.String(http.StatusBadRequest, "Not a multipart form")
	}
	files := form.File["avatar"]
	if len(files) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Không có hình ảnh được chọn")
	}
	if len(files) > 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "không được truyền quá 1 bức ảnh")
	}
	var valiadatorImage helpers.Validator
	valiadatorImage.Chain = append(valiadatorImage.Chain, valiadatorImage.Required(files[0].Filename, "Hình ảnh không được để trống"))
	valiadatorImage.Chain = append(valiadatorImage.Chain, valiadatorImage.IsImage(files[0]))
	errValidate := valiadatorImage.Validate()
	if errValidate != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errValidate)
	}
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	errChangeImage := UserService.ChangeAvatar(files, user)
	if errChangeImage != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errChangeImage)
	}
	http.Redirect(e.Response(), e.Request(), e.Request().Header.Get("Referer"), 302)
	return e.JSON(http.StatusOK, map[string]string{
		"message": "cập nhật ảnh đại diện thành công",
	})
}

func (UserController *UserController) UpdateInformation(e echo.Context) error {
	phone := e.FormValue("phone")
	address := e.FormValue("address")
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	var validator helpers.Validator
	validator.Chain = append(validator.Chain, validator.Required(phone, "số điện thoại không được để trống"))
	validator.Chain = append(validator.Chain, validator.Required(address, "địa chỉ không được để trống"))
	validator.Chain = append(validator.Chain, validator.IsPhoneNumber(phone, "định dạng số điện thoại sai"))
	if err := validator.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	UserService.UpdateInformation(phone, address, user)
	return e.JSON(http.StatusOK, map[string]string{
		"message": "cập nhật thông tin thành công",
	})
}

func (UserController *UserController) GetAllUsers(e echo.Context) error {
	var roleService services.RoleService
	roles, err := roleService.GetAllRoles()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "Cant get all roles",
		})
	}
	result, err := UserService.GetAllUsers()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "Cant get all users",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"users": result,
		"roles": roles,
	})
}

func (UserController *UserController) BlockUser(e echo.Context) error {
	userIdString := e.Param("id")
	id, err := strconv.Atoi(userIdString)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	errBlock := UserService.BlockUser(id)
	if errBlock != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errBlock.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Block user successfully",
	})
}
