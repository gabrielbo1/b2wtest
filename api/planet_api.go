package api

import (
	"encoding/json"
	"github.com/gabrielbo1/b2wtest/config"
	"github.com/gabrielbo1/b2wtest/domain"
	"github.com/gabrielbo1/b2wtest/infrastructure/repository"
	"github.com/gorilla/mux"
	"github.com/peterhellberg/swapi"
	"net/http"
	"strconv"
)

var rep domain.PlanetRepository = repository.NewPlanetRepository()
var c *swapi.Client = swapi.DefaultClient

func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet domain.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid request payload"))
		return
	}
	if err := rep.Save(planet); err != nil {
		respJSON(w, http.StatusBadRequest, err)
		return
	}
	respJSON(w, http.StatusOK, planet)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet domain.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid request payload"))
		return
	}
	if err := rep.Update(planet); err != nil {
		respJSON(w, http.StatusBadRequest, err)
		return
	}
	respJSON(w, http.StatusOK, planet)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var planets []domain.Planet
	var err *config.Err
	if planets, err = rep.FindAll(); err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid request payload"))
		return
	}
	for i := range planets {
		findPlanetMovies(&planets[i])
	}
	respJSON(w, http.StatusOK, planets)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := rep.FindById(params["id"])
	if err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid ID"))
		return
	}
	findPlanetMovies(&planet)
	respJSON(w, http.StatusOK, planet)
}

func FindByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := rep.FindByName(params["name"])
	for i := range planet {
		findPlanetMovies(&planet[i])
	}
	if err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid Name"))
		return
	}
	respJSON(w, http.StatusOK, planet)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := rep.Delete(params["id"])
	if err != nil {
		respJSON(w, http.StatusBadRequest, config.NewErr().WithMessage("Invalid ID"))
		return
	}
	respJSON(w, http.StatusOK, message{
		Message: "OK",
	})
}

func findPlanetMovies(planet *domain.Planet) {
	s := strconv.FormatInt(planet.Number, 10)
	if n, err := strconv.Atoi(s); err == nil {
		if p, err := c.Planet(n); err == nil {
			planet.FilmURLs = p.FilmURLs
		}
	}
}
