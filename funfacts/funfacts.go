package funfacts

import (
	"fmt"
	"strings"

	"github.com/kjellern98/funtemps/conv"
)

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string, temperatureScale string) []string {

	var sunCoreTemp = 15000000.0 // Celsius
	var sunSurfaceTemp = 5778.0  // Kelvin

	var moonSurfaceTempNight = -183.0 // Celsius
	var moonSurfaceTempDay = 106.0    // Celsius

	var earthCoreTemp = 9392.0    // Kelvin
	var lowestTempSurface = -89.4 // Celsius
	var highestTempSurface = 56.7 // Celsius

	var funFactToReturn = FunFacts{}

	switch temperatureScale {
	case "C":
		switch about {
		case "sun":
			sunFactOne := "Temperatur på ytre lag av Solen er " + getFunfactTemperature(sunSurfaceTemp, "K", "C")
			sunFactTwo := "Temperatur i Solens kjerne er " + getFunfactTemperature(sunCoreTemp, "C", "C")
			funFactToReturn.Sun = append(funFactToReturn.Sun, sunFactOne, sunFactTwo)
			return funFactToReturn.Sun
		case "luna":
			lunaFactOne := "Temperatur på månens overflate om natten er " + getFunfactTemperature(moonSurfaceTempNight, "C", "C")
			lunaFactTwo := "Temperatur på månens overflate om dagen er " + getFunfactTemperature(moonSurfaceTempDay, "C", "C")
			funFactToReturn.Luna = append(funFactToReturn.Luna, lunaFactOne, lunaFactTwo)
			return funFactToReturn.Luna
		case "terra":
			terraFactOne := "Laveste målte Temperatur på jordens overflate er " + getFunfactTemperature(lowestTempSurface, "C", "C")
			terraFactTwo := "Høyeste temperatur målt på Jordens overflate " + getFunfactTemperature(highestTempSurface, "C", "C")
			terraFactThree := "Temperatur i jordens kjerne er " + getFunfactTemperature(earthCoreTemp, "K", "C")
			funFactToReturn.Terra = append(funFactToReturn.Terra, terraFactOne, terraFactTwo, terraFactThree)
			return funFactToReturn.Terra
		default:
			return []string{"Invalid argument"}
		}
	case "F":
		switch about {
		case "sun":
			sunFactOne := "Temperatur på ytre lag av Solen er " + getFunfactTemperature(sunSurfaceTemp, "K", "F")
			sunFactTwo := "Temperatur i Solens kjerne er " + getFunfactTemperature(sunCoreTemp, "C", "F")
			funFactToReturn.Sun = append(funFactToReturn.Sun, sunFactOne, sunFactTwo)
			return funFactToReturn.Sun
		case "luna":
			lunaFactOne := "Temperatur på månens overflate om natten er " + getFunfactTemperature(moonSurfaceTempNight, "C", "F")
			lunaFactTwo := "Temperatur på månens overflate om dagen er " + getFunfactTemperature(moonSurfaceTempDay, "C", "F")
			funFactToReturn.Luna = append(funFactToReturn.Luna, lunaFactOne, lunaFactTwo)
			return funFactToReturn.Luna
		case "terra":
			terraFactOne := "Laveste målte Temperatur på jordens overflate er " + getFunfactTemperature(lowestTempSurface, "C", "F")
			terraFactTwo := "Høyeste temperatur målt på Jordens overflate " + getFunfactTemperature(highestTempSurface, "C", "F")
			terraFactThree := "Temperatur i jordens kjerne er " + getFunfactTemperature(earthCoreTemp, "K", "F")
			funFactToReturn.Terra = append(funFactToReturn.Terra, terraFactOne, terraFactTwo, terraFactThree)
			return funFactToReturn.Terra
		default:
			return []string{"Invalid argument"}
		}
	case "K":
		switch about {
		case "sun":
			sunFactOne := "Temperatur på ytre lag av Solen er " + getFunfactTemperature(sunSurfaceTemp, "K", "K")
			sunFactTwo := "Temperatur i Solens kjerne er " + getFunfactTemperature(sunCoreTemp, "C", "K")
			funFactToReturn.Sun = append(funFactToReturn.Sun, sunFactOne, sunFactTwo)
			return funFactToReturn.Sun

		case "luna":
			lunaFactOne := "Temperatur på månens overflate om natten er " + getFunfactTemperature(moonSurfaceTempNight, "C", "K")
			lunaFactTwo := "Temperatur på månens overflate om dagen er " + getFunfactTemperature(moonSurfaceTempDay, "C", "K")
			funFactToReturn.Luna = append(funFactToReturn.Luna, lunaFactOne, lunaFactTwo)
			return funFactToReturn.Luna
		case "terra":
			terraFactOne := "Laveste målte Temperatur på jordens overflate er " + getFunfactTemperature(lowestTempSurface, "C", "K")
			terraFactTwo := "Høyeste temperatur målt på Jordens overflate " + getFunfactTemperature(highestTempSurface, "C", "K")
			terraFactThree := "Temperatur i jordens kjerne er " + getFunfactTemperature(earthCoreTemp, "K", "K")
			funFactToReturn.Terra = append(funFactToReturn.Terra, terraFactOne, terraFactTwo, terraFactThree)
			return funFactToReturn.Terra
		default:
			return []string{"Invalid temperature scale"}
		}
	default:
		return []string{"Invalid temperature scale"}
	}
}

func getFunfactTemperature(temp float64, inputTemp string, outputScale string) string {
	var convertedTemp float64

	switch inputTemp {
	case "C":
		if outputScale == "K" {
			convertedTemp = conv.CelsiusToKelvin(temp)
		} else if outputScale == "F" {
			convertedTemp = conv.CelsiusToFahrenheit(temp)
		} else if outputScale == "C" {
			convertedTemp = temp
		}
	case "F":
		if outputScale == "C" {
			convertedTemp = conv.FahrenheitToCelsius(temp)
		} else if outputScale == "K" {
			convertedTemp = conv.FahrenheitToKelvin(temp)
		} else if outputScale == "F" {
			convertedTemp = temp
		}
	case "K":
		if outputScale == "C" {
			convertedTemp = conv.KelvinToCelsius(temp)
		} else if outputScale == "F" {
			convertedTemp = conv.KelvinToFahrenheit(temp)
		} else if outputScale == "K" {
			convertedTemp = temp
		}
	}

	outputTemp := fmt.Sprintf("%.2f", convertedTemp)
	if strings.HasSuffix(outputTemp, ".00") {
		outputTemp = outputTemp[:len(outputTemp)-3]
	} else if strings.HasSuffix(outputTemp, "0") {
		outputTemp = outputTemp[:len(outputTemp)-1]
	}

	// add spaces between every three digits in the converted temperature
	parts := strings.Split(outputTemp, ".")
	intPart := parts[0]
	var formattedIntPart string
	for i := len(intPart) - 1; i >= 0; i-- {
		formattedIntPart = string(intPart[i]) + formattedIntPart
		if (len(intPart)-i)%3 == 0 && i != 0 {
			formattedIntPart = " " + formattedIntPart
		}
	}
	if outputScale == "C" || outputScale == "F" {
		outputScale = "°" + outputScale
	}
	if len(parts) == 1 {
		return fmt.Sprintf("%s%s", formattedIntPart, outputScale)
	}
	return fmt.Sprintf("%s.%s%s", formattedIntPart, parts[1], strings.TrimSpace(outputScale))

}
