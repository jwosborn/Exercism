package space

// Planet Name
type Planet string

//Age returns age on Planet
func Age(seconds float64, p Planet) float64 {
	var yearInSeconds float64 = 31557600
	var year float64 = seconds / yearInSeconds
	var op float64
	switch p {
	case "Mercury":
		op = 0.2408467
	case "Venus":
		op = 0.61519726
	case "Earth":
		op = 1
	case "Mars":
		op = 1.8808158
	case "Jupiter":
		op = 11.862615
	case "Saturn":
		op = 29.447498
	case "Uranus":
		op = 84.016846
	case "Neptune":
		op = 164.79132
	}
	return year / op
}
