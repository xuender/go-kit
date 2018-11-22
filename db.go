package kit

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/xuender/go-utils"
)

// DB 数据库对象
type DB struct {
	db *leveldb.DB // 数据对象
}

// Close 关闭数据库
func (d *DB) Close() error {
	return d.db.Close()
}

// Get 数据库数据读取
func (d *DB) Get(key []byte, p interface{}) error {
	data, err := d.db.Get(key, nil)
	if err != nil {
		return err
	}
	return utils.Decode(data, p)
}

// Has 包含
func (d *DB) Has(key []byte) (bool, error) {
	return d.db.Has(key, nil)
}

// Put 保存数据
func (d *DB) Put(key []byte, obj interface{}) error {
	bs, err := utils.Encode(obj)
	if err != nil {
		return err
	}
	return d.db.Put(key, bs, nil)
}

// Delete 删除
func (d *DB) Delete(key []byte) error {
	return d.db.Delete(key, nil)
}

// Iterator 迭代获取数据
func (d *DB) Iterator(prefix []byte, f func(key, value []byte) bool) error {
	iter := d.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		if f(iter.Key(), iter.Value()) {
			break
		}
	}
	iter.Release()
	return iter.Error()
}

// NewDB 新建数据库对象
func NewDB(dir string) (*DB, error) {
	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
