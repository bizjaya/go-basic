package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {



	
	openFile()
	//fmt.Print("Hlello ", test, "w orld")
}


func openFile() {
	file, ferr := os.Open("customers.csv")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	// ctr := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Print("Line %d: %s\n", ctr, line)
		//ctr++
		items := strings.Split(line, ",")
		
		fmt.Printf("Name: %s %s Email: %s\n", items[1],items[2],items[3])
		fmt.Println("------")
	}

}