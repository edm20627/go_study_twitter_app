// package mysql
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 検索
	result := []*Users{}
	error := db.Find(&result).Error

	if error != nil || len(result) == 0 {
		return
	}
	for _, user := range result {
		fmt.Println(user.ID, user.Name)
	}
	update(db, 2 , "test")
}

func addUserDate(db *gorm.DB) {
	error := db.Create(&Users{
		Name:     "テスト太郎",
		Age:      20,
		Address:  "東京都杉並区",
		UpdateAt: getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

func findLike(db *gorm.DB, keyword string) {
	result := []*Users{}
	error := db.Where("name LIKE ?", "%"+keyword+"%").Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	for _, user := range result {
		fmt.Println(user.Name)
	}
}

func update(db *gorm.DB, id int, name string) {
	fmt.Println("update")
	error := db.Model(Users{}).Where("id = ?", id).Update(Users{
		Name: name,
		UpdateAt: getDate(),
	}).Error

	if error != nil {
		fmt.Println(error)
	}

	result := []*Users{}
	db.Find(&result)
	for _, user := range result {
		fmt.Println(user.ID, user.Name)
	}
}

func delete(db *gorm.DB, id int) {
	fmt.Println("delete")
	error := db.Where("id = ?", id).Delete(Users{}).Error

	if error != nil {
		fmt.Println(error)
	}

	result := []*Users{}
	db.Find(&result)
	for _, user := range result {
		fmt.Println(user.ID, user.Name)
	}
}

func getDate() string {
	const layout = "2000-01-01 00:00:00"
	now := time.Now()
	return now.Format(layout)
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "go_sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

type Users struct {
	ID       int
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UpdateAt string `json:"updateAt" sql:"not null;type:date`
}
