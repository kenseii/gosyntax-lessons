package main

import (
	"fmt"
	"strings"
	"sort"
	"os"
	"log"
	"io/ioutil"
)

// interactions with strings in golang

func main()  {
	sampString:= "Hello world"
	// check if a string contains a certain word
	fmt.Println(strings.Contains(sampString,"Hello"))
	// its index in case it is there
	fmt.Println(strings.Index(sampString,"llo"))
	// counts how often the word is there
	fmt.Println(strings.Count(sampString,"l"))
	// replacing ls with xs for the first 3 ls
	fmt.Println(strings.Replace(sampString,"l", "x", 3))

	// split by, separator
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	// sorting letters
	listOfLetters := []string{"c", "a","b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)

	//join
	listOfNums := strings.Join([]string{"3","2","1"}, ",")
	fmt.Println(listOfNums)


	// file io part
	file, err:= os.Create("sample.txt")
	// in case there is an error display it in the log
	if err != nil {
		log.Fatal(err)
	}
	// actuall writting
	file.WriteString("hahahahaha")

	file.Close()

	// reading the file

	stream, error := ioutil.ReadFile("sample.txt")

	if error != nil {
		log.Fatal(error)
	}

	readString := string(stream)
	fmt.Println(readString)
}
