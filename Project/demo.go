package main

import "fmt"

type db_User struct {
	user     string
	password string
	uid      int
}

var (
	Len db_User
)

func main() {
	Len.user = "LenQin"
	Len.password = "root@LenQin123"
	Len.uid = 10001
	dbList(&Len)
}
func dbList(data *db_User) {
	fmt.Printf("User:[%v]\n", data.user)
	fmt.Printf("password:[%v]\n", data.password)
	fmt.Printf("uid:[%v]", data.uid)

}
