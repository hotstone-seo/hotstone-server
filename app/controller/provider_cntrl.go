package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/hotstone-seo/hotstone-server/app/urlstore"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// ProviderCntrl is controller for provider function
type ProviderCntrl struct {
	dig.In
	service.ProviderService
	urlstore.URLStoreServer
}

// Route to define API Route
func (c *ProviderCntrl) Route(e *echo.Echo) {
	e.POST("provider/matchRule", c.MatchRule)
	e.POST("provider/retrieveData", c.RetrieveData)
	e.GET("provider/tags", c.Tags)
}

// MatchRule to match rule
func (p *ProviderCntrl) MatchRule(ctx echo.Context) (err error) {
	var (
		req  service.MatchRuleRequest
		resp *service.MatchRuleResponse
	)
	if err = ctx.Bind(&req); err != nil {
		return err
	}
	if resp, err = p.ProviderService.MatchRule(p.URLStoreServer, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (p *ProviderCntrl) RetrieveData(c echo.Context) (err error) {
	var (
		req  service.RetrieveDataRequest
		resp *http.Response
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return err
	}
	if resp, err = p.ProviderService.RetrieveData(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	c.Response().WriteHeader(resp.StatusCode)
	_, err = c.Response().Write(body)
	return
}

func (p *ProviderCntrl) Tags(c echo.Context) (err error) {
	var (
		tags []*service.InterpolatedTag
		ctx  = c.Request().Context()
	)
	if tags, err = p.ProviderService.Tags(ctx, service.ProvideTagsRequest{
		RuleID: 0, // TODO: get from body
	}); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}
