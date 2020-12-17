package main

import (
	"fmt"
	"log"
)

// Collection ...
type Collection interface {
	Iterator() Iterator
	Push(interface{})
	Len() int
}

// Iterator ...
type Iterator interface {
	hasNext() bool
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
func (uc *UserCollection) Push(user interface{}) {
	if u, ok := user.(*User); ok {
		uc.users = append(uc.users, u)
	} else {
		log.Println("warning: fail to push into UserCollection", u)
	}
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

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

// Next ...
func (u *UserIterator) Next() interface{} {
	if u.hasNext() {
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
	uc.Push(admin)

	iter := uc.Iterator()

	for iter.hasNext() {
		user := iter.Next()
		fmt.Printf("User is %+v\n", user)
	}
}
