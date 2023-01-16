package main

/*
This is the sorted Word Database

WordLength map is indexed by the length of the word
*/
type Database struct {
	WordLength map[int]SameLengthWord
}

/*
All words contained in SameLengthWord have the same length

classifiedWords map is indexed by a value computed from the characters contained in a word

the key is computed as follows:

a-z --> 0000, ä --> 0001, ö --> 0010, ü --> 0100, ß --> 1000

Whenever a word contains one of the characters, OR the value with the sum. At the end, the sum is your key
*/
type SameLengthWord struct {
	classifiedWords map[uint32][]string
}
