package goaway

// defaultFalseNegatives is a list of defaultProfanities that are checked for before the defaultFalsePositives are removed
//
// This is reserved for words that may be incorrectly filtered as false positives.
//
// Alternatively, words that are long, or that should mark a string as profane no what the context is
// or whether the word is part of another word can also be included.
//
// Note that there is a test that prevents words from being in both defaultProfanities and defaultFalseNegatives,
var defaultFalseNegatives = []string{
	"asshole",
	"dumbass", // ass -> bASS (FP) -> dumBASS
	"nigger",
}
