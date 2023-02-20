package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kjellern98/funtemps/conv"
	"github.com/kjellern98/funtemps/funfacts"
)

var fahr float64
var celsius float64
var kelvin float64
var out string
var funfact string
var temperatureScale string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "Input in Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "Input in Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "Input in Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&temperatureScale, "t", "C", "temperaturskala som skal brukes når fun-facts vises")

}

func main() {
	flag.Parse()

	/*
		fmt.Println("Input F = ", fahr)
		fmt.Println("Input C =", celsius)
		fmt.Println("Input K =", kelvin)
		fmt.Println("Output Type =", out)
	*/

	switch {
	case isFlagPassed("F"):
		if isFlagPassed("C") || isFlagPassed("K") {
			fmt.Println("-F, -C, -K cannot be used together")
			return
		} else if isFlagPassed("funfacts") {
			fmt.Println("-F, -funfacts cannot be used together")
			return
		}
		switch out {
		case "C":
			fmt.Println(convertTemperature(fahr, "F", "C"))
		case "K":
			fmt.Println(convertTemperature(fahr, "F", "K"))
		case "F":
			fmt.Println("cannot convert to same temperature scale")
			fmt.Println("Try using -out C or -out K instead")
		default:
			fmt.Println("Invalid output type")
		}
	case isFlagPassed("C"):
		if isFlagPassed("F") || isFlagPassed("K") {
			fmt.Println("-F, -C, -K cannot be used together")
			return
		} else if isFlagPassed("funfacts") {
			fmt.Println("-F, -funfacts cannot be used together")
			return
		}
		switch out {
		case "F":
			fmt.Println(convertTemperature(celsius, "C", "F"))
		case "K":
			fmt.Println(convertTemperature(celsius, "C", "K"))
		case "C":
			fmt.Println("cannot convert to same temperature scale")
			fmt.Println("Try using -out F or -out K instead")
		default:
			fmt.Println("Invalid output type")
		}
	case isFlagPassed("K"):
		if isFlagPassed("F") || isFlagPassed("C") {
			fmt.Println("-F, -C, -K cannot be used together")
			return
		} else if isFlagPassed("funfacts") {
			fmt.Println("-F, -funfacts cannot be used together")
			return
		}
		switch out {
		case "F":
			fmt.Println(convertTemperature(kelvin, "K", "F"))
		case "C":
			fmt.Println(convertTemperature(kelvin, "K", "C"))
		case "K":
			fmt.Println("cannot convert to same temperature scale")
			fmt.Println("Try using -out C or -out F instead")
		default:
			fmt.Println("Invalid output type")
		}
	case isFlagPassed("funfacts"):
		if !isFlagPassed("t") {
			fmt.Println("Temperature scale must be specified when using fun-facts")
			return
		}
		switch funfact {
		case "sun":
			var sunFacts = funfacts.GetFunFacts("sun", temperatureScale)
			for _, fact := range sunFacts {
				fmt.Println(fact)
			}
		case "luna":
			var lunaFacts = funfacts.GetFunFacts("luna", temperatureScale)
			for _, fact := range lunaFacts {
				fmt.Println(fact)
			}
		case "terra":
			var terraFacts = funfacts.GetFunFacts("terra", temperatureScale)
			for _, fact := range terraFacts {
				fmt.Println(fact)
			}
		default:
			fmt.Println("Invalid fun-facts name")
		}
	default:
		fmt.Println("No input given")
		fmt.Println("Example to try: go run main.go -F 100 -out C")
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func convertTemperature(temp float64, inputScale string, outputScale string) string {
	var convertedTemp float64

	switch inputScale {
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
	if inputScale == "C" || inputScale == "F" {
		inputScale = "°" + inputScale
	}

	if len(parts) == 1 {
		return fmt.Sprintf("%g%s is %s%s", temp, inputScale, formattedIntPart, outputScale)
	}
	return fmt.Sprintf("%g%s is %s.%s%s", temp, inputScale, formattedIntPart, parts[1], outputScale)
}
