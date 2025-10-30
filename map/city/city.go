package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Moscaov", "Klinton", "kirovsk"}, // города с населением 10-99 тыс. человек
		100:  []string{"CPB", "NVZ", "Briansk"},         // города с населением 100-999 тыс. человек
		1000: []string{"KIpr", "Harkov", "Novosibirsk"}, // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Moscaov":     54,
		"Klinton":     65,
		"kirovsk":     789,
		"CPB":         345,
		"NVZ":         548,
		"Briansk":     579,
		"KIpr":        65468,
		"Harkov":      78645,
		"Novosibirsk": 57498,
	}

	fmt.Println(groupCity)

	for k, v := range cityPopulation {
		fmt.Println(k, " - ", v)
	}

	fmt.Println("--------------")

	for keyCity, _ := range cityPopulation {
		var found bool
		for _, valueGroup := range groupCity[100] {
			if keyCity == valueGroup {
				found = true
				break
			}
		}

		if !found {
			delete(cityPopulation, keyCity)
		}
	}

	for k, v := range cityPopulation {
		fmt.Println(k, " - ", v)
	}

}
