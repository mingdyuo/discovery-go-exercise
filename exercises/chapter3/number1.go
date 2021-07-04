package chapter3

import "fmt"

var (
	hangeulStart = rune(44032)
	hangeulEnd   = rune(55204)
)

func EndOfWord(word string) string {
	numEnds := 28
	runeWord := []rune(word)
	lastWord := runeWord[len(runeWord)-1]
	if hangeulStart <= lastWord && lastWord <= hangeulEnd {
		index := int(lastWord - hangeulStart)
		if index%numEnds != 0 {
			return "은"
		}
	}
	return "는"
}

func Example_array() {
	fruits := []string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		fmt.Printf("%s%s 맛있다.\n", fruit, EndOfWord(fruit))
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 수박은 맛있다.
}
