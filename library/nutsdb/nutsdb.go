package nustdb

import (
	"github.com/xujiajun/nutsdb"
	"log"
)

var DBs *DBClient

// 创建并打开数据库
func init(){
	DBs=&DBClient{
		bucket:"db0",
		dbdir:"/tmp/nutsdb",
	}
	DBs.OpenDB()
}

// IDBClient interface
type IDBClient interface {
	OpenDB()
	GetData(string)
	SetData(string,string)
}

// DBClient struct
type DBClient struct {
	db     *nutsdb.DB
	dbdir  string
	bucket string
}

// OpenDB() 打开数据库
func (d *DBClient) OpenDB(){
	opt := nutsdb.DefaultOptions
	opt.Dir = d.dbdir
	db, err := nutsdb.Open(opt)
	if err != nil {log.Fatal(err)}
	d.db=db
}

// GetData() 读取数据
func (d *DBClient) GetData(keys string)(string,error){
	key := []byte(keys)
	data:=""
	if err := d.db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(d.bucket, key);err!=nil{
				log.Fatal(err)
			}else{
				data=string(e.Value)
			}
			return nil
		}); err != nil {
		return "",err
	}
	return data, nil
}

// SetData(keys,value) 写入数据
func (d *DBClient) SetData(keys,value string) error{
	key := []byte(keys)
	val := []byte(value)
	if err := d.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(d.bucket, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}