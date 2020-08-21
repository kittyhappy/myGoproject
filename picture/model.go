package main

// 操作数据库相关的
import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Info 一个照片信息
type Info struct {
	ID   int64  //Id 主键
	Name string //文件名称
	Path string //保存路径
	Note string //备注
	Unix int64  //时间戳
}

//Db 一个数据库操作全局变量
var Db *sqlx.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ztcbase?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	Db = db
}

// InfoAdd 添加
func InfoAdd(mod *Info) error {
	result, err := Db.Exec("insert into info (`name`,path,note,unix) values (?,?,?,?)", mod.Name, mod.Path, mod.Note, mod.Unix)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	if id < 1 {
		return errors.New("添加失败")
	}
	return nil
}

// InfoGet 查询某条相片
func InfoGet(id int64) (Info, error) {
	mod := Info{}
	err := Db.Get(&mod, "select * from info where id = ?", id)
	return mod, err
}

// InfoList 返回相册列表
func InfoList() ([]Info, error) {
	mods := make([]Info, 0, 8)
	err := Db.Select(&mods, "select * from info")
	return mods, err
}

// InfoDrop 删除
func InfoDrop(id int64) error {
	result, err := Db.Exec("delete from info where id =?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows != 1 {
		return errors.New("删除失败")
	}
	return nil
}
