package main

/*
#cgo CFLAGS:-I/Library/Java/Home/include -I/Library/Java/Home/include/darwin
#cgo LDFLAGS:-ldl

#include<jni.h>
*/
import "C"

import (
	"fmt"
	"log"

	"github.com/juntaki/jnigo"
)

func main() {
	vm := jnigo.CreateJVM()

	args := []jnigo.JObject{}

	if str, e := vm.NewString("java.version"); e == nil {
		args = append(args, str)
	}

	v, err := vm.CallStaticFunction("java/lang/System", "getProperty", "(Ljava/lang/String;)Ljava/lang/String;", args)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", v.GoValue())
}
