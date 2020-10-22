package main

func exchange(x, y *int64) {
	temp := *x
	*x = *y
	*y = temp
}

func Burbuja(s []int64) {
	length := len(s)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if s[j] > s[j+1] {
				exchange(&s[j], &s[j+1])
			}
		}
	}
}
