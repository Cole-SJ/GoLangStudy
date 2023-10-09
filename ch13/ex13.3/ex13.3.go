package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}

	vip := VIPUser{User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("%s %s %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("%s %s %d %d %d",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
