package mongo

import (
	"errors"
	"labix.org/v2/mgo"
)

const (
	cProfile   = "Profiles"
	cCommunity = "Communities"
)

var (
	c struct {
		Community *mgo.Collection
		Profile   *mgo.Collection
	}
	db *mgo.Database
	s  *mgo.Session
)

var (
	ErrNotConnected = errors.New("No database connection")
)

func Connect(host, dbName string) (err error) {
	if s, err = mgo.Dial(host); err != nil {
		return
	}
	db = s.DB(dbName)
	c.Profile = db.C(cProfile)
	//c.Community = db.C(cCommunity)
	if err = ensureIndexes(); err != nil {
		return
	}
	return
}

func Connected() bool {
	return db != nil
}

func ensureIndexes() (err error) {
	sets := map[string][]mgo.Index{
		cProfile: []mgo.Index{
			mgo.Index{
				Key:        []string{"username"},
				Background: true,
				DropDups:   true,
				Sparse:     false,
				Unique:     true,
			},
			mgo.Index{
				Key:        []string{"id"},
				Background: true,
				DropDups:   true,
				Sparse:     false,
				Unique:     true,
			},
		},
		cCommunity: []mgo.Index{},
	}
	for coll, set := range sets {
		for _, idx := range set {
			db.C(coll).EnsureIndex(idx)
		}
	}
	return
}
