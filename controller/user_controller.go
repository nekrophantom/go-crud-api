package controller

import (
	"crud_api/db"
	"crud_api/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context) error {
	var users []models.User

	result := db.GetDB().Find(&users)
	
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	
	if err := c.Bind(user); err != nil {
		return err
	}

	result := db.GetDB().Create(user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)

	result := db.GetDB().First(user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return result.Error
	}

	if err := c.Bind(user); err != nil {
		return err
	}

	result = db.GetDB().Save(user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)

	result := db.GetDB().First(user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return result.Error
	}

	if err := c.Bind(user); err != nil {
		return err
	}

	result = db.GetDB().Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}