package main

// import (
//     "fmt"
//     "bufio"
//     "os"
//     "strconv"
// )
//
//
// func main()  {
//     scanner := bufio.NewScanner(os.Stdin)
//     fmt.Printf("What year were you born: \n")
//     scanner.Scan()
//     input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
//     fmt.Printf("you will be %d year old at the end of 2022\n", 2022-input)
// }
// ###############################################################################

import (
	"fmt"
)

//
// func main()  {
//     /*
//     x := 0
//     for x <= 5 {
//         x++
//     }*/
// // these two are the same just different syntax
//     for x := 0; x <= 5; x++ {
//         fmt.Println(x)
//     }
//    }
// #$#############################################################################j

// func main() {
//     var a []int = []int{1,3,4,56,7,12,4,6}
//
//     for i, element := range a {
//         for j := i + 1; j < len(a); j++ {
//             element2 := a[j]
//             if element2 == element {
//                 fmt.Println(element)
//             }
//         }
//     }
// }

// func add(x, y int) (z1 int, z2 int) {
//     defer fmt.Println("hello")
//     z1 = x + y
//     z2 = x - y
//     fmt.Println("before return")
//     return
// }
//
// func main() {
//     ans1,ans2 := add(6, 7)
//     fmt.Println(ans1, ans2)
// }

//	func returnFunc(x string) func() {
//	    return func() { fmt.Println(x)}
//	}
//
//	func main()  {
//	    returnFunc("hello")()
//	}
func main() {
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)
}
