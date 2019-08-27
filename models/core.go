package models

import (
	"database/sql"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/jinzhu/gorm"
	//导入数据库初始化
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) Begin() {
	db.db = db.db.Begin()
}

func (db *DB) Rollback() {
	db.db = db.db.Rollback()
}

func (db *DB) Commit() {
	db.db = db.db.Commit()
}

var (
	db *gorm.DB
)

func NewDB() *DB {
	return &DB{db: db}
}

func init() {
	var err error
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed" + err.Error())
	}
	if err = initDB(); err != nil {
		panic("failed to connect database," + err.Error())
	}
	db.SetLogger(logs.GetLogger("orm"))
	db.LogMode(true)
	// 自动同步表结构 自动迁移数据
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{})
	var count int
	//如果从数据库user表中查询不到数据， 则新增一条admin管理员记录
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@qq.com",
			Pwd:    "123",
			Avatar: "/static/images/info-img.png",
			Role:   0,
		})
	}
}

func initDB() error {
	var err error
	dbconf, err := beego.AppConfig.GetSection("database")
	if err != nil {
		logs.Error(err)
		dbconf = map[string]string{
			"type": "sqlite3",
		}
	}
	switch dbconf["type"] {
	case "mysql":
		db, err = gorm.Open("mysql", dbconf["url"])
	default:
		db, err = gorm.Open("sqlite3", "data/data.db")
	}
	if err != nil {
		return err
	}
	return nil
}

//Model 是文章 的数据表
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createtime"`
	UpdatedAt time.Time  `json:"updatetime"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (db *DB) GetDBTime() *time.Time {
	var t *time.Time
	rows, err := db.db.DB().Query("select datetime(CURRENT_TIMESTAMP,'localtime');")

	// reslist, _ := db.DoHandle(rows)
	// for _, v := range reslist[0] {
	// 	fmt.Println("value: ", v, "reflect.typeof: ", reflect.TypeOf(v))
	// }

	if err != nil {
		logs.Error(err)
		return nil
	}
	rows.Scan(t)
	return t
}

func (db *DB) DoHandle(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}
