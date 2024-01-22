package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func preprocessWords(filename string) []string {
	var words []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	return words
}

func sortString(s string) string {
	characters := strings.Split(s, "")
	sort.Strings(characters)
	return strings.Join(characters, "")
}

func countLetters(s string) map[rune]int {
	letterCount := make(map[rune]int)
	for _, letter := range s {
		letterCount[letter]++
	}
	return letterCount
}

func findWords(input string, words []string) []string {
	var result []string
	inputCount := countLetters(input)
	var sortedWordChars = make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(word)
		sortedWordChars[sortedWord] = append(sortedWordChars[sortedWord], word)
	}

	for sortedWord, word := range sortedWordChars {
		if canFormWord(sortedWord, inputCount) {
			result = append(result, word...)
		}
	}

	return result
}

func canFormWord(word string, inputCount map[rune]int) bool {
	wordCount := countLetters(word)
	for letter, count := range wordCount {
		if inputCount[letter] < count {
			return false
		}
	}
	return true
}

func TestFunction(t *testing.T) {
	words := preprocessWords("words.txt")
	input := "saet"
	res := findWords(input, words)
	expected := []string{"ate", "tea"}
	fmt.Println(res)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("FAIL")
	}
}

func BenchmarkFunction(b *testing.B) {
	b.ResetTimer()
	words := preprocessWords("words.txt")
	input := "saeqwoeiqjbowenopfjpowqijrpowhbetateteadqiwujeqwionfeiwqnropinewpognverojfwroinroiqngpoeirngopeinrgopienvopienv"

	findWords(input, words)
}
