package data_types

func data_types() {
	// integers
	var a int = 0   // 4 bytes or 8 bytes (-,+)
	var b int8 = 0  // 8 bit (-128, +127)
	var c int16 = 0 // 16 bit (-32768, +32767)
	var d int32 = 0 // 32 bit = 4 bytes (-...., +....)
	var e int64 = 0 // 64 bit = 8 bytes (-...., +....)

	// uint
	var a uint = 0   // 4 bytes or 8 bytes (0+)
	var b uint8 = 0  // 8 bit (0, 255)
	var c uint16 = 0 // 16 bit (0, 65535)
	var d uint32 = 0 // 32 bit (0+)
	var e uint64 = 0 // 64 bit (0+)

	// other numeric
	var a byte = 0 // synonim uint8 (0, 255)
	var b rune = 0 // synonim int32 (-...,+...)

	// floats
	var a float32 = 0.1  // float 32 bit (single acc)
	var b float64 = 0.11 // float 64 bit (double acc)

	// complex
	var c complex64 = 0 + 0i  // real and imag == float32
	var d complex128 = 0 + 0i // real and imag == float64

	// boolean
	var a bool = true
	var b bool = false

	// string
	var s string = "string"

	// example
	var name, age = "Name", 1
}
