/* Alta3 Research - Creating a slice from an array */

package main

import "fmt"

func main() {

    // 
    myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // we need an array

    // slice from the 1st up to (non inclusive) of the 5th
    mySlice1 := myArray[1:5] // mySlice1 = {1, 2, 3, 4, 5}

    // mySlice1 = {1, 2, 3, 4, 5}  take positions 0 upto (non inclusive) of 2
    mySlice2 := mySlice1[0:2] // mySlice2 = {1, 2}

    fmt.Printf("The type of myArray is %T\n", myArray) // [10]int
    fmt.Println("Displaying myArray:", myArray)        // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9

    fmt.Printf("The type of mySlice1 is %T\n", mySlice1) // []int
    fmt.Println("Displaying mySlice1:", mySlice1)        // 1, 2, 3, 4, 5

    fmt.Printf("The type of mySlice2 is %T\n", mySlice2) // []int
    fmt.Println("Displaying mySlice2:", mySlice2)        // 1, 2

    mySlice2[1] = 22 // updating mySlice2 will also update mySlice1 and myArray

    fmt.Println(mySlice1) // this now reads [1 22 3 4]
    fmt.Println(mySlice2) // this now reads [1 22 3 4]
    fmt.Println(myArray)  // this now reads [0 1 22 3 4 5 6 7 8 9]
    //      updating the slice updates the array

}

