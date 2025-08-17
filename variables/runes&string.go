package main

import (
	"fmt"
)


func runesAndStrings() {
	// Runes are 32-bit integers representing Unicode code points
	var r rune = 'A'
	fmt.Println("Rune:", r)

	// Strings are a sequence of bytes
	var str string = "Hello, World!"
	fmt.Println("String:", str)

	fmt.Printf("type of each element in string: %T\n", str[0]) // --> byte or uint8

	// here each emoji has more than one byte
	emojis := "😊😊😊😊"

	// indexing a string gives us the first BYTE of the UTF-8 encoding
	fmt.Printf("First index of emojis in character format: %c\n", emojis[0]) // prints � (garbled, since it's just one byte of emoji)
	fmt.Println("First index of emojis in byte format (decimal):", emojis[0]) // prints 240
	fmt.Printf("First index of emojis in byte format (hex): %#x\n", emojis[0]) // prints 0xf0

	// to actually see ALL bytes of the first emoji
	fmt.Printf("All bytes of first emoji: % x\n", emojis[:4])

	// to get the emoji itself we need a bigger container that is a rune
	var emojiRunes = []rune(emojis)
	fmt.Printf("Emoji as rune (decimal code point): %d\n", emojiRunes[0]) // prints 128522
	fmt.Printf("Emoji as rune (hex code point): %#x\n", emojiRunes[0])    // prints 0x1f60a
	fmt.Printf("Emoji as rune (character): %c\n", emojiRunes[0])       
}

