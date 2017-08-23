package mgo

import (
	"time"
	"tvac/web"

	"gopkg.in/mgo.v2/bson"
)

type Table struct {
	*UnsafeTable
}

func NewTable(db *Database, name string) *Table {
	var t = &Table{UnsafeTable: NewUnsafeTable(db, name)}
	return t
}

func (t *Table) Create(i IModel) error {
	i.BeforeCreate()
	return t.UnsafeInsert(i)
}

func (t *Table) UpdateByID(id string, i IModel) error {
	i.BeforeUpdate()
	return t.UnsafeUpdateByID(id, i)
}

func (t *Table) MarkDelete(id string) error {
	var data = bson.M{
		"dtime": time.Now().Unix(),
	}
	return t.UnsafeUpdateByID(id, data)
}

func (t *Table) ReadAll(ptr interface{}) error {
	return t.UnsafeReadMany(bson.M{"dtime": 0}, ptr)
}

func (t *Table) ReadMany(key string, values []string, ptr interface{}) error {
	return t.UnsafeReadMany(bson.M{"dtime": 0, key: bson.M{"$in": values}}, ptr)
}

func (t *Table) ReadOne(where interface{}, ptr interface{}) error {
	return t.UnsafeReadOne(where, ptr)
}

func (t *Table) ReadByID(id string, ptr interface{}) error {
	return t.UnsafeGetByID(id, ptr)
}

func (t *Table) NotExist(where map[string]interface{}) error {
	where["dtime"] = 0
	var c, err = t.UnsafeTable.UnsafeCount(where)
	if err != nil {
		return err
	}
	if c > 0 {
		return web.BadRequest("already exist")
	}
	return nil
}

func (t *Table) ReadByArrID(ids []string, ptr interface{}) error {
	return t.UnsafeRunGetAll(bson.M{"_id": bson.M{"$in": ids}}, ptr)
}
