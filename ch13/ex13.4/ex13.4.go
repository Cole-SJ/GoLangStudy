package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 20}

	vip := VIPUser{User{"화랑", "hwarang", 40, 30},
		3,
		250,
	}

	fmt.Printf("%s %s %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("%s %s %d %d %d %d",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.Price,
		vip.User.Level,
	)
}
