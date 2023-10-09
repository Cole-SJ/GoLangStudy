package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VipLevel int
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
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VipLevel,
		vip.Price,
	)
}
