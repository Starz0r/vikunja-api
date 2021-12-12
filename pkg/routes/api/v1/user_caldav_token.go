// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2021 Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package v1

import (
	"net/http"

	"code.vikunja.io/api/pkg/user"
	"code.vikunja.io/web/handler"
	"github.com/labstack/echo/v4"
)

type TokenResponse struct {
	Token string `json:"token"`
}

// GenerateCaldavToken is the handler to create a caldav token
// @Summary Generate a caldav token
// @Description Generates a caldav token which can be used for the caldav api. It is not possible to see the token again after it was generated.
// @tags user
// @Accept json
// @Produce json
// @Security JWTKeyAuth
// @Success 200 {object} v1.TokenResponse
// @Failure 400 {object} web.HTTPError "Something's invalid."
// @Failure 404 {object} web.HTTPError "User does not exist."
// @Failure 500 {object} models.Message "Internal server error."
// @Router /user/settings/token/caldav [post]
func GenerateCaldavToken(c echo.Context) (err error) {

	u, err := user.GetCurrentUser(c)
	if err != nil {
		return handler.HandleHTTPError(err, c)
	}

	token, err := user.GenerateNewCaldavToken(u)
	if err != nil {
		return handler.HandleHTTPError(err, c)
	}

	return c.JSON(http.StatusCreated, &TokenResponse{
		Token: token.ClearTextToken,
	})
}

// GetCaldavTokens is the handler to return a list of all caldav tokens for the current user
// @Summary Returns the caldav tokens for the current user
// @Description Return the IDs and created dates of all caldav tokens for the current user.
// @tags user
// @Accept json
// @Produce json
// @Security JWTKeyAuth
// @Success 200 {array} user.Token
// @Failure 400 {object} web.HTTPError "Something's invalid."
// @Failure 404 {object} web.HTTPError "User does not exist."
// @Failure 500 {object} models.Message "Internal server error."
// @Router /user/settings/token/caldav [get]
func GetCaldavTokens(c echo.Context) error {
	u, err := user.GetCurrentUser(c)
	if err != nil {
		return handler.HandleHTTPError(err, c)
	}

	tokens, err := user.GetCaldavTokens(u)
	if err != nil {
		return handler.HandleHTTPError(err, c)
	}

	return c.JSON(http.StatusCreated, tokens)
}
