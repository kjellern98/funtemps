package main

import (
	"flag"
	"fmt"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64
var out string
var funfacts string
var temperatureScale string

func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in degrees kelvin")

	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

	flag.StringVar(&temperatureScale, "t", "C", "temperaturskala som skal brukes når fun-facts vises")

}

func main() {

	flag.Parse()

	// Check for conflicting flags
	if isFlagPassed("F") && isFlagPassed("C") {
		flagError("-F", "-C")
		return
	} else if isFlagPassed("F") && isFlagPassed("K") {
		flagError("-F", "-K")
		return
	} else if isFlagPassed("C") && isFlagPassed("K") {
		flagError("-C", "-K")
		return
	} else if isFlagPassed("funfacts") && (isFlagPassed("F") || isFlagPassed("C") || isFlagPassed("K")) {
		fmt.Println("Error: -funfacts cannot be used with -F, -C, or -K flags")
		return
	} else if isFlagPassed("funfacts") && !isFlagPassed("t") {
		fmt.Println("Error: -funfacts must be used with -t flag")
		return
	}

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	if isFlagPassed("F") && isFlagPassed("out") {
		result := FahrenheitToCelsius(fahr)
		fmt.Printf("%0.2f°F is %0.2f°C\n", fahr, result)
	} else if isFlagPassed("C") && isFlagPassed("out") {
		result := CelsiusToFahrenheit(celsius)
		fmt.Printf("%0.2f°C is %0.2f°F\n", celsius, result)
	} else if isFlagPassed("K") && isFlagPassed("out") {
		result := KelvinToCelsius(kelvin)
		fmt.Printf("%0.2f°K is %0.2f°C\n", kelvin, result)
	} else if isFlagPassed("C") && isFlagPassed("fact") {
		result := funfacts.RandomCelsiusFact()
		fmt.Println(result)
	} else if isFlagPassed("F") && isFlagPassed("fact") {
		result := funfacts.RandomFahrenheitFact()
		fmt.Println(result)
	} else if isFlagPassed("K") && isFlagPassed("fact") {
		result := funfacts.RandomKelvinFact()
		fmt.Println(result)
	} else if isFlagPassed("conv") {
		result := conv.Convert(value, unit)
		fmt.Printf("%0.2f %s is %0.2f %s\n", value, unit, result.Value, result.Unit)
	} else {
		flag.Usage()
	}

	/*
		// Her er noen eksempler du kan bruke i den manuelle testingen
		fmt.Println(fahr, out, funfacts)

		fmt.Println("len(flag.Args())", len(flag.Args()))
		fmt.Println("flag.NFlag()", flag.NFlag())

		fmt.Println(isFlagPassed("out"))

		// Eksempel på enkel logikk
		if out == "C" && isFlagPassed("F") {
			// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
			// skal returnere °C
			fmt.Println("0°F er -17.78°C")
		}
	*/
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func flagError(flag1, flag2 string) {
	fmt.Printf("Error: %s and %s flags cannot be used simultaneously\n", flag1, flag2)
	return
}
