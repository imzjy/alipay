package alipay

import "fmt"


type Alipay struct {
	
}

func VerifySign(data string) {
	fmt.Println("VerifySign", data)
}

func VerifyNotify(data string) {
	fmt.Println("VerifyNotify", data)
}


func Verify(data string) {
	VerifySign(data)
	VerifyNotify(data)
}