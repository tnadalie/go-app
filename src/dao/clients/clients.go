package clients

import (
	"github.com/tnadalie/go-app/src/dao"

	clients "github.com/tnadalie/go-app/src/models/clients"
	"gopkg.in/mgo.v2/bson"
)

func (m *dao.AppDAO) FindAll() ([]clients.Client, error) {
	var movies []Movie
	err := dao.db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *dao.AppDAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *dao.AppDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *dao.AppDAO) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *dao.AppDAO) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
