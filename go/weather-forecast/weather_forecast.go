// Package weather for getting current weather conditions and locations.
package weather

// CurrentCondition string is a package variable for current weather condition.
var CurrentCondition string

// CurrentLocation string is a package variable for current weather location.
var CurrentLocation string

// Forecast returns the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
