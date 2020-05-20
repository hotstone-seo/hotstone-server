package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// SettingCntrl is controller to rule entity
type SettingCntrl struct {
	dig.In
	service.SettingService
}

// Route to define API Route
func (c *SettingCntrl) Route(e *echo.Group) {
	e.GET("/settings", c.Find)
	e.PUT("/settings/:key", c.Update)
	e.GET("/settings/:key", c.FindOne)
}

// Find api to get all setting
func (c *SettingCntrl) Find(ec echo.Context) (err error) {
	settings, err := c.SettingService.Find(ec.Request().Context())
	if err != nil {
		return httpError(err)
	}
	return ec.JSON(http.StatusOK, settings)
}

// FindOne api to get one setting
func (c *SettingCntrl) FindOne(ec echo.Context) (err error) {
	setting, err := c.SettingService.FindOne(
		ec.Request().Context(),
		ec.Param("key"),
	)
	if err != nil {
		return httpError(err)
	}
	return ec.JSON(http.StatusOK, setting)
}

// Update api to update setting
func (c *SettingCntrl) Update(ec echo.Context) (err error) {
	var setting repository.Setting

	if err = ec.Bind(&setting); err != nil {
		return err
	}

	if err = c.SettingService.Update(
		ec.Request().Context(),
		ec.Param("key"),
		&setting,
	); err != nil {
		return httpError(err)
	}

	return ec.NoContent(http.StatusOK)
}
