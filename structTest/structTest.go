// structTest project structTest.go
package structTest

import (
	"fmt"
)

type UserN struct {
	Id   int
	Name string
}

func User_Test() {
	user1 := UserN{1, "Tom"}
	user2 := UserN{2, "Jack"}

	fmt.Println(user1.Id, user1.Name)
	fmt.Println(user2.Id, user2.Name)
}

type PointerStruct struct {
	_     int
	Value string
	Point *PointerStruct
}

func PointStruct() {
	a := &PointerStruct{Value: "a"}
	b := &PointerStruct{Value: "b", Point: a}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
}

type AnonsStruct struct {
	Id      int
	Name    string
	contact struct {
		phone    string
		address  string
		postcode string
	}
}

func AnonsStructInit() {
	var d = struct {
		name  string
		age   int
		title string
	}{"zhangql", 111, "Test Anons Var"}

	d2 := AnonsStruct{
		Name: "shaolize",
		Id:   120,
	}
	d2.contact.phone = "13800138000"

	fmt.Println(d, d2)

	c := map[string]struct {
		a int
		b int
	}{
		"a": {1, 2},
		"b": {3, 4},
	}

	fmt.Println(c)

}

func (this *UserN) getUser() {
	println(this.Id, this.Name)
}

func Init_Test() {
	a := UserN{1, "Tom"}
	b := &UserN{1, "Tom"}

	fmt.Println(a.Id, a.Name)
	fmt.Println(b.Id, b.Name)

	a.Id = 3
	b.Id = 4

	fmt.Println(a.Id, a.Name)
	fmt.Println(b.Id, b.Name)

	a.getUser()
}

func (this UserN) Text(x int) {
}

func (this *UserN) Text2(s string) {}

func Method_struct() {
	u := UserN{1, "Tom"}

	u.Text(123)
	u.Text2("abc")

	var f1 func(int) = u.Text

	var f2 func(string) = u.Text2

	f1(456)
	f2("def")

	UserN.Text(u, 123)

	(*UserN).Text2(&u, "hij")

	var f3 func(UserN, int) = UserN.Text

	var f4 func(*UserN, string) = (*UserN).Text2

	f3(u, 789)
	f4(&u, "uio")

}
