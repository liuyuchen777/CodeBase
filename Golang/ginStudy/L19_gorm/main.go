/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 02:21:55
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 08:53:50
 * @Description: ORM
 * @FilePath: /go/src/goStudy/ginStudy/L19_gorm/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

/*
什么是ORM (Object Relational Mapping)
对象关系映射
对象 -> go语言中的结构体实例
关系 -> 关系数据库
映射 ->
*/

package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

// define model -> gorm.Model
/*
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
*/
// struct tag 反射机制
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Namer    string
	Age      int64
}

func (Animal) TableName() string {
	// 指定表的名字
	return "myAnimal"
}

func main() {
	db, err := gorm.Open("mysql", "root:lyc7758321321@(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	/*
		// 自动迁移
		db.AutoMigrate(&UserInfo{})
		// create a user
		// insert data to mysql
		// insert into UserInfo values(1, "lyc", "male", "basketball")
		u1 := UserInfo{1, "lyc", "male", "basketball"}
		u2 := UserInfo{2, "zzh", "female", "basketball"}
		u3 := UserInfo{3, "xxl", "male", "coding"}
		db.Create(&u1)
		db.Create(&u2)
			db.Create(&u3)
		// search
		var u = new(UserInfo)
		db.First(u)
		fmt.Printf("%#v\n", u)
		// update
		db.Model(&u).Update("hobby", "study")
		// delete
		db.Delete(&u)
	*/
	db.AutoMigrate(&User{})
	// 表名默认是结构体名称的复数
	db.AutoMigrate(&Animal{})
	db.Table("someAnimal").CreateTable(&Animal{})
}
