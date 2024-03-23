package search_suggestions_system

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	mapSub := map[string][]string{}
	for _, product := range products {
		for i := 1; i <= len(searchWord); i++ {
			if i > len(product) {
				continue
			}

			s := product[:i]

			if _, ok := mapSub[s]; !ok {
				mapSub[s] = []string{}
			}

			if len(mapSub[s]) < 3 {
				mapSub[s] = append(mapSub[s], product)
			}
		}
	}

	suggested := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		s := searchWord[:i]

		if v, ok := mapSub[s]; ok {
			suggested = append(suggested, v)
		} else {
			suggested = append(suggested, []string{})
		}
	}

	return suggested
}
