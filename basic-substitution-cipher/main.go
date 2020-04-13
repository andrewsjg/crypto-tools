package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func main() {

	key := makeKey()

	file, err := os.Open("shakespeare.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	for _, char := range b {
		//Convert Char to Uppercase
		upperCaseLetter := unicode.ToUpper(rune(char))

		if int(upperCaseLetter) >= 65 && int(upperCaseLetter) <= 90 {
			fmt.Println(key[string(upperCaseLetter)])
		}
	}
}

// Make a random substitution key
func makeKey() map[string]string {

	a := []int{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	key := make(map[string]string)

	for count, num := range a {
		indexLetter := string(count + 65)
		key[indexLetter] = string(num)
	}

	return key
}
