package main

import (
	"fmt"
	"unicode"
)

func main() {
	cc_test1 := "oneTwoThree"
	cc_test2 := "saveChangesInTheEditor"

	fmt.Println(camelcase(cc_test1))
	fmt.Println(camelcase(cc_test2))

	cipher_test1 := "abc-xyz"
	cipher_test2 := "There's-a-starman-waiting-in-the-sky"
	cipher_test3 := "middle-Outz"

	fmt.Println(caesarCipher(cipher_test1, 2))
	fmt.Println(caesarCipher(cipher_test2, 3))
	fmt.Println(caesarCipher(cipher_test3, 2))
}

func camelcase(s string) int32 {
	if s == "" {
		return 0
	}

	word_count := 1 // First word will be lowercase

	for _, char := range s {
		if unicode.IsUpper(char) {
			word_count += 1
		}
	} 

	return int32(word_count)
}

func caesarCipher(s string, k int32) string {
    // Write your code here

	var res string

	for _, char := range s {
		if !unicode.IsLetter(char) {
			res += string(char)
		} else {

			new_val := char+k

			if new_val > 'z' {
				new_val = new_val - 26
			}

			res += (string(new_val))
		}


	} 

	return res

}