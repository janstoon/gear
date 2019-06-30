package gear

import (
	"database/sql"
	"io"
	"time"
)

// Key/Value Storage
type Dictionary interface {
	Put(key string, value interface{}) error
	Has(key string) bool
	Get(key string) interface{}
	Delete(key string) error
	GetDefault(key string, value interface{}) interface{}
}

type RelationalSqlDatabase interface {
	Begin() (*sql.Tx, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func LimitClause(limit, offset uint) (string, []interface{}) {
	if offset+limit > 0 {
		if 0 == limit {
			return "LIMIT ?, 18446744073709551615", []interface{}{offset}
		} else if offset == 0 {
			return "LIMIT ?", []interface{}{limit}
		} else {
			return "LIMIT ?, ?", []interface{}{offset, limit}
		}
	}

	return "", nil
}

type Documentor interface {
	DB(name string) DocDatabase
	Close()
}

type DocDatabase interface {
	Aggregate(name string) DocAggregate
	FileStorage(prefix string) DocFileStorage
}

type DocAggregate interface {
	// TODO: Create(info *CollectionInfo) error
	Drop() error

	// TODO: EnsureIndex(index Index) error
	EnsureIndexKey(key ...string) error
	DropIndex(key ...string) error
	DropIndexName(name string) error

	Count() (n int, err error)
	Find(query interface{}) DocQuery
	FindId(id interface{}) DocQuery

	Insert(docs ...interface{}) error

	// TODO: Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	// TODO: UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)

	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
	// TODO: UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)

	Remove(selector interface{}) error
	RemoveId(id interface{}) error
	// TODO: RemoveAll(selector interface{}) (info *ChangeInfo, err error)

	// TODO: Bulk() *Bulk
	// TODO: Pipe(pipeline interface{}) *Pipe
}

type DocFileStorage interface {
	Create(name string) (DocFile, error)
	Open(name string) (DocFile, error)
	OpenId(id interface{}) (DocFile, error)
	Find(query interface{}) DocQuery
	Remove(name string) (err error)
	RemoveId(id interface{}) error
}

type DocQuery interface {
	Count() (n int, err error)
	All(result interface{}) error
	// TODO: Apply(change Change, result interface{}) (info *ChangeInfo, err error)
	Update(update interface{}, result interface{}) (int, error)
	Batch(n int) DocQuery
	Distinct(key string, result interface{}) error
	One(result interface{}) error
	Limit(n int) DocQuery
	Skip(n int) DocQuery
	Prefetch(p float64) DocQuery
	Select(selector interface{}) DocQuery
	Sort(fields ...string) DocQuery
}

type DocFile interface {
	io.Reader
	io.Writer
	io.Closer
	io.Seeker

	Id() interface{}
	SetId(id interface{})

	Name() string
	SetName(name string)

	Size() int64

	ContentType() string
	SetContentType(ctype string)

	GetMeta(result interface{}) (err error)
	SetMeta(metadata interface{})

	MD5() string

	SetChunkSize(bytes int)

	UploadDate() time.Time
	SetUploadDate(t time.Time)

	Abort()
}
