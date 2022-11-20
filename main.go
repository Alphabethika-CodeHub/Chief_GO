package main

import (
	"fmt"
	"time"
)

// Interfaces
// A struct is a collection of fields.
type myTypeArr []string
type objBio struct {
	name string
	age  int
}

// 1. var card float64
func main() {
	// ASSIGN VARIABLE With 4 Options
	// 1  card = 3.14
	// 2. card := "Ace of Spades"; it Will Auto Determine What Type of Variable is
	// 3. var card string = "Ace of Spades";
	// 4. var card int = 5

	// Using Mixed Variable Declarations:
	var a, b, c int = 5, 5, 5
	myBio := objBio{"Dafa", 23}

  // Array
	var myArrOfFavoriteFoods [3]string // With Fixed Size
	// myArrOfFavoriteFoods := [3]string{"Pizza", "Nasi Goreng", "Bakmi"} // With Fixed Size
  myArrOfFavoriteFoods[0] = "Pizza"
  myArrOfFavoriteFoods[1] = "Nasi Goreng"
  myArrOfFavoriteFoods[2] = "Bakmi"
  fmt.Println(myArrOfFavoriteFoods)

  // Slices
  // var mySlicesOfFavoriteDrinks 

	// Pointer Explained
	// https://www.geeksforgeeks.org/pointers-in-golang/#:~:text=Pointers%20in%20Go%20programming%20language,memory%20address%20in%20the%20system.
	// Using Pointer to Calculate Trapezoid Area
	// Formula ( base_a + base_b ) / 2 * height
	base_a, base_b, height := 1, 1, 1 // It is a Default Value

	// Alias p_ as Pointer
	p_base_a := &base_a
	// This (*) Known as |dereferencing" or "indirecting
	*p_base_a = 3 // Set base_a Through Pointer

	p_base_b := &base_b
	*p_base_b = 3 // Set base_b Through Pointer

	p_height := &height
	*p_height = 6 // Set height Through Pointer

	// This Will Return Hexadecimal Value
	fmt.Println("Pointer Return:", p_base_a)

	// This Will Return The Value
	// After Being Set Some Value Using Pointer
	fmt.Println("Value After Set by Pointer:", base_a)
	fmt.Println("Value After Set by Pointer:", base_b)
	fmt.Println("Value After Set by Pointer:", height)

	currTime := time.Now()
	fmt.Println("Time Started at:", currTime)
	fmt.Println("Using Mixed Variable Declarations:", a+b+c)
	fmt.Println("Struct Field Using Dot:", myBio.name)

	// %T it is Called FORMATTING DIRECTIVES
	// https://pkg.go.dev/fmt
	// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	fmt.Printf("Using Typeof & Format Directive: %T\n", a)
	fmt.Println("Circle Area of 20: ", CircleArea(25))

	studentGrade := []int{90, 95, 70, 72, 69, 20, 55, 50}
	// CalcStudentGradeWithIfElse(studentGrade)
	CalcStudentGradeWithSwitchCase(studentGrade)

}

func CircleArea(r float64) float64 {
	const pi float64 = 3.14
	return pi * r * r
}

func CalcStudentGradeWithSwitchCase(arrOfGrade []int) {
	for i, grade := range arrOfGrade {
		// First Method
		// switch grade {
		// case 90:
		// 	fmt.Println(i, "Student Grade: ", grade, "= A")
		// case 70:
		// 	fmt.Println(i, "Student Grade: ", grade, "= B")
		// case 50:
		// 	fmt.Println(i, "Student Grade: ", grade, "= C")
		// default:
		// 	fmt.Println(i, "Student Grade: ", grade, "= D")
		// }

		// Second Method
		switch {
		case grade >= 90:
			fmt.Println(i, "Student Grade: ", grade, "= A")
		case grade >= 70:
			fmt.Println(i, "Student Grade: ", grade, "= B")
		case grade >= 50:
			fmt.Println(i, "Student Grade: ", grade, "= C")
		default:
			fmt.Println(i, "Student Grade: ", grade, "= D")
		}
	}
}

func CalcStudentGradeWithIfElse(arrOfGrade []int) {
	for i, grade := range arrOfGrade {
		if grade >= 90 {
			fmt.Println(i, "Student Grade: ", grade, "= A")
		} else if grade >= 70 {
			fmt.Println(i, "Student Grade: ", grade, "= B")
		} else if grade >= 50 {
			fmt.Println(i, "Student Grade: ", grade, "= C")
		} else {
			fmt.Println(i, "Student Grade: ", grade, "= D")
		}
	}
}

func RandomRectangle() {
	// Just Proudly Random Rectangle.
	for i := 1; i <= 5; i++ {
		fmt.Println("")
		for j := 1; j <= 5; j++ {
			fmt.Print("* ")
		}
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("")
		for j := 1; j <= i; j += 1 {
			fmt.Print("* ")
		}
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("")
		for j := 5; j >= i; j -= 1 {
			fmt.Print("* ")
		}
	}
}
