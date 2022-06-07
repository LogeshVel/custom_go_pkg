package main

import (
	"cool_project/src/package1"
	"cool_project/src/package2"
	"fmt"
)

func main() {
	fmt.Println("With go mod init")
	fmt.Println(package1.PackageOnevar)
	fmt.Println(package2.VarfromPackage2)
}
