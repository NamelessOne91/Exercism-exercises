package pangram

// upperToLowerDistance represents the number of ASCII encoded characters
// between a given uppercase letter and its lowercase representation
const upperToLowerDistance = int('a') - int('A')

// alphabetLen represents the number of different letters contained in the english alphabet
const alphabetLen = int('z') - int('a') + 1

// IsPangram returns a boolean representing whether the given string contains every letter
// of the alphabet at least once
func IsPangram(input string) bool {
	if len(input) < alphabetLen {
		return false
	}
	var checker = make(map[int]int, alphabetLen)

	for i := 0; i < len(input); i++ {
		ascii := int(input[i])
		if ascii <= int('Z') && ascii >= int('A') {
			ascii += upperToLowerDistance
		}
		if ascii <= int('z') && ascii >= 'a' {
			checker[ascii]++
		}
	}
	return len(checker) == alphabetLen
}
