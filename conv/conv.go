package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	// Returnerer Kelvin
	return celsius + 273.15
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
