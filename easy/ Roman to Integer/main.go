package main

// Given a roman numeral, convert it to an integer.

func romanToInt(s string) int {
	numbers := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var sum int
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sum += numbers[string(s[i])]
			break
		}
		if _, ok := numbers[string(s[i])+string(s[i+1])]; ok {
			switch string(s[i]) + string(s[i+1]) {
			case "IV":
				sum += 4
			case "IX":
				sum += 9
			case "XL":
				sum += 40
			case "XC":
				sum += 90
			case "CD":
				sum += 400
			case "CM":
				sum += 900
			}
			if i+1 == len(s)-1 {
				break
			}
			i++
			continue
		}
		sum += numbers[string(s[i])]
	}
	return sum
}
