package vlc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type binaryChunk string
type binaryChunks []binaryChunk
type encodingTable map[rune]string

const chunkSize = 8
func Execute(str string) string {
	// prepare text: M -> !m
	prepStr := prepareText(str)
	// encode to binary: text -> 0101001011
	bStr := encodeBin(prepStr)
	// split binary by chanks (8): bits to bytes -> 0111000 00001111 ...Execute
	chunks := splitByChunks(bStr, chunkSize)
	// bytes to hex: 00001111 11110000 -> 0A A0

	fmt.Println(chunks)
	return ""
}

// prepareText prepares text to be fit for encode:
// changes upper case letters into !+lowecase letter
// i.g. My name is World -> !my name is !world
func prepareText(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch){
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}


//encodeBin function encodes source string into binary without spaces
func encodeBin(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

//splitByChunks function splits source bitSting by chunks with given size
// i.g. '0000111111110000' -> '00001111 11110000'
func splitByChunks(bStr string, chunkSize int) binaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize
	if strLen%chunkSize!=0{
		chunksCount++
	}
	res := make(binaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))
		if (i+1)%chunkSize == 0 {
			res = append(res, binaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0{
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, binaryChunk(lastChunk))
	}
	return res
}

func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}