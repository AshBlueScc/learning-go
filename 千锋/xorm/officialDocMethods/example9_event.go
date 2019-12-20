package officialDocMethods

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

func (user User) BeforeInsert() {
	//do something before insert
}

func (user User) BeforeUpdate() {
	//do something before Update
}

func (user User) BeforeDelete() {
	//do something before delete
}

func BeforeSet(name string, cell xorm.Cell) {
	//do something after data is retrieve from database and before data will be set to struct.
}

func AfterSet(name string, cell xorm.Cell) {
	//do something after data is retrieve from database and after data was set to struct.
}

func (user User) AfterInsert() {
	//do something after insert
}

func (user User) AfterUpdate() {
	//do something after update
}

func (user User) AfterDelete() {
	//do something after delete
}

func BasicMethodsTest() {
	_, _ = engine.Before(func(bean interface{}){
		fmt.Println("before", bean)
	}).Insert(new(User))
	//Do before before insert

	_, _ = engine.After(func(bean interface{}) {
		fmt.Println("after", bean)
	}).Insert(new(User))
	//do after after inert
}
