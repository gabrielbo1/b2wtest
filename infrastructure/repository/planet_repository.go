package repository

import (
	"github.com/gabrielbo1/b2wtest/config"
	"github.com/gabrielbo1/b2wtest/domain"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const PlanetEntity = "planet"

//planetRepMongo - Implements planet entity repository with MongoDB.
type planetRepMongo struct {
	mongoDB *mgo.Database
}

func newPlanetRepMongo(mongoDB *mgo.Database) *planetRepMongo {
	return &planetRepMongo{mongoDB: mongoDB}
}

//Save - Save planet with MongoDB.
func (rep *planetRepMongo) Save(planet domain.Planet) *config.Err {
	ai.Connect(rep.mongoDB.C(PlanetEntity))
	planet.ID = bson.NewObjectId().Hex()
	planet.Number = int64(ai.Next(PlanetEntity))
	if err := rep.mongoDB.C(PlanetEntity).Insert(&planet); err != nil {
		return config.NewErr().WithError(err)
	}
	return nil
}

//Update - Update planet with MongoDB.
func (rep *planetRepMongo) Update(planet domain.Planet) *config.Err {
	if err := rep.mongoDB.C(PlanetEntity).UpdateId(planet.ID, &planet); err != nil {
		return config.NewErr().WithError(err)
	}
	return nil
}

//FindAll - FindAll planet with MongoDB.
func (rep *planetRepMongo) FindAll() (planets []domain.Planet, err *config.Err) {
	if err := rep.mongoDB.C(PlanetEntity).Find(bson.M{}).All(&planets); err != nil {
		return nil, config.NewErr().WithError(err)
	}
	return planets, nil
}

//FindById - FindById planet with MongoDB.
func (rep *planetRepMongo) FindById(id string) (planet domain.Planet, err *config.Err) {
	if err := rep.mongoDB.C(PlanetEntity).FindId(id).One(&planet); err != nil {
		return domain.Planet{}, config.NewErr().WithError(err)
	}
	return planet, nil
}

//FindByName - FindByName planet with MongoDB.
func (rep *planetRepMongo) FindByName(name string) (planets []domain.Planet, err *config.Err) {
	if err := rep.mongoDB.C(PlanetEntity).Find(nil).Select(bson.M{"name": name}).All(&planets); err != nil {
		return nil, config.NewErr().WithError(err)
	}

	return planets, nil
}

//Delete - Delete planet with MongoDB.
func (rep *planetRepMongo) Delete(id string) *config.Err {
	if err := rep.mongoDB.C(PlanetEntity).RemoveId(id); err != nil {
		return config.NewErr().WithError(err)
	}
	return nil
}
