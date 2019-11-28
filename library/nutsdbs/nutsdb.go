package nustdbs

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
	"runtime"
)

var DBs *DBClient
var Dirs = map[string]string{
	"windows": "tmp/db",
	"darwin":  "/tmp/db",
	"linux":   "/tmp/db",
}

// 创建并打开数据库
func init(){
	DBs=&DBClient{
		bucket:"db0",
		dbdir:Dirs[runtime.GOOS],
	}
	fmt.Println("[modb]",Dirs[runtime.GOOS])
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

func (d *DBClient) CloseDB(){
	_ = d.db.Close()
}


// SetData(keys,value) 写入数据
func (d *DBClient) SetData(keys,value string){
	key := []byte(keys)
	val := []byte(value)
	if err := d.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(d.bucket, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
}

// GetData() 读取数据
func (d *DBClient) GetData(keys string) string {
	key := []byte(keys)
	data:=""
	if err := d.db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(d.bucket, key);err!=nil{
				if err.Error()=="key not found"{
					d.SetData(keys,"")
				}else{
					return err
				}
			}else{
				data=string(e.Value)
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
	return data
}

// GetDatas() 读取key前缀的所有数据
func (d *DBClient) GetDatas(prefix string,limitNum int) []map[string]string{
	var datas []map[string]string
	if err := d.db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.PrefixScan(d.bucket,[]byte(prefix),limitNum)
			if err != nil {
				return err
			}
			data:=make(map[string]string)
			for _, entry := range entries {
				data[string(entry.Key)]=string(entry.Value)
			}
			datas=append(datas,data)
			return nil
		}); err != nil {
		log.Println(err)
	}
	return datas
}

// DelData() 删除数据
func (d *DBClient) DelData(keys string) {
	key := []byte(keys)
	if err := d.db.View(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(d.bucket, key);err!=nil{
				return err
			}
			return nil
		}); err != nil {
		log.Fatal("DelData('",keys,"') error：",err)
	}
}