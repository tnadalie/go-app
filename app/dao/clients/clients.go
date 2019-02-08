package clients

import (
	"log"

	clients "github.com/tnadalie/go-app/src/models/clients"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type AppDAO struct {
	Server   string
	Database string
}

var DB *mgo.Database

const (
	COLLECTION = "clients"
)

func (m *AppDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}

func (m *AppDAO) FindAll() ([]clients.Client, error) {
	var clients []clients.Client
	err := DB.C(COLLECTION).Find(bson.M{}).All(&clients)
	return clients, err
}

func (m *AppDAO) FindById(id string) (clients.Client, error) {
	var client clients.Client
	err := DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&client)
	return client, err
}

func (m *AppDAO) Insert(movie clients.Client) error {
	err := DB.C(COLLECTION).Insert(&movie)
	return err
}

func (m *AppDAO) Delete(movie clients.Client) error {
	err := DB.C(COLLECTION).Remove(&movie)
	return err
}

func (m *AppDAO) Update(movie clients.Client) error {
	err := DB.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
