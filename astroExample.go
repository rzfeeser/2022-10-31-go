/* RZFeeser | Alta3 Research
   see http://api.open-notify.org/astros.json */

package main

import "fmt"

type Astro struct {
    name     string
    craft    string
}

// this is our new struct
type SpacePeoples struct {
        people []Astro
        number int
        message string
}

func main() {

    ast1 := Astro{"Cai Xuzhe", "Tiangong"}
    ast2 := Astro{"Chen Dong", "Tiangong"}
    ast3 := Astro{"Liu Yang", "Tiangong"}
    ast4 := Astro{"Sergey Prokopyev", "ISS"}

    // create a slice of people (we need this to create a SpacePeoples
    p := []Astro{ast1, ast2, ast3, ast4}

    // initialize a SpacePeoples struct, "s"
    s := SpacePeoples{p, len(p), "success"}

    fmt.Println(s)

}
