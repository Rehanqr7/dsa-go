package math

/*
GCD and LCM (Easy)
*/

func GCD(a, b int) int {
	if a < 0 { a = -a }
	if b < 0 { b = -b }
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 { return 0 }
	g := GCD(a, b)
	return a / g * b
}
