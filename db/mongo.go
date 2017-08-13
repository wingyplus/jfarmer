package db

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	session *mgo.Session
}

func (db *DB) All(collection string, v interface{}) error {
	s := db.session.Copy()
	defer s.Close()
	col := s.DB("QRTrace").C(collection)
	col.Find(bson.M{}).All(v)
	// ...
}
