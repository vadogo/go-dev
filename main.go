package main

import "fmt"

func main() {
	fmt.Print("Hello Git \n")
	a, b, c, d := 6, 8, 3, 1
<<<<<<< HEAD
	fmt.Printf("6 + 8 + 3 + 1 = %v", a+b+c+d)
=======
	fmt.Printf("6 + 8 + 3 + 1 = %v", a+b+c+d) //comment
>>>>>>> fa879f108658e4694635b1c1d4d34e200442136f
}
