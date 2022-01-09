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
package service

import (
	"github.com/superhero-match/superhero-profile/cmd/api/model"
	"github.com/superhero-match/superhero-profile/internal/cache"
	"github.com/superhero-match/superhero-profile/internal/config"
	"github.com/superhero-match/superhero-profile/internal/es"
)

// Service interface defines service methods.
type Service interface {
	GetCachedSuggestion(key string) (*model.Superhero, error)
	GetESSuggestion(superheroID string) (*model.Superhero, error)
}

// service holds all the different services that are used when handling request.
type service struct {
	ES    es.ES
	Cache cache.Cache
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (Service, error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &service{
		ES:    e,
		Cache: c,
	}, nil
}
