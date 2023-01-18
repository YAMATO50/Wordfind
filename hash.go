package main

import "strings"

func hashWord(word string) uint32 {
	var sum uint32 = 0

	for _, letter := range strings.ToLower(word) {
		switch letter {
		case 'a':
			sum = sum | 0b00000000000000000000000000000001
			break
		case 'b':
			sum = sum | 0b00000000000000000000000000000010
			break
		case 'c':
			sum = sum | 0b00000000000000000000000000000100
			break
		case 'd':
			sum = sum | 0b00000000000000000000000000001000
			break
		case 'e':
			sum = sum | 0b00000000000000000000000000010000
			break
		case 'f':
			sum = sum | 0b00000000000000000000000000100000
			break
		case 'g':
			sum = sum | 0b00000000000000000000000001000000
			break
		case 'h':
			sum = sum | 0b00000000000000000000000010000000
			break
		case 'i':
			sum = sum | 0b00000000000000000000000100000000
			break
		case 'j':
			sum = sum | 0b00000000000000000000001000000000
			break
		case 'k':
			sum = sum | 0b00000000000000000000010000000000
			break
		case 'l':
			sum = sum | 0b00000000000000000000100000000000
			break
		case 'm':
			sum = sum | 0b00000000000000000001000000000000
			break
		case 'n':
			sum = sum | 0b00000000000000000010000000000000
			break
		case 'o':
			sum = sum | 0b00000000000000000100000000000000
			break
		case 'p':
			sum = sum | 0b00000000000000001000000000000000
			break
		case 'q':
			sum = sum | 0b00000000000000010000000000000000
			break
		case 'r':
			sum = sum | 0b00000000000000100000000000000000
			break
		case 's':
			sum = sum | 0b00000000000001000000000000000000
			break
		case 't':
			sum = sum | 0b00000000000010000000000000000000
			break
		case 'u':
			sum = sum | 0b00000000000100000000000000000000
			break
		case 'v':
			sum = sum | 0b00000000001000000000000000000000
			break
		case 'w':
			sum = sum | 0b00000000010000000000000000000000
			break
		case 'x':
			sum = sum | 0b00000000100000000000000000000000
			break
		case 'y':
			sum = sum | 0b00000001000000000000000000000000
			break
		case 'z':
			sum = sum | 0b00000010000000000000000000000000
			break
		case 'ä':
			sum = sum | 0b00000100000000000000000000000000
			break
		case 'ö':
			sum = sum | 0b00001000000000000000000000000000
			break
		case 'ü':
			sum = sum | 0b00010000000000000000000000000000
			break
		case 'ß':
			sum = sum | 0b00100000000000000000000000000000
			break
		default:
			sum = sum | 0b10000000000000000000000000000000
		}
	}

	return sum
}
