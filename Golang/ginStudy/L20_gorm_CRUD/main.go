/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 09:27:06
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 09:46:16
 * @Description:
 * @FilePath: /go/src/goStudy/ginStudy/L20_gorm_CRUD/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1. define model
type User struct {
	ID   int64
	Name *string `gorm:"default:'King James'"`
	Age  int64
}

/*
if you want to send "" to db, but when don't have name, give default value
method 1: *string
method 2: 
*/

func main() {
	// link db
	db, err := gorm.Open("mysql", "root:lyc7758321321@(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2. map model and table
	db.AutoMigrate(&User{})

	// 3. create
	/*
		u := User{Name: "lyc", Age: 18} // create a User object
		fmt.Println(db.NewRecord(&u))   // 判断主键是否为空
		db.Create(&u)
		fmt.Println(db.NewRecord(&u))
	*/

	u := User{Name: new(string), Age: 33}
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空, true
	db.Create(&u)
	fmt.Println(db.NewRecord(&u)) // false
}
