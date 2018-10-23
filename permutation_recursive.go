package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
	"strconv"
	"strings"
)


// To Do all permutations: what if a word is like a circle ?
func permutationRecursive(word string, position int) (listPermutations[] string ) {
	listPermutations = append(listPermutations, word)
	arrWord := strings.Split(word, "")
	for i := position ;  i < len(arrWord) ; i++ {
		for j := i + 1; j < len(arrWord); j++ {
			arrWord[i], arrWord[j] = arrWord[j], arrWord[i]
			listPermutations = append(listPermutations,
				permutationRecursive(strings.Join(arrWord, ""), i+1)...
			)
			arrWord[i], arrWord[j] = arrWord[j], arrWord[i]
		}
	}
	return
}

func main()  {
	word := utils.StringInput("")
	listPermutations := permutationRecursive(word, 0)
	fmt.Print("List of possible permutations: ")
	fmt.Println(listPermutations)
	fmt.Println("Check: nb of letter(s) in the word " + strconv.Itoa(len(word)) +
		" give " + strconv.Itoa(len(listPermutations)) + " permutations")
}