package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func user_type() {
	fmt.Println("user_type()")

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println(lisa)

	lisa = user{"Lisa", "lisa@gmail.com", 234, true}
	fmt.Println(lisa)
}

func user_type_field() {
	fmt.Println("user_type_field()")

	type admin struct {
		person user
		level  string
	}

	fred := admin{
		person: user{
			name:       "Fred",
			email:      "fred@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(fred)
}

func assign_diff_type() {
	fmt.Println("assign_diff_type()")

	//type duration int64
	//var i int64
	//var d duration
	//
	//i = 77
	//d = i    // compile error
}

type user1 struct {
	name  string
	email string
}

func (u user1) notify() {
	fmt.Printf("Sending email to user: %s<%s>\n", u.name, u.email)
}

func (u *user1) changeEmail(email string) {
	u.email = email
}

func method() {
	bill := user1{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user1{"Lisa", "lisa@email.com"}
	lisa.notify()
	(*lisa).notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	(&bill).changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}

func main() {
	user_type()
	user_type_field()
	//assign_diff_type()
	method()
}
