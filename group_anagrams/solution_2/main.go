package main

import (
	"fmt"
	"sort"
)

// Helper function to sort a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))             // [["eat","tea","ate"],["tan","nat"],["bat"]]
	fmt.Println(groupAnagrams([]string{""}))                                                   // [[""]]
	fmt.Println(groupAnagrams([]string{"a"}))                                                  // [["a"]]
	fmt.Println(groupAnagrams([]string{"hhhhu", "tttti", "tttit", "hhhuh", "hhuhh", "tittt"})) // [["hhhhu","hhhuh","hhuhh"],["tttti","tttit","tittt"]]
}
