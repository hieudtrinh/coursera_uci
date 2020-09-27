package main
import "fmt"
import "math"

func main(){
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))
	fmt.Printf("%.2f\n", math.Trunc(-1.2345))
	fmt.Printf("%.0f\n", math.Trunc(-1.2345))
}

// func main() {
// 	var xtemp int
// 	x1 := 0
// 	x2 := 1
// 	for x:=0; x<5; x++ {
// 	  xtemp = x2
// 	  x2 = x2 + x1
// 	  x1 = xtemp
// 	}
// 	fmt.Println(x2)
//   }
// func main() {
// 	x:=7
// 	switch {
// 	  case x>3:
// 		fmt.Printf("1")
// 	  case x>5:
// 		fmt.Printf("2")
// 	  case x==7:
// 		fmt.Printf("3")
// 	  default: 
// 		fmt.Printf("4")
// 	}
//   }
  