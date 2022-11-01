/* Alta3 Research - Slice and append function */

package main

import "fmt"

func main() {

        // append an element to a slice
        a := []int{1, 2}
        fmt.Println("the length of a", len(a))
        fmt.Println("the capacity of a", cap(a))

        a = append(a, 3) // a == [1 2 3]
        fmt.Println("the length of a", len(a))
        fmt.Println("the capacity of a", cap(a))

        // underlying array can still hold this data!
        a = append(a, 42) // a == [1 2 3]
        fmt.Println("the length of a", len(a))
        fmt.Println("the capacity of a", cap(a))

        // out of cap! but thats okay! the slice managers
        // will create a larger array, and dump our old array into the new larger one
        // then update the pointer on the slice
        a = append(a, 88)
        fmt.Println("the length of a", len(a))
        fmt.Println("the capacity of a", cap(a))




        fmt.Println(a)   // [123]





        // Concatenate two slices
        b := []int{1, 2}
        c := []int{11, 22}
        b = append(b, c...) // b = append(b, 11, 22)
        fmt.Println(b) // 1, 2, 11, 22
        
        // The result does not depend on whether the arguments overlap
        // so we can concatenate a slice with itself
        z := []int{1, 2}
        z = append(z, z...) // z == [1 2 1 2]
        fmt.Println(z)
        
}

