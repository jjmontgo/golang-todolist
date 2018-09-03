package frame

import "upper.io/db.v2" // required for db.Collection

// the primary key is expected to be named "id" across all tables
type Record interface {
	PrimaryKey() uint
	SetPrimaryKey(uint)
	Collection() db.Collection
}

func SaveRecord(record Record) error {
	var err error
	var id interface{}
	primaryKey := record.PrimaryKey()
	collection := record.Collection()
	if (primaryKey > 0) {
		err = collection.Find("id", primaryKey).Update(record)
	} else {
		id, err = collection.Insert(record)
		primaryKey := id.(int64)
		record.SetPrimaryKey(uint(primaryKey))
	}
	return err
}

func DeleteRecord(record Record) {
	record.Collection().Find("id", record.PrimaryKey()).Delete()
}
