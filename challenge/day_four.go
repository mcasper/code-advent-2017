package challenge

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func DayFour() {
	input, err := ioutil.ReadFile("challenge/inputs/day_four.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	inputString := string(input)
	inputArray := strings.Split(inputString, "\n")
	validInputs := make([]string, 0)
	invalidPassphrasesCount := 0

	for _, v := range inputArray {
		if v != "" {
			validInputs = append(validInputs, v)
		}
	}

	for _, passphrase := range validInputs {
		words := strings.Split(passphrase, " ")
		seenWords := map[string]bool{}
		for _, word := range words {
			if _, seen := seenWords[word]; !seen {
				seenWords[word] = true
			} else {
				invalidPassphrasesCount += 1
				break
			}
		}
	}

	fmt.Println(len(validInputs), "total passphrases")
	fmt.Println(len(validInputs)-invalidPassphrasesCount, "valid passphrases")
	fmt.Println(invalidPassphrasesCount, "invalid passphrases")

	// Part 2

	invalidPassphrasesCount2 := 0

	for _, passphrase := range validInputs {
		words := strings.Split(passphrase, " ")
		sortedWords := make([]string, 0)
		for _, word := range words {
			sortedWords = append(sortedWords, SortStringByCharacter(word))
		}

		seenWords := map[string]bool{}
		for _, word := range sortedWords {
			if _, seen := seenWords[word]; !seen {
				seenWords[word] = true
			} else {
				invalidPassphrasesCount2 += 1
				break
			}
		}
	}

	fmt.Println(len(validInputs), "total passphrases")
	fmt.Println(len(validInputs)-invalidPassphrasesCount2, "valid passphrases 2")
	fmt.Println(invalidPassphrasesCount2, "invalid passphrases 2")
}
