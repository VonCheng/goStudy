package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name   string
	Phone  string
	Sex    string
	Adress string
}
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	session, err := mgo.Dial("172.16.21.151:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("boyAndDog").C("user")

	// err = c.Insert(&Person{"GM", "+55 53 8116 9639", "male", "北京市海淀区清河小营桥"},
	// 	&Person{"VonCheng", "+55 53 8402 8510", "femal", "河北省廊坊市大城县"},
	// 	&Person{"泾河", "+55 53 8402 2465", "male", "天津市河东区万新村"},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

  var arr []int;


  s := arr[:]
  fmt.Println(s)
	result := []int{}
	err = c.Find(bson.M{"name": "VonCheng"}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	// var buffer bytes.Buffer
	// buffer.WriteString()
	// this.Ctx.buffer.string()

	// var slice1 = make([]string, 10, 40)
	fmt.Println(result)

	fmt.Println("hello", "world")
	this.Data["json"] = result
	this.ServeJSON()
}
