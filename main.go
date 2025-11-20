/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will display a multiplication table of 9 from 1-12
 */

package main

import "fmt"

func main() {
	// INPUT (Displaying Constants)
	const NUMBERONE int = 1
	const NUMBERTWO int = 2
	const NUMBERTHREE int = 3
	const NUMBERFOUR int = 4
	const NUMBERFIVE int = 5
	const NUMBERSIX int = 6
	const NUMBERSEVEN int = 7
	const NUMBEREIGHT int = 8
	const NUMBERNINE int = 9
	const NUMBERTEN int = 10
	const NUMBERELEVEN int = 11
	const NUMBERTWELVE int = 12

	// PROCESSING (Displaying Calculations)
	answer_1 := NUMBERONE * NUMBERNINE
	answer_2 := NUMBERTWO * NUMBERNINE
	answer_3 := NUMBERTHREE * NUMBERNINE
	answer_4 := NUMBERFOUR * NUMBERNINE
	answer_5 := NUMBERFIVE * NUMBERNINE
	answer_6 := NUMBERSIX * NUMBERNINE
	answer_7 := NUMBERSEVEN * NUMBERNINE
	answer_8 := NUMBEREIGHT * NUMBERNINE
	answer_9 := NUMBERNINE * NUMBERNINE
	answer_10 := NUMBERTEN * NUMBERNINE
	answer_11 := NUMBERELEVEN * NUMBERNINE
	answer_12 := NUMBERTWELVE * NUMBERNINE

	// OUTPUT (Showcasing Calculations)
	fmt.Println("9 x 1 = ", answer_1)
	fmt.Println("9 x 2 = ", answer_2)
	fmt.Println("9 x 3 = ", answer_3)
	fmt.Println("9 x 4 = ", answer_4)
	fmt.Println("9 x 5 = ", answer_5)
	fmt.Println("9 x 6 = ", answer_6)
	fmt.Println("9 x 7 = ", answer_7)
	fmt.Println("9 x 8 = ", answer_8)
	fmt.Println("9 x 9 = ", answer_9)
	fmt.Println("9 x 10 = ", answer_10)
	fmt.Println("9 x 11 = ", answer_11)
	fmt.Println("9 x 12 = ", answer_12)

	fmt.Println("\nDone.")
}