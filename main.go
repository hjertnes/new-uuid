package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hjertnes/utils"
)

func getNumberOfUUIDsToGenerate() int{
	if len(os.Args) == 1 {
		return 1
	}

	i, err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println("Input is not a valid number")
		os.Exit(0)
	}

	return i
}

func main(){
	numberOfUUIDsToGenerate := getNumberOfUUIDsToGenerate()
	for {
		if numberOfUUIDsToGenerate == 0{
			break
		}
		numberOfUUIDsToGenerate--
		fmt.Println(utils.CreateUUID())
	}
}