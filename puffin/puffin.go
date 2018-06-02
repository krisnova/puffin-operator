package puffin

import "fmt"

// PrintPuffin will print an ascii puffing to STDOUT
func PrintPuffin() {

	puffin := `
         .-"-.
        /  ,~a\_
        \  \__))>
        ,) ." \
       /  (    \
      /   )    ;
     /   /     /
   ,/_."''  _.-'
    /_/''"\\___
           '~~~'`

	fmt.Println(puffin)
}
