package main

import (
	"errors"
	"fmt"
	"log"
)

// 要思考的是，程式有沒有必要把 Iterator 跟 Collection 抽象化出來，
// 以此例來說，抽象化成 interface 的好處，是當有多個實踐 Iterator 的實例要做同一個function進行處理時，可以節省程式碼。
// 參考 interface/polymorphism

// 但其缺點是，Push 變成要填入空界面，才能容納各種 type，導致靜態型別無法判定，須多做 type Assertion
// 如果不做抽象化，UserCollection 的 Push 方法就能在靜態下就約束其型別

// Collection ...
type Collection interface {
	Iterator() Iterator
	Push(interface{}) error
	Len() int
}

// Iterator ...
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// User ...
type User struct {
	name string
	age  int
}

// UserCollection ...
type UserCollection struct {
	users []*User
}

// Len ...
func (uc *UserCollection) Len() int {
	return len(uc.users)
}

// Push ...
func (uc *UserCollection) Push(user interface{}) error {
	if u, ok := user.(*User); ok {
		uc.users = append(uc.users, u)
		return nil
	}
	return errors.New("warning: fail to push into UserCollection")
}

// UserIterator embed UserCollection for iterating
type UserIterator struct {
	UserCollection
	index int
}

// Iterator ...
func (uc *UserCollection) Iterator() Iterator {
	return &UserIterator{
		UserCollection: UserCollection{uc.users},
	}
}

// HasNext ...
func (u *UserIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

// Next ...
func (u *UserIterator) Next() interface{} {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// Admin ...
type Admin struct {
	name  string
	age   int
	admin bool
}

// 要關注的是：client 使用到的東西只有 User, UserCollection
func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}

	user2 := &User{
		name: "b",
		age:  20,
	}

	admin := &Admin{
		name:  "bear",
		age:   42,
		admin: true,
	}

	uc := &UserCollection{}
	uc.Push(user1)
	uc.Push(user2)

	// 嘗試 push 非 User type 的 struct 會造成執行時才能發現錯誤
	err := uc.Push(admin)
	if err != nil {
		log.Println(err)
	}

	iter := uc.Iterator()

	for iter.HasNext() {
		user := iter.Next()
		fmt.Printf("User is %+v\n", user)
	}
}
