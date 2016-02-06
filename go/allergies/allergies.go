package allergies

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(num int) (allergies []string) {
	for _, allergen := range allergens {
		if AllergicTo(num, allergen) {
			allergies = append(allergies, []string{allergen}...)
		}
	}
	return allergies
}

func AllergicTo(num int, allergen string) bool {
	for i, allergy := range allergens {
		if allergy == allergen && (num&int(1<<uint(i))) > 0 {
			return true
		}
	}
	return false
}
