/* Alta3 Research | Author: RZFeeser
   Constants - Introduction to constants   */

package main

import "fmt"

// constant created in the package level
// constants can appear anywhere a var can appear
const Pi = 3.14

func main() {

    // constant created at a function level
    const MyURI = "http://example.com:5000/v2/"
    fmt.Println("The protocol, authority, and path /v2/ cannot be modified:", MyURI)
    
    // invoke a variable created at the package level
    fmt.Println("Happy", Pi, "Day")

    // the variable "Xfiles" has been set to a bool type
    const Xfiles = true
    fmt.Println("The truth is out there?", Xfiles)



    var myFloat float64 = 21.54
    var myInt int = 562
    var myInt64 int64 = 120

    var res1 = myFloat + float64(myInt)  // Not Allowed (Compiler Error)
    var res2 = myInt + int(myInt64)  // Not Allowed (Compiler Error)

    fmt.Println(res1, res2)
}

