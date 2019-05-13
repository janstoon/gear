package mongo

import (
	"gopkg.in/mgo.v2"

	"gitlab.com/janstun/actor"
)

type documentor struct {
	sess *mgo.Session
}

func NewDocumentor(dsn string) (actor.Documentor, error) {
	if s, err := mgo.Dial(dsn); err != nil {
		return nil, err
	} else {
		return documentor{s}, nil
	}
}

func (s documentor) DB(name string) actor.DocDatabase {
	return database{s.sess.DB(name)}
}

func (s documentor) Close() {
	s.sess.Close()
}

// Database

type database struct {
	db *mgo.Database
}

func (d database) Aggregate(name string) actor.DocAggregate {
	return collection{d.db.C(name)}
}

func (d database) FileStorage(prefix string) actor.DocFileStorage {
	return gridfs{d.db.GridFS(prefix)}
}

// DocAggregate

type collection struct {
	coll *mgo.Collection
}

//func (c collection) Create(info *CollectionInfo) error {
//	return c.coll.Create(info)
//}

func (c collection) Drop() error {
	return c.coll.DropCollection()
}

//func (c collection) EnsureIndex(index Index) error {
//	return c.coll.EnsureIndex(index)
//}

func (c collection) EnsureIndexKey(key ...string) error {
	return c.coll.EnsureIndexKey(key...)
}

func (c collection) DropIndex(key ...string) error {
	return c.coll.DropIndex(key...)
}

func (c collection) DropIndexName(name string) error {
	return c.coll.DropIndexName(name)
}

func (c collection) Count() (n int, err error) {
	return c.coll.Count()
}

func (c collection) Find(query interface{}) actor.DocQuery {
	return documentQuery{c.coll.Find(query)}
}

func (c collection) FindId(id interface{}) actor.DocQuery {
	return documentQuery{c.coll.FindId(id)}
}

func (c collection) Insert(docs ...interface{}) error {
	return c.coll.Insert(docs...)
}

//func (c collection) Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error) {
//	return c.coll.Upsert(selector, update)
//}

//func (c collection) UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error) {
//	return c.coll.UpsertId(id, update)
//}

func (c collection) Update(selector interface{}, update interface{}) error {
	return c.coll.Update(selector, update)
}

func (c collection) UpdateId(id interface{}, update interface{}) error {
	return c.coll.UpdateId(id, update)
}

//func (c collection) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error) {
//	return c.coll.UpdateAll(selector, update)
//}

func (c collection) Remove(selector interface{}) error {
	return c.coll.Remove(selector)
}

func (c collection) RemoveId(id interface{}) error {
	return c.coll.RemoveId(id)
}

//func (c collection) RemoveAll(selector interface{}) (info *ChangeInfo, err error) {
//	return c.coll.RemoveAll(selector)
//}

//func (c collection) Bulk() *Bulk {
//	return c.coll.Bulk()
//}

//func (c collection) Pipe(pipeline interface{}) *Pipe {
//	return c.coll.Pipe(pipeline)
//}

type gridfs struct {
	fs *mgo.GridFS
}

func (g gridfs) Create(name string) (actor.DocFile, error) {
	return g.fs.Create(name)
}

func (g gridfs) Open(name string) (actor.DocFile, error) {
	return g.fs.Open(name)
}

func (g gridfs) OpenId(id interface{}) (actor.DocFile, error) {
	return g.fs.OpenId(id)
}

func (g gridfs) Find(query interface{}) actor.DocQuery {
	return documentQuery{g.fs.Find(query)}
}

func (g gridfs) Remove(name string) (err error) {
	return g.fs.Remove(name)
}

func (g gridfs) RemoveId(id interface{}) error {
	return g.fs.RemoveId(id)
}

// Query

type documentQuery struct {
	query *mgo.Query
}

func (q documentQuery) Count() (n int, err error) {
	return q.query.Count()
}

func (q documentQuery) All(result interface{}) error {
	return q.query.All(result)
}

func (q documentQuery) Update(update interface{}, result interface{}) (int, error) {
	changed, err := q.query.Apply(mgo.Change{Update: update}, result)

	return changed.Updated, err
}

func (q documentQuery) Batch(n int) actor.DocQuery {
	q.query.Batch(n)

	return q
}

func (q documentQuery) Distinct(key string, result interface{}) error {
	return q.query.Distinct(key, result)
}

func (q documentQuery) One(result interface{}) error {
	return q.query.One(result)
}

func (q documentQuery) Limit(n int) actor.DocQuery {
	q.query.Limit(n)

	return q
}

func (q documentQuery) Skip(n int) actor.DocQuery {
	q.query.Skip(n)

	return q
}

func (q documentQuery) Prefetch(p float64) actor.DocQuery {
	q.query.Prefetch(p)

	return q
}

func (q documentQuery) Select(selector interface{}) actor.DocQuery {
	q.query.Select(selector)

	return q
}

func (q documentQuery) Sort(fields ...string) actor.DocQuery {
	q.query.Sort(fields...)

	return q
}
