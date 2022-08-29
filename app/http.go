package app

import (
	"belajar/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"  validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

type M map[string]interface{}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Welcome(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	//data := M{"Message": "Service Running"}
	return c.JSON(http.StatusOK, u)
}

func GroupApi(c echo.Context) error {
	data := M{"Message": "Service Running GroupApi"}
	return c.JSON(http.StatusOK, data)
}

func GetApi(echo *echo.Echo) {
	echo.GET("/api", getDBAll)

}

func GetApiDelete(echo *echo.Echo) {
	echo.DELETE("/api/delete/:email", deleteUserDB)

}

func GetApiUpdate(echo *echo.Echo) {
	echo.PUT("/api/update/:email", updateDBUser)

}

func GetApiAdd(echo *echo.Echo) {
	echo.POST("/api/create", insertDB)

}

func UseSubGroup(group *echo.Group) {
	group.GET("/", getDBAll)
	group.POST("/create", insertDB)
	group.DELETE("/delete/:email", deleteUserDB)
	group.PUT("/update/:email", updateDBUser)

}

func updateDBUser(c echo.Context) (err error) {

	user := new(model.Users)
	c.Bind(user)
	response := new(Response)
	if user.UpdateUser(c.Param("email")) != nil { // method update user
		response.ErrorCode = 10
		response.Message = "Gagal update data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)

}

// ///
func getDBAll(c echo.Context) (err error) {

	response := new(Response)
	users, err := model.GetAll(c.QueryParam("keywords")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data user"
		response.Data = users
	}
	return c.JSON(http.StatusOK, response)

}

func insertDB(c echo.Context) (err error) {

	user := new(model.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
		file, err := c.FormFile("ktp")
		if err != nil {
			fmt.Println("Ktp kosong")
		} else {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			user.Ktp = file.Filename
			fmt.Println("Ada file, akan disimpan")
		}
	}
	response := new(Response)
	if user.CreateUser() != nil { // method create user
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)

}

func deleteUserDB(c echo.Context) (err error) {

	user, _ := model.GetOneByEmail(c.Param("email")) // method get by email
	response := new(Response)

	if user.DeleteUser() != nil { // method update user
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data user"
	}
	c.Logger().Error("report")
	return c.JSON(http.StatusOK, response)

}
