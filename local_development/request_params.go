package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "name.sandesh.sirname.more"
	splits := strings.Split(str, ".")
	field, value := splits[0], splits[1]
	fmt.Println(map[string]string{field: value})
	fmt.Println(splits)

}
