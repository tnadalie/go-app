package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type AppDAO struct {
	Server   string
	Database string
}

var DB *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *AppDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}
