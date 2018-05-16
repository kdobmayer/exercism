package allergies

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries",
	"tomatoes", "chocolate", "pollen", "cats"}

func Allergies(score uint) []string {
	var allergies []string
	for _, allergen := range allergens {
		if AllergicTo(score, allergen) {
			allergies = append(allergies, allergen)
		}
	}
	return allergies
}

func AllergicTo(score uint, substance string) bool {
	for i, allergen := range allergens {
		if allergen == substance {
			return score&(1<<uint(i)) > 0
		}
	}
	return false
}