package communities

import (
	"labix.org/v2/mgo/bson"
)

type Community struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string
	Url  string
}

var Communities = []Community{
	Community{
		Id:   bson.ObjectIdHex("517097eeca1d531967000001"),
		Name: "DMV Resistance",
		Url:  "https://plus.google.com/communities/103349576921336760265",
	},
	Community{
		Id:   bson.ObjectIdHex("517097eeca1d531967000002"),
		Name: "HR Resistance",
		Url:  "https://plus.google.com/communities/107799543110624446726",
	},
}

func All() []Community {
	return Communities
}
