package space

import (
	"math"
)

// SecondsInYear is the number of seconds in an Earth year
var SecondsInYear uint64 = 31557600

// Planet is a large body orbiting a star with some period
type Planet string

// Mercury is the 1st planet from the sun
const Mercury = 0.2408467

// Venus is the 2nd planet from the sun
const Venus = 0.61519726

// Earth is the 3rd planet from the sun
const Earth = 1.0

// Mars is the 4th planet from the sun
const Mars = 1.8808158

// Jupiter is the 5th planet from the sun
const Jupiter = 11.862615

// Saturn is the 6th planet from the sun
const Saturn = 29.447498

// Uranus is the 7th planet from the sun
const Uranus = 84.016846

// Neptune is the 8th planet from the sun
const Neptune = 164.79132

func planetPeriod(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return Mercury
	case "Venus":
		return Venus
	case "Earth":
		return Earth
	case "Mars":
		return Mars
	case "Jupiter":
		return Jupiter
	case "Saturn":
		return Saturn
	case "Uranus":
		return Uranus
	case "Neptune":
		return Neptune
	default:
		return math.MaxFloat64
	}
}

// Age is the given seconds represented in orbital periods of the given planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (planetPeriod(planet) * float64(SecondsInYear))
}
