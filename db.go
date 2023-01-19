package main

/*
This is the sorted Word Database

WordLength map is indexed by the length of the word
*/
type Database struct {
	WordLength map[int]SameLengthWord `json:"WordLength"`
}

/*
All words contained in SameLengthWord have the same length

classifiedWords map is indexed by a value computed from the characters contained in a word

the key is computed as follows:

a --> 0b00000000000000000000000000000001
b --> 0b00000000000000000000000000000010
...
ü --> 0b00010000000000000000000000000000
ß --> 0b00100000000000000000000000000000
default: 0b10000000000000000000000000000000

Whenever a word contains one of the characters, OR the value with the sum. At the end, the sum is your key
*/
type SameLengthWord struct {
	ClassifiedWords map[uint32][]string `json:"HashedWords"`
}
