package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/juju/errors"

	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type GoogleOauth2UserInfoResp map[string]interface{}

// NewOauth2Config return new instance of oauth2.Config [constructor]
func NewOauth2Config(config config.Config) *oauth2.Config {

	c := oauth2.Config{
		RedirectURL:  config.Oauth2GoogleCallback,
		ClientID:     config.Oauth2GoogleClientID,
		ClientSecret: config.Oauth2GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &c
}

// AuthCntrl is controller to handle authentication
type AuthCntrl struct {
	dig.In
	config.Config
	Oauth2Config *oauth2.Config
}

// Route to define API Route
func (c *AuthCntrl) Route(e *echo.Echo) {
	e.GET("auth/google/login", c.AuthGoogleLogin)
	e.GET("auth/google/callback", c.AuthGoogleCallback)
}

// AuthGoogleLogin handle Google auth login
func (c *AuthCntrl) AuthGoogleLogin(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	// Create oauthState cookie
	oauthState := c.setRandomCookie(ce, "oauthstate", time.Now().Add(20*time.Minute))

	// AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	// validate that it matches the the state query parameter on your redirect callback.
	url := c.Oauth2Config.AuthCodeURL(oauthState)
	// log.Errorf("[auth/google/login] AUTH URL:\n%s\n\n", url)

	return ce.Redirect(http.StatusTemporaryRedirect, url)
}

// AuthGoogleLogin handle Google auth callback
func (c *AuthCntrl) AuthGoogleCallback(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
	// }

	authCallback := func(ce echo.Context) (err error) {
		// Read oauthState from Cookie
		oauthState, err := ce.Cookie("oauthstate")
		if err != nil {
			return errors.Trace(err)
		}

		if ce.QueryParam("state") != oauthState.Value {
			return errors.New("invalid oauth google state")
		}

		userInfoResp, err := c.getUserInfoFromGoogle(ce.QueryParam("code"))
		if err != nil {
			return errors.Trace(err)
		}

		err = c.validateUserInfoResp(userInfoResp)
		if err != nil {
			return errors.Trace(err)
		}

		jwtToken, err := c.generateJwtToken(userInfoResp)
		if err != nil {
			return errors.Trace(err)
		}

		secureTokenCookie := &http.Cookie{
			Name: "secure_token", Value: jwtToken,
			Expires:  time.Now().Add(time.Hour * 72),
			HttpOnly: true, Secure: c.Config.CookieSecure,
		}
		ce.SetCookie(secureTokenCookie)

		tokenCookie := &http.Cookie{
			Name: "token", Value: jwtToken,
			Expires:  time.Now().Add(time.Hour * 72),
			HttpOnly: true, Secure: false,
		}
		ce.SetCookie(tokenCookie)

		return ce.Redirect(http.StatusTemporaryRedirect, c.Oauth2GoogleRedirectSuccess)
	}

	err = authCallback(ce)
	if err != nil {
		log.Error(errors.Details(err))
		return ce.Redirect(http.StatusTemporaryRedirect, c.Oauth2GoogleRedirectFailure)
	}

	return err
}

func (c *AuthCntrl) setRandomCookie(ce echo.Context, cookieName string, expiration time.Time) string {
	b := make([]byte, 16)
	rand.Read(b)

	randomVal := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: cookieName, Value: randomVal, Expires: expiration, HttpOnly: true, Secure: c.Config.CookieSecure}
	ce.SetCookie(cookie)

	return randomVal
}

func (c *AuthCntrl) getUserInfoFromGoogle(code string) (userInfoResp GoogleOauth2UserInfoResp, err error) {
	// Use code to get token and get user info from Google.
	token, err := c.Oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if !token.Valid() {
		return nil, errors.New("invalid token")
	}

	response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userInfoResp)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return userInfoResp, nil
}

func (c *AuthCntrl) validateUserInfoResp(userInfoResp GoogleOauth2UserInfoResp) error {
	if verifiedEmail, ok := userInfoResp["verified_email"]; !ok || !verifiedEmail.(bool) {
		return errors.New("invalid or empty verified_email")
	}

	if c.Config.Oauth2GoogleHostedDomain != "" {
		if hd, ok := userInfoResp["hd"]; !ok || hd != c.Config.Oauth2GoogleHostedDomain {
			return errors.New("invalid or empty hd")
		}
	}
	return nil
}

func (c *AuthCntrl) generateJwtToken(userInfoResp GoogleOauth2UserInfoResp) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userInfoResp["email"]
	claims["picture"] = userInfoResp["picture"]
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(c.Config.JwtSecret))
	if err != nil {
		return "", errors.Trace(err)
	}

	return t, nil
}