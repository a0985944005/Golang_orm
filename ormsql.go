package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// orm資料表若沒設定會預設為自己宣告的struct
type Test struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Chips int    `json:"chips"`
	Phone string `json:"phone"`
	Time  string `json:"time"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterModel(new(Test))

	orm.RegisterDataBase("default", "mysql", "dev:dev123@tcp(192.168.50.253:3306)/CatHome?charset=utf8")
}

func main() {

	// insert()
	// update()
	// delete()
	query()

}

func query() {

	var j, i int64

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默認使用 default，你可以指定為其他database

	var tests []Test
	num, err := o.Raw("SELECT * FROM test ").QueryRows(&tests)
	fmt.Println(tests)

	if err == nil {
		fmt.Println("test nums: ", num)
	}

	i = num
	for j = 0; j < i; j++ {
		fmt.Println(j, tests[j])
	}
	fmt.Printf("%T", tests)
}

func delete() {

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默認使用 default，你可以指定為其他database

	test := Test{Id: 11}

	fmt.Println(o.Delete(&test))
}

func insert() {

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默認使用 default，你可以指定為其他database

	test := Test{
		Name:  "test",
		Chips: 5000,
		Phone: "099999999",
		Time:  time.Now().Format("2006-01-02 15:04:05"),
	}
	fmt.Println(o.Insert(&test))
}

func update() {

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默認使用 default，你可以指定為其他database

	test := Test{
		Id:    2,
		Name:  "Rosco",
		Chips: 5000,
		Phone: "0985944005",
		Time:  time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Println(o.Update(&test))
}
