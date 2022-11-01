/* RZFeeser | Alta3 Research
   With Pointers - Why we need pointers  */        

package main

import (
   "fmt"
)

type User struct {
   Name string
   Pet  []string
}

// a method for type User
func (p2 User) newPet() {
   p2.Pet = append(p2.Pet, "Fido")
   return len(p2.Pet)
}

func main() {
   u := User{Name: "Anna", Pet: []string{"Bailey"}}         // this time we'll generate a pointer!
   
   fmt.Println(u, "u before")
  
   u.newPet()
   
   fmt.Println(u, "u after")                                // Does Anna have a new pet now? Yes!
}
