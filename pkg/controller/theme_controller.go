package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_resource_handler/pkg/generator"
	"pixri_resource_handler/pkg/model"
)

func GenerateTheme(c echo.Context) error {
	application := model.Application{}
	if error := c.Bind(&application); error != nil {
		return error
	}

	themes := generator.GenerateTheme(application.Name,application.Description,application.ID)

	fmt.Println(themes)
	return c.JSON(http.StatusOK, themes)
}


func ThemeController(g *echo.Group, contextRoot string) {
	g.POST(contextRoot+"/themes/generate", GenerateTheme)
}