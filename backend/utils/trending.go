package utils

// hardcoded it for now, when the data is populated, it will be revisted
func GetTrendingDestinations() []Place {
	destinations := []Place{
		{"lagos", 120}, {"abuja", 112}, {"ikeja", 20}, {"lekki", 15}, {"ibadan", 58}, {"benin city", 40}, {"sabon gari", 15},
	}
	return destinations
}

type Place struct {
	City   string
	Hotels int
}

// hardcoded it for now, when the data is populated, it will be revisted
func GetPropertyTypes() []Property {
	properties := []Property{
		{"hotel & suite", "hotel_and_suite.jpeg"}, {"apartment", "apartment.jpeg"},
		{"resort", "resort.jpeg"}, {"guest house", "guest_houe.jpeg"},
	}
	return properties
}

type Property struct {
	Name  string
	Image string
}
