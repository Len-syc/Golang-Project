package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

type GetPlayerTokenRsp struct {
	user          string
	uid           int
	Token         string
	base          string
	PublicRSAKey  string
	PrivateRSAKey string
	ip            string
	password      string
}

var (
	tempUser     *string
	tempPassword *string
	test         int
)

func GetHTTP() (string, string) {
	var (
		user     string
		password string
	)
	fmt.Print("账号:")
	fmt.Scanln(&user)
	fmt.Print("密码:")
	fmt.Scanln(&password)
	tempUser = &user
	tempPassword = &password
	return *tempUser, *tempPassword
}

func Check() bool {
	switch {
	case tempUser == nil && tempPassword == nil:
		fmt.Print("Error")
		return false
	case tempUser == nil && tempPassword != nil:
		fmt.Printf("[Error]tempUser Error!Value[%v]", *tempUser)
		return false
	case tempUser != nil && tempPassword == nil:
		fmt.Printf("[Error]temppassword Error!Value[%v]", *tempPassword)
		return false
	case (len(*tempUser) != 11 && len(*tempPassword) != 12):
		fmt.Println("账号密码的长度错误(+86长度为11位,密码长度为12为)")
		return false
	case (len(*tempUser) != 11 && len(*tempPassword) == 12):
		fmt.Println("账号的长度错误(+86长度为11位)")
		return false
	case (len(*tempUser) == 11 && len(*tempPassword) != 12):
		fmt.Println("密码的长度错误(密码长度为12)")
		return false
	}
	return true
}

func main() {
	var (
		Len       GetPlayerTokenRsp
		CheckRoot bool
	)
	GetHTTP()
	CheckRoot = Check()
	if CheckRoot == true {
		fmt.Print("验证成功")
	} else if CheckRoot == false {
		fmt.Print("验证失败[code:1]")
		fmt.Print("程序结束,任意键退出")
		fmt.Scanln(&test)
		os.Exit(1)
	} else {
		fmt.Print("未知错误\n")
		fmt.Printf("[Error]System Log(:CheckRoot[%v]\ntempUser[%s]\ntemppassword[%s]", CheckRoot, *tempUser, *tempPassword)
		fmt.Print("验证失败[code:2]")
		fmt.Print("程序结束,任意键退出")
		os.Exit(2)
	}
	Len.user = *tempUser
	Len.password = *tempPassword
	Len.Token = *tempUser + *tempPassword
	Len.base = base64.StdEncoding.EncodeToString([]byte(Len.user))
	Len.ip = "127.0.0.1"
	Len.PublicRSAKey = "<RSAKeyValue><Modulus>xbbx26551G99YiFP2nqHVJ5OMQ" + Len.base + "==</Modulus><Exponent>AQAB</Exponent></RSAKeyValue>"
	Len.PrivateRSAKey = "<RSAKeyValue><Modulus>FtuG7e8nD9geqG48wUEw1kQQ7AFTiTTC1SWeyqDKUM" + Len.base + "==</Modulus><Exponent>AQAB</Exponent></RSAKeyValue>"
	fmt.Printf("Len user: %v\n", Len.user)
	fmt.Printf("Len password: %v\n", Len.password)
	fmt.Printf("Len Token: %v\n", Len.Token)
	fmt.Printf("Len base: %v\n", Len.base)
	fmt.Printf("Len ip: %v\n", Len.ip)
	fmt.Printf("Len PublicRSAKey: %v\n", Len.PublicRSAKey)
	fmt.Printf("Len PrivateRSAKey: %v\n", Len.PrivateRSAKey)
}
