/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package service

import (
	"github.com/superhero-match/superhero-profile/cmd/api/model"
	"github.com/superhero-match/superhero-profile/cmd/api/service/mapper"
	"sort"
)

// GetCachedSuggestion fetches suggestion from cache and maps it into result.
func (srv *Service) GetCachedSuggestion(key string) (*model.Superhero, error) {
	cachedSuggestion, err := srv.Cache.GetSuggestion(key)
	if err != nil {
		return nil, err
	}

	if cachedSuggestion == nil {
		return nil, nil
	}

	result := mapper.MapCacheSuggestionToResult(*cachedSuggestion)

	if len(result.ProfilePictures) > 0 {
		sort.Slice(result.ProfilePictures, func(i, j int) bool {
			return result.ProfilePictures[i].Position < result.ProfilePictures[j].Position
		})
	}

	return &result, nil
}
