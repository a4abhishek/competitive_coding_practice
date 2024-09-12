package main

import "fmt"

type Anagram struct {
	ID        map[rune]int
	CharCount int
	Members   []string
}

func groupAnagrams(strs []string) [][]string {
	anagrams := []*Anagram{}

	for _, s := range strs {
		curId := map[rune]int{}
		currCharCount := 0
		for _, r := range s {
			curId[r]++
			currCharCount++
		}

		matched := false
		for _, anagram := range anagrams {
			if currCharCount != anagram.CharCount {
				continue
			}

			matched = true
			for char, count := range anagram.ID {
				if curId[char] != count {
					matched = false
					break
				}
			}

			if matched {
				anagram.Members = append(anagram.Members, s)
				break
			}
		}

		if !matched {
			anagrams = append(anagrams, &Anagram{
				ID:        curId,
				CharCount: currCharCount,
				Members:   []string{s},
			})
		}
	}

	answer := [][]string{}
	for _, anagram := range anagrams {
		answer = append(answer, anagram.Members)
	}

	return answer
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))             // [["eat","tea","ate"],["tan","nat"],["bat"]]
	fmt.Println(groupAnagrams([]string{""}))                                                   // [[""]]
	fmt.Println(groupAnagrams([]string{"a"}))                                                  // [["a"]]
	fmt.Println(groupAnagrams([]string{"hhhhu", "tttti", "tttit", "hhhuh", "hhuhh", "tittt"})) // [["hhhhu","hhhuh","hhuhh"],["tttti","tttit","tittt"]]
}
