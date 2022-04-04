package main

import "fmt"

type T struct {
	User User
	// User
}

type User struct {
	Name string
	Age int
}

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

//スライス
type Users []*User

// type Users struct {
// 	Users []*Users
// }

//コンストラクト
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func (u *User) SetName3() {
	u.Name = "A"
}


func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("A")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName2("8")
	user2.SayName()

	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	// fmt.Println(t.Name)

	t.User.SetName3()
	fmt.Println(t)

	user3 := NewUser("user3", 10) 
	fmt.Println(user3)
	fmt.Println(*user3)

	user4 := User{Name: "user4", Age:40}
	user5 := User{Name: "user5", Age:50}
	user6 := User{Name: "user6", Age:60}
	user7 := User{Name: "user7", Age:70}

	users := Users{}

	users = append(users, &user4)
	users = append(users, &user5)
	users = append(users, &user6, &user7)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user4, &user5)

	for _, u := range users2{
		fmt.Println(*u)
	}

	m := map[int]User{
		1: {Name: "user8", Age: 80},
		2: {Name: "user9", Age: 90},
	}

	fmt.Println(m)

	m2 := map[User]string {
		{Name: "user10", Age: 100}: "Tokyo",
		{Name: "user11", Age: 110}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user12"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}

	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// fmt.Println(mi + i)
	mi.Print()
}