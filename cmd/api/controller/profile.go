/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctrl "github.com/superhero-match/superhero-profile/cmd/api/model"
)

// Profile returns Superhero data with specified id.
func (ctl *Controller) Profile(c *gin.Context) {
	var req ctrl.ProfileRequest

	err := c.BindJSON(&req)
	if checkProfileRequestError(err, c) {
		ctl.Logger.Error(
			"failed to bind JSON to value of type ProfileRequest",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	superhero, err := ctl.Service.GetCachedSuggestion(
		fmt.Sprintf(ctl.SuggestionKeyFormat, req.SuperheroID),
	)
	if checkProfileRequestError(err, c) {
		ctl.Logger.Error(
			"failed to fetch superhero profile from cache",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	fmt.Println("GetCachedSuggestion:")
	fmt.Printf("%+v", superhero)
	fmt.Println("")

	if superhero != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"profile": superhero,
		})

		return
	}

	superhero, err = ctl.Service.GetESSuggestion(req.SuperheroID)
	if checkProfileRequestError(err, c) {
		ctl.Logger.Error(
			"failed to fetch superhero profile from ES",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	fmt.Println("GetESSuggestion:")
	fmt.Printf("%+v", superhero)
	fmt.Println("")

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"profile": superhero,
	})
}

func checkProfileRequestError(err error, c *gin.Context) bool {
	if err != nil {
		var suggestion ctrl.Superhero
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"profile": suggestion,
		})

		return true
	}

	return false
}
