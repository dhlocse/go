package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
  }

func main() {
  fmt.Println(myFunction(1, 2))
  familyName("loc")
}