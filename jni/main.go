package main

import (
	"fmt"
	"log"

	"github.com/nqbao/jnigo"
)

func getJavaVersion(vm *jnigo.JVM) string {
	args := []jnigo.JObject{}

	if str, e := vm.NewString("java.version"); e == nil {
		args = append(args, str)
	}

	v, err := vm.CallStaticFunction("java/lang/System", "getProperty", "(Ljava/lang/String;)Ljava/lang/String;", args)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", v.GoValue())
}

func testHashcode(vm *jnigo.JVM) int32 {
	// simulate the example from https://github.com/timob/jnigi#example
	obj, err := vm.NewJClass("java/lang/Object", []jnigo.JObject{})

	if err != nil {
		log.Fatal(err)
	}

	v, err := obj.CallFunction("hashCode", "()I", nil)

	if err != nil {
		log.Fatal(err)
	}

	return v.GoValue().(int32)
}

func main() {
	vm := jnigo.CreateJVM()

	fmt.Printf("Java version: %v\n", getJavaVersion(vm))
	fmt.Printf("Object hash code: %v\n", testHashcode(vm))
}
