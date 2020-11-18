package domain

import (
	"github.com/gabrielbo1/b2wtest/config"
)

//Planet - Planet entity.
type Planet struct {
	ID       string   `bson:"_id" json:"id"`
	Number   int64    `bson:"number" json:"number"`
	FilmURLs []string `json:"films"`
	Name     string   `bson:"name" json:"name"`
	Climate  string   `bson:"climate" json:"climate"`
	Terrain  string   `bson:"terrain" json:"terrain"`
}

//PlanetRepository - Planet repository entity.
type PlanetRepository interface {
	Save(planet Planet) *config.Err

	Update(planet Planet) *config.Err

	FindAll() (planets []Planet, err *config.Err)

	FindById(id string) (planet Planet, err *config.Err)

	FindByName(name string) (planets []Planet, err *config.Err)

	Delete(id string) *config.Err
}
