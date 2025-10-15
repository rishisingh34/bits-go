package conversions

import (
	"fmt"
	"strconv"
)

func Convert() {
	// ----- 1. Numeric conversions -----
	var i int = 42
	var f float64 = float64(i) // int → float64
	var j int = int(f)         // float64 → int (truncates decimal)
	fmt.Println("Numeric conversions:", i, f, j)

	// ----- 2. String ↔ Number -----
	str := "123"
	num, _ := strconv.Atoi(str)    // string → int
	flt, _ := strconv.ParseFloat(str, 64) // string → float64
	s := strconv.Itoa(num)         // int → string
	fmt.Println("String ↔ Number:", num, flt, s)

	// ----- 3. Custom type conversion -----
	type Celsius float64
	type Fahrenheit float64

	var c Celsius = 100
	fahrenheit := Fahrenheit(c) // same underlying type (float64)
	fmt.Println("Custom type conversion:", fahrenheit)

	// ----- 4. String ↔ Byte/Rune slices -----
	text := "GoLang"
	bytes := []byte(text)
	runes := []rune(text)
	fmt.Println("String to []byte:", bytes)
	fmt.Println("String to []rune:", runes)
	fmt.Println("[]byte to string:", string(bytes))

	// ----- 5. Interface (type assertion) -----
	var x interface{} = "hello"
	strVal, ok := x.(string)
	if ok {
		fmt.Println("Interface assertion successful:", strVal)
	} else {
		fmt.Println("Interface assertion failed")
	}

	// ----- 6. Boolean to string -----
	boolVal := true
	boolStr := strconv.FormatBool(boolVal)
	fmt.Println("Bool to string:", boolStr)

	// ----- 7. Float to string -----
	floatStr := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println("Float to string:", floatStr)
}
