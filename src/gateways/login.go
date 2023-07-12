package gateways

import (
	"bn-login-repositories/domain/entities"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	Login_Liff(user_id string) *entities.LineLoginData
}

func (h HTTPGateway) Login_Liff(c echo.Context) (err error) {
	data := new(entities.LineLoginData)
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	fmt.Println("User LOGIN:", data.Sub)
	checkBlacklist := user.checkBlacklist(data.Sub)
	if checkBlacklist != "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"detail": "Could not validate credentials",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
	})
}
