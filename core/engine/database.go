package engine

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey;auto_increment" json:"id" form:"id"`
	CreatedAt time.Time      `json:"create_at" form:"create_at"`
	UpdatedAt time.Time      `json:"update_at" form:"update_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at" form:"delete_at"`
}

type ModelID struct {
	ID int `gorm:"primary_key;auto_increment" json:"id"`
}

/***** LocalTime *****/
type LocalTime struct {
	time.Time
}

// LocalTime 自定义输出JSON格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// LocalTime 写入数据库时会调用
func (t LocalTime) Value() (driver.Value, error) {
	log.Println("localtime:", t)
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// LocalTime 读取数据库时会调用该方法
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Paginate 数据库GORM数据分页
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * pageSize
		//fmt.Println("page:", page, " limit", pageSize)
		return db.Offset(offset).Limit(pageSize)
	}
}

// BuildQuery 含分页的Sql生成器
func BuildQuery(
	c *gin.Context,
	db *gorm.DB,
	wheres interface{},
	columns interface{},
	orderBy interface{},
) (*gorm.DB, int64, error) {
	var err error
	//查询条件
	db, err = BuildWhere(db, wheres)
	if err != nil {
		return nil, 0, err
	}
	//选择字段
	db = db.Select(columns)

	//排序方式
	if orderBy != nil && orderBy != "" {
		db = db.Order(orderBy)
	}
	//数据总数
	var count int64
	db.Count(&count)

	//数据分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * pageSize
	//fmt.Println("page:", page, " limit", pageSize)
	//return db.Offset(offset).Limit(pageSize)
	if page > 0 && pageSize > 0 {
		db.Offset(offset).Limit(pageSize)
	}
	return db, count, err
}

// BuildWhere 构建where条件
/**
1、结构体条件测试
where := user.User{ID: 1, UserName: "chen"}
db, err = BuildWhere(db, where)
db.Find(&users)
2、and条件测试
where := []interface{}{
	[]interface{}{"id", "=", 1},
	[]interface{}{"username", "chen"},
}
// SELECT * FROM `users`  WHERE (id = 1) and (username = 'chen')

3、in,or条件测试
where := []interface{}{
	[]interface{}{"id", "in", []int{1, 2}},
	[]interface{}{"username", "=", "chen", "or"},
}
// SELECT * FROM `users`  WHERE (id in ('1','2')) OR (username = 'chen')

4、not in,or条件测试
where := []interface{}{
	[]interface{}{"id", "not in", []int{1}},
	[]interface{}{"username", "=", "chen", "or"},
}
// SELECT * FROM `users`  WHERE (id not in ('1')) OR (username = 'chen')

5、map条件测试
where := map[string]interface{}{"id": 1, "username": "chen"}
// SELECT * FROM `users`  WHERE (`users`.`id` = '1') AND (`users`.`username` = 'chen')

6、and,or混合条件测试
where := []interface{}{
	[]interface{}{"id", "in", []int{1, 2}},
	[]interface{}{"username = ? or nickname = ?", "chen", "yond"},
}
// SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen' or nickname = 'yond')
*/
func BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	var err error
	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					return nil, errors.New("切片长度不能小于2")
				}
				columnstr := column.(string)
				// 拼接参数形式
				//if strings.Index(columnstr, "?") > -1 {
				if strings.Contains(columnstr, "?") {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/
					optList := " = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp"

					//if strings.Index(" in notin ", _opt) > -1 {
					if strings.Contains(" in notin ", _opt) {
						// val 是数组类型
						column = columnstr + " " + opt + " (?)"
						//} else if strings.Index(optList, _opt) > -1 {
					} else if strings.Contains(optList, _opt) {
						column = columnstr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			} else if t == reflect.Map /*Map*/ {
				db = db.Where(item)
			} else {
				/*
					// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						db, err = BuildWhere(db, item)
						if err != nil {
							panic(err)
						}
						return db
					})*/

				db, err = BuildWhere(db, item)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		return nil, errors.New("参数有误")
	}
	return db, nil
}
