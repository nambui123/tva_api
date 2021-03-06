package mgo

import (
	"cetm/qapi/x/web"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UnsafeTable struct {
	collection *mgo.Collection
	Name       string
}

func NewUnsafeTable(db *Database, Name string) *UnsafeTable {
	return &UnsafeTable{
		collection: db.C(Name),
		Name:       Name,
	}
}

const errRecordNotFound = web.NotFound("record not found")
const errReadDataFailed = web.InternalServerError("read data failed")
const errInsertDataFailed = web.InternalServerError("insert data failed")
const errUpdateDataFailed = web.InternalServerError("update data failed")
const errRemoveDataFailed = web.InternalServerError("remove data failed")
const errCountDataFailed = web.InternalServerError("count data failed")
const errNoOutput = web.InternalServerError("no ouput for data")

func (t *UnsafeTable) C() *mgo.Collection {
	return t.collection
}


func (t *UnsafeTable) UnsafeRunGetAll(query interface{}, ptr interface{}) error {
	if ptr == nil {
		return errNoOutput
	}
	var err = t.collection.Find(query).All(ptr)
	if err != nil {
		return errReadDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeRunGetOne(query interface{}, ptr interface{}) error {
	if ptr == nil {
		return errNoOutput
	}

	var cursor = t.collection.Find(query).Iter()
	var err = cursor.Err()
	if err != nil {
		return errReadDataFailed
	}
	defer cursor.Close()

	if cursor.Next(ptr) {
		return nil
	}
	return errRecordNotFound
}

func (t *UnsafeTable) UnsafeInsert(obj interface{}) error {
	err := t.collection.Insert(obj)
	if err != nil {
		return errInsertDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeCount(where interface{}) (int, error) {
	count, err := t.collection.Find(where).Count()
	if err != nil {
		return 0, errCountDataFailed
	}
	return count, nil
}

func (t *UnsafeTable) UnsafeGetByID(id string, ptr interface{}) error {
	return t.UnsafeRunGetOne(bson.M{"_id": id}, ptr)
}

func (t *UnsafeTable) UnsafeUpdateByID(id string, data interface{}) error {
	err := t.collection.UpdateId(id, bson.M{"$set": data})
	if err != nil {
		//mongoDBLog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UpsertByID(id string, data interface{}) (*mgo.ChangeInfo, error) {
	res, err := t.collection.UpsertId(id, bson.M{"$set": data})
	if err != nil {
		return nil, errUpdateDataFailed
	}
	return res, nil
}

func (t *UnsafeTable) UnsafeDeleteByID(id string) error {
	err := t.collection.RemoveId(id)
	if err != nil {
		return errRemoveDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeUpdateWhere(where, data interface{}) error {
	err := t.collection.Update(where, bson.M{"$set": data})
	if err != nil {
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeReadAll(ptr interface{}) error {
	return t.UnsafeRunGetAll(nil, ptr)
}

func (t *UnsafeTable) UnsafeReadMany(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetAll(where, ptr)
}

func (t *UnsafeTable) UnsafeReadOne(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetOne(where, ptr)
}

func (t *UnsafeTable) EnsureIndex(field string) error {
	return t.collection.EnsureIndex(mgo.Index{
		Key:        []string{field},
		Background: true,
	})
}

func (t *UnsafeTable) IsErrNotFound(err error) bool {
	return err == errRecordNotFound
}
