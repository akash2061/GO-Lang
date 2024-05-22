package main

import (
	"bufio"        //! Read-Input
	"fmt"          //! Print-Format
	"log"          //! For Errors
	"math"         //! Math-lib
	"math/rand"    //! Random Numbers
	"os"           //! Input-Output
	"reflect"      //! TypeOf
	"strconv"      //! To Convert Strings
	"strings"      //! To Deal with Strings
	"time"         //! Time
	"unicode/utf8" //! Runes-Count
	"unsafe"       //! SizeOf
)

var pl = fmt.Println               //? Print - Alias
var text = "This is a global text" //? Global var
var count = 0

func main() {
	fmt.Println("Hello... Go...!")

	var name string = "Morningstar_2061"
	var name1 = "akash2061"
	var name2 string
	name3 := "Hello"
	name2 = "Foo..."
	fmt.Println(name, "\n", name1, "\n", name2, "\n", name3, "\n", text)

	var age uint8 = 20
	num := 2.01
	fmt.Println(age, "\t", reflect.TypeOf(age), "\tSize = ", unsafe.Sizeof(age))
	fmt.Println(num, "\t", reflect.TypeOf(num), "\tSize = ", unsafe.Sizeof(num))
	fmt.Println(name2, "\t", reflect.TypeOf(name2), "\tSize = ", unsafe.Sizeof(name2))

	fmt.Printf("My name is %v and age is %v\n", name, age)
	fmt.Printf("My name is %s and age is %d\n", name, age)
	fmt.Printf("name type = %T\nage type = %T\nnum type = %T\n", name, age, num)

	str := fmt.Sprintf("My name is %v and age is %v\n", name, age)
	fmt.Println("\nString is: ", str)

	pl("Hello... Alias... :)")

	pl("\nWhat is your Name?")
	reader := bufio.NewReader(os.Stdin)
	usr_name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", usr_name)
	} else {
		log.Fatal(err)
	}

	pl("Typecast:")
	var v1, v2 = 10, 1.5
	pl(float64(v1) + v2)
	v1 = int(v2)
	pl("Floor of v2:", v1)

	cv := "50"
	cv1, err := strconv.Atoi(cv)
	pl("String: cv", cv, "Atoi: cv1", cv1, "ERROR:", err, "Type1: cv", reflect.TypeOf(cv), "Type2: cv1", reflect.TypeOf(cv1))
	cv2 := "3.14"
	if cv3, err := strconv.ParseFloat(cv2, 32); err == nil {
		pl(cv3)
	}

	pl("\nConditional Statements:")
	iAge := 100
	if (iAge >= 1) && (iAge <= 18) {
		pl("Happy Birthday... Underage... ( ͡° ͜ʖ ͡°) ")
	} else if (iAge == 20) || (iAge == 50) {
		pl("Happy Birthday... 20 or 50? (ᵕ—ᴗ—) ")
	} else if iAge >= 65 {
		pl("Too Old... (╥﹏╥)")
	} else {
		pl("ERROR 404: Birthday Not Found... \\(^o^)/")
	}
	pl("ERROR 404: !true ==", !true)

	pl("\nStrings:")
	str = "A word"
	replacer := strings.NewReplacer("A", "Another")
	str1 := replacer.Replace(str)
	pl("First string:", str)
	pl("Second string:", str1)
	pl("Length of string1:", len(str1))
	pl("Contains Another:", strings.Contains(str1, "Another"))
	pl("o Index:", strings.Index(str1, "o"))
	pl("Replace:", strings.Replace(str1, "o", "0", -1))
	pl("String1 to Lower:", strings.ToLower(str1))
	pl("String1 to Upper:", strings.ToUpper(str1))
	str2 := "H-e-l-l-o--W-o-r-l-d-!"
	pl("String2: ", str2)
	pl("Split string2:", strings.Split(str2, "-"))

	pl("\nString Prefix-Suffix: ( NOTE: Word is a tool name <LolCat> )")
	pl("Prefix:", strings.HasPrefix("LolCat", "Lol"), "\tTrimming Prefix", strings.TrimPrefix("LolCat", "Lol"))
	pl("Suffix:", strings.HasSuffix("LolCat", "Cat"), "\tTrimming Suffix", strings.TrimSuffix("LolCat", "Cat"))

	//! Runes:
	pl("\nRunes:\n")
	rStr := "abcdef"
	pl("Rune Count: ", utf8.RuneCountInString(rStr))
	for i, rune := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, rune, rune)
	}

	//! Time
	pl("\nTime:")
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), ":", now.Minute(), ":", now.Second())

	count++
	pl("Incrementing Count by 1:", count)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("\nRandom :", randNum)
	
	//! There are many math functions
	pl("\nMath :")
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	pl("\nFormatting:")
	fmt.Printf("%9f\n",3.141)
	fmt.Printf("%.2f\n",3.141592)
	fmt.Printf("%9.f\n",3.141592)
}
