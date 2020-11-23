/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package mapper

import (
	"github.com/superhero-match/superhero-profile/cmd/api/model"
	es "github.com/superhero-match/superhero-profile/internal/es/model"
)

// MapESSuggestionToResult maps ES Superhero to result Superhero.
func MapESSuggestionToResult(s es.Superhero) model.Superhero {
		superhero := model.Superhero{
			ID:                s.ID,
			SuperheroName:     s.SuperheroName,
			MainProfilePicURL: s.MainProfilePicURL,
			ProfilePictures:   make([]model.ProfilePicture, 0),
			Gender:            s.Gender,
			Age:               s.Age,
			Lat:               s.Location.Lat,
			Lon:               s.Location.Lon,
			Birthday:          s.Birthday,
			Country:           s.Country,
			City:              s.City,
			SuperPower:        s.SuperPower,
			AccountType:       s.AccountType,
			CreatedAt:         s.CreatedAt,
		}

		for _, profilePicture := range s.ProfilePictures {
			superhero.ProfilePictures = append(superhero.ProfilePictures, model.ProfilePicture{
				SuperheroID:       profilePicture.SuperheroID,
				ProfilePictureURL: profilePicture.ProfilePictureURL,
				Position:          profilePicture.Position,
			})
		}

	return superhero
}
