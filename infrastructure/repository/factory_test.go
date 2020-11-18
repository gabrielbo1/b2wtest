package repository

import (
	"github.com/gabrielbo1/b2wtest/config"
	"github.com/gabrielbo1/b2wtest/domain"
	"testing"
)

func init() {
	getDataBase()
}

func TestNewPlanetRepository(t *testing.T) {
	if config.EnvVal(config.Base) == string(MongoDB) {
		var e *config.Err
		var rep domain.PlanetRepository
		if rep = NewPlanetRepository(); rep == nil {
			t.Fatal("Error creating repository")
		}
		planet := domain.Planet{Name: "Mart", Climate: "Arid", Terrain: "terrain"}
		if e = rep.Save(planet); e != nil {

		}

		planets, e := rep.FindAll()
		if e != nil || len(planets) <= 0 {
			t.Fatal(e)
		}

		planet = planets[0]
		planet.Climate = "Arctic"

		e = rep.Update(planet)
		if e != nil {
			t.Fatal(e)
		}

		planet2, err2 := rep.FindById(planet.ID)

		if err2 != nil {
			t.Fatal(e)
		}

		if planet.Climate != planet2.Climate {
			t.Fatal("Update error")
		}

		if e = rep.Delete(planet2.ID); e != nil {
			t.Fatal("Delete error")
		}

		if planets, err := rep.FindByName("Mart"); len(planets) <= 0 || err != nil {
			t.Log(err.OnError())
			t.Fatal("FindByName error ")
		}
	}
}
