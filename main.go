package main

import (
	"fmt"
	"github.com/kgoval/erresp"
)

func init()  {
	// registering error message
	erresp.Register(ErrorMessage)
}

func main(){

	fmt.Println(erresp.New(ErrInvalidPhoneFormat, LangID, ""))
}
