package repository

import (
	"github.com/gabrielbo1/b2wtest/config"
	"github.com/gabrielbo1/b2wtest/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"time"
)

//DataBase - Defines the type of the constant to
// define supported databases.
type DataBase string

//MongoDB - MongoDB data base.
const MongoDB DataBase = "MONGO"

//mongoDB - MongoDB connector.
var mongoDB *mgo.Database

func getDataBase() {
	switch config.EnvVal(config.Base) {
	case string(MongoDB):
		mongoURI := "mongodb://"
		mongoURI += config.EnvVal(config.BaseUser)
		mongoURI += ":" + config.EnvVal(config.BasePassword)
		mongoURI += "@" + config.EnvVal(config.BaseAddress) + ":" + config.EnvVal(config.BasePort)
		mongoURI += "/" + config.EnvVal(config.BaseName)
		info := &mgo.DialInfo{
			Addrs:    []string{config.EnvVal(config.BaseAddress) + ":" + config.EnvVal(config.BasePort)},
			Timeout:  60 * time.Second,
			Database: config.EnvVal(config.BaseName),
			Username: config.EnvVal(config.BaseUser),
			Password: config.EnvVal(config.BasePassword),
		}

		session, err := mgo.DialWithInfo(info)
		if err != nil {
			panic(err)
		}
		mongoDB = session.DB(config.EnvVal(config.BaseName))
		break
	default:
		log.Fatal("Invalid data base configuration.")
	}

}

//NewPlanetRepository - Return implementation planet repository.
func NewPlanetRepository() domain.PlanetRepository {
	getDataBase()
	switch config.EnvVal(config.Base) {
	case string(MongoDB):
		return newPlanetRepMongo(mongoDB)
	}
	return nil
}
