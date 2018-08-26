package frame

import "upper.io/db.v2" // required for db.Collection

// the primary key is expected to be named "id" across all tables
type Record interface {
	PrimaryKey() string
	SetPrimaryKey(string)
	Collection() db.Collection
}

func SaveRecord(record Record) error {
	var err error
	var id interface{}
	primaryKey := record.PrimaryKey()
	collection := record.Collection()
	if (primaryKey == "") {
		id, err = collection.Insert(record)
		record.SetPrimaryKey(ToString(id))
	} else {
		err = collection.Find("id", primaryKey).Update(record)
	}
	return err
}

func DeleteRecord(record Record) {
	record.Collection().Find("id", record.PrimaryKey()).Delete()
}
