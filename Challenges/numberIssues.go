

// Challenge 2.1

package Challenges

import (
	"fmt"
)
// Challenge 2
/* You and your colleagues want to create a code
that displays all numbers between 1 and 100
that are divisible by 3. */
func NumberIssue1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Challenge 2
/* When speaking the numbers, whenever a multiple of 3 appears,
the participant must say "Pin" and in multiples of 5, the participant
must say "Pan". So, your program must print numbers from 1 to 100 and
in the cases mentioned, the numbers must not appear, but rather,
what the participant must say. */
func NumberIssue2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
