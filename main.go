package main

import (
	"bufio"        //! Read-Input
	"errors"       //! Handle-Errors
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

	custom "github.com/akash2061/GO-Lang/custom"
)

var pl = fmt.Println               //? Print - Alias
var text = "This is a global text" //? Global var
var count = 0                      //? Global Count

func SayHello() {
	pl("Hello Function...")
}
func add(x int, y int) int {
	return x + y
}
func multi(x float64, y int) float64 {
	return x * float64(y) * float64(y)
}
func getTwo(x int) (int, int) {
	return x + 1, x - 1
}
func fact(f int) int {
	if f == 0 {
		return 1
	}
	return f * fact(f-1)
}
func getSum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
func arrSum(arr []int) int {
	s := 0
	for _, num := range arr {
		s += num
	}
	return s
}
func swap(ptr1 *int, ptr2 *int) {
	temp := *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp
}
func double(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
		pl("Address of:", arr[x], ":", &arr[x])
	}
}
func custom_date() {
	func_date := custom.Date{}
	pl("Function Date:", func_date)
	func_err := func_date.SetDay(12)
	if func_err != nil {
		log.Fatal(func_err)
	}
	func_err = func_date.SetMonth(6)
	if func_err != nil {
		log.Fatal(func_err)
	}
	func_err = func_date.SetYear(1974)
	if func_err != nil {
		log.Fatal(func_err)
	}
	pl("Function Date:", func_date.Day(), "/", func_date.Month(), "/", func_date.Year())
}
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

	//! Typecast
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

	//! Strings
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
	//? Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	//? Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	//? Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	pl("\nFormatting:")
	fmt.Printf("%9f\n", 3.141)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	//! Loops
	pl("\nFor Loops:")
	for x := 65; x <= 70; x++ {
		fmt.Printf("%d - %c\n", x, x)
	}
	pl("\nWhile Loops:")
	fx := 0
	for fx < 3 {
		pl("While: ", fx)
		fx++
	}
	true1 := true
	for true1 {
		// seedSecs := time.Now().Unix()
		// rand.Seed(seedSecs)
		// randNum := rand.Intn(50) + 1
		// pl("\nRandom :", randNum)
		pl("True count:", count)
		count++
		if count == 5 {
			break
		}
	}

	//! Array
	pl("\nRange:")
	aNum := []int{10, 20, 30}
	for _, num := range aNum {
		pl(num)
	}
	pl("\nArray:")
	var arr1 [5]int
	anum := [5]int{1, 2, 3}
	pl("Default Value with Index:")
	for i, num := range arr1 {
		pl(i, " : ", num)
	}
	pl("Array-1:\t", arr1)
	pl("Array-2:\t", anum)

	matrix := [3][3]int{}
	pl("\nMatrix:", matrix)
	matrix_1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	pl("\nMatrix:", matrix_1)
	pl()
	for i, row := range matrix_1 {
		pl("Row", i, row)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("\nI am a string :", bStr)

	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("\nSlice Size :", len(sl1))

	// Cycle with for
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}

	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	pl("Full Array: ", sArr)
	pl("1st 2:")
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)
	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("\nAppend Array with Slice:")
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	sl4 := make([]string, 5)
	pl("\nSl4-Make:", sl4)
	pl()
	SayHello()
	pl("\nAddition Function:", add(3, 5))
	pl("\nArea of Circle Function:", multi(3.141592, 5))

	pl("\nReturn 2 values:")
	pl(getTwo(10))
	a, b := getTwo(10)
	pl("A :", a, "\nB :", b)

	pl("\nFactorial of 5:", fact(5))
	pl("\ngetSum of:", aNum, ":", getSum(aNum...))

	arr := [5]int{1, 3, 5, 7, 9}
	pl("\nArray Sum:", arr, ":", arrSum(arr[:]))

	//! Pointers
	a = 2061
	b = 47
	pl("\nA & B Before Swap:", "\nA :", a, "\nB :", b)
	swap(&a, &b)
	pl("\nA & B After Swap:", "\nA :", a, "\nB :", b)
	pl()
	var aPtr *int = &a
	var bPtr *int = &b
	pl("Address of A:", aPtr)
	pl("Address of B:", bPtr)
	pl("Value of A:  ", *aPtr)
	pl("Value of B:  ", *bPtr)
	pl()

	//! Array using Pointer
	parr := [4]int{1, 2, 3, 4}
	double(&parr)
	pl("Double of Array:", parr)

	//! File-Handling
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		pl("\nFile Created Successfully...")
	}
	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 11}
	for _, num := range iPrimeArr {
		_, err := f.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	pl("\nPrinting File Contant...")
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl(scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
	pl("\nFile-Description:")
	pl(os.Stat("data.txt"))
	pl()
	pl("File Name:", f.Name())
	pl("File Descriptor:", f.Fd()) //? Number of time File opened...

	//! Delete File
	/*
		err = os.Remove("data.txt")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("File deleted successfully.")
		}
	*/
	//! Append to file
	/*
		Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write

		These can be or'ed

		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	//! Check if file exists
	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err = os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		} else {
			pl("\nData Added Successfully...")
		}
	}
	pl()
	pl("Return Form map_func:")
	map_1 := custom.Maps()
	pl(map_1)

	custom.Gene()
	custom.Cyber_Str()
	custom.Compos()
	custom.Udt()
	pl()
	custom_date()
	custom.Cli()
	custom.Func_inter()
}
