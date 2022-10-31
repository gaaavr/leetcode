package main

import "fmt"

func main() {

	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	BubbleSort(ar)
	fmt.Println(ar)

}

func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
}
