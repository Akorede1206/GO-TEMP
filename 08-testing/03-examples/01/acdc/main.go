package acdc

// func main() {
// 	a := Sum(1, 2)
// 	fmt.Println(a)

// }

func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
