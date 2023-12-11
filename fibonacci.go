package exercises

// F(0)=0，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N*）
func Fibonacci(step int) int {
	if step == 0 {
		return 0
	} else if step == 1 {
		return 1
	} else {
		a := 1
		b := 1
		for i := step; i > 2; i-- {
			tmp := a
			a = a + b
			b = tmp
		}
		return a
	}
}
