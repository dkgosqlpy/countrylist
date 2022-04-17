package countrylist

func Countrycodes() map[int]string {
	country := map[int]string{99: "India", 54: "Argentina", 55: "Brazil", 855: "Cambodia", 62: "Indonesia", 60: "Malaysia", 375: "Belarus", 61: "Australia", 43: "Austria"}

	return country
}

func Countrynames() []string {
	country := []string{"India", "Canada", "Japan", "Russia", "Brazil"}
	return country
}
