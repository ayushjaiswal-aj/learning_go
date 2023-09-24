package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl("sV1: ", sV1)
	pl("sV2: ", sV2)

	pl("Length of string sV1: ", len(sV1))
	pl("sV2 contains Another: ", strings.Contains(sV2, "Another"))
	pl("index of 'o' in sV2: ", strings.Index(sV2, "o"))
	pl("Replace 'o'  with zero in sV2", sV2, strings.Replace(sV2, "o", "0", -2))

	sV3 := "some words    \n"
	pl("sV3: ", sV3)
	sV3 = strings.TrimSpace(sV3)
	pl("sV3 after trim: ", sV3)
	pl("split based on - in string a-b-c-d-e: ", strings.Split("-", "a-b-c-d-e"))
	pl("turn into upper case abcde", strings.ToUpper("abcde"))
	pl("turn into lower case ABCDE", strings.ToLower("ABCDE"))
	pl("tacocat has prefix taco: ", strings.HasPrefix("tacocat", "taco"))
	pl("tacocat has suffix cat: ", strings.HasSuffix("tacocat", "cat"))

	rStr := "abcdefg"
	pl("Rune Count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}
