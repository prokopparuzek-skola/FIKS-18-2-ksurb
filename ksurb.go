package main

import "fmt"

func swap(num []uint8, a, b uint64) {
	swap := num[a]
	num[a] = num[b]
	num[b] = swap
}

func qsort(num []uint8) {
	end := uint64(len(num) - 1)
	pivot := num[0]
	less := uint64(0)

	swap(num, 0, end)
	for i := uint64(0); i < end; i++ {
		if num[i] > pivot {
			swap(num, uint64(i), less)
			less++
		}
	}
	swap(num, end, less)
	if less != 0 {
		qsort(num[0:less])
	}
	if less != end {
		qsort(num[less+1:])
	}
}

func double(num []uint8) {
	for i := range num {
		num[i] *= 2
	}
}

func compute(num []uint8, vysledek uint64) bool {
	var sum uint64

	for i := range num {
		sum += uint64(num[i])
	}
	qsort(num)
	double(num)
	//fmt.Println(num)

Compute:
	for i := range num {
		dif := sum - vysledek
		for _, val := range num[i:] {
			if uint64(val) <= dif {
				dif -= uint64(val)
			}
			switch true {
			case dif == 0:
				return true
			case dif < 0:
				continue Compute
			}
		}
	}
	return false
}

func main() {
	var T, X, N uint64

	fmt.Scanf("%d", &T)
	for i := uint64(0); i < T; i++ {
		fmt.Scanf("%d%d", &X, &N)
		num := make([]uint8, N)
		for j := range num {
			fmt.Scanf("%d", &num[j])
		}
		ans := compute(num, X)
		if ans == true {
			fmt.Println("LZE")
		} else {
			fmt.Println("NELZE")
		}
	}
}
