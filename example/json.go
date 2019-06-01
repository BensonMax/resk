package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type User struct {
	Id      int    `json:"id,string"`     //[string]标记可以将数字类型以字段串形式来编码和解码
	Name    string `json:"username"`      //使用json tag 可以在序列化时将key 替换成自定义的tag， username 替换了Name
	Age     int    `json:"age,omitempty"` //omitempty 表示忽略空值
	Address string `json:"address"`
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := User{
		Id:      12,
		Name:    "wendell",
		Age:     1,
		Address: "成都高新区",
	}
	// 传入结构体指针， 返回序列化成json的数据结构data err
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	//data:=[]byte(`{"id":"12","username":"wendell","age":1,"Address":"北京"}`)
	u2 := &User{}
	err = json.Unmarshal(data, u2) //将json结构反序列化成 数据结构指针
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", u2)
	fmt.Println(u2.Id)
	fmt.Println(u2.Age)
	fmt.Println(u2.Name)
	fmt.Println(u2.Address)
}
