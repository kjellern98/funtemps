package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input  string
		output string
		want   []string
	}

	var sunFunFactsCelcius = FunFacts{}
	var sunFunFactKelvin = FunFacts{}
	var sunFunfactFeherenheit = FunFacts{}

	var moonFunFactCelcius = FunFacts{}
	var moonFunFactKelvin = FunFacts{}
	var moonFunFactFeherenheit = FunFacts{}

	var terraFunFactCelcius = FunFacts{}
	var terraFunFactKelvin = FunFacts{}
	var terraFunFactFeherenheit = FunFacts{}

	sunFactOneCelcius := "Temperatur på ytre lag av Solen er 5 504.85°C"
	sunFactTwoCelcius := "Temperatur i Solens kjerne er 15 000 000°C"
	sunFunFactsCelcius.Sun = append(sunFunFactsCelcius.Sun, sunFactOneCelcius, sunFactTwoCelcius)

	sunFactOneKelvin := "Temperatur på ytre lag av Solen er 5 778K"
	sunFactTwoKelvin := "Temperatur i Solens kjerne er 15 000 273.15K"
	sunFunFactKelvin.Sun = append(sunFunFactKelvin.Sun, sunFactOneKelvin, sunFactTwoKelvin)

	sunFactOneFeherenheit := "Temperatur på ytre lag av Solen er 9 940.73°F"
	sunFactTwoFeherenheit := "Temperatur i Solens kjerne er 27 000 032°F"
	sunFunfactFeherenheit.Sun = append(sunFunfactFeherenheit.Sun, sunFactOneFeherenheit, sunFactTwoFeherenheit)

	lunaFactOneCelcius := "Temperatur på månens overflate om natten er - 183°C"
	lunaFactTwoCelcius := "Temperatur på månens overflate om dagen er 106°C"
	moonFunFactCelcius.Luna = append(moonFunFactCelcius.Luna, lunaFactOneCelcius, lunaFactTwoCelcius)

	lunaFactOneKelvin := "Temperatur på månens overflate om natten er 90.15K"
	lunaFactTwoKelvin := "Temperatur på månens overflate om dagen er 379.15K"
	moonFunFactKelvin.Luna = append(moonFunFactKelvin.Luna, lunaFactOneKelvin, lunaFactTwoKelvin)

	lunaFactOneFeherenheit := "Temperatur på månens overflate om natten er - 297.4°F"
	lunaFactTwoFeherenheit := "Temperatur på månens overflate om dagen er 222.8°F"
	moonFunFactFeherenheit.Luna = append(moonFunFactFeherenheit.Luna, lunaFactOneFeherenheit, lunaFactTwoFeherenheit)

	terraFactOneCelcius := "Laveste målte Temperatur på jordens overflate er -89.4°C"
	terraFactTwoCelcius := "Høyeste temperatur målt på Jordens overflate 56.7°C"
	terraFactThreeCelcius := "Temperatur i jordens kjerne er 9 118.85°C"
	terraFunFactCelcius.Terra = append(terraFunFactCelcius.Terra, terraFactOneCelcius, terraFactTwoCelcius, terraFactThreeCelcius)

	terraFactOneKelvin := "Laveste målte Temperatur på jordens overflate er 183.75K"
	terraFactTwoKelvin := "Høyeste temperatur målt på Jordens overflate 329.85K"
	terraFactThreeKelvin := "Temperatur i jordens kjerne er 9 392K"
	terraFunFactKelvin.Terra = append(terraFunFactKelvin.Terra, terraFactOneKelvin, terraFactTwoKelvin, terraFactThreeKelvin)

	terraFactOneFeherenheit := "Laveste målte Temperatur på jordens overflate er - 128.92°F"
	terraFactTwoFeherenheit := "Høyeste temperatur målt på Jordens overflate 134.06°F"
	terraFactThreeFeherenheit := "Temperatur i jordens kjerne er 16 445.93°F"
	terraFunFactFeherenheit.Terra = append(terraFunFactFeherenheit.Terra, terraFactOneFeherenheit, terraFactTwoFeherenheit, terraFactThreeFeherenheit)

	tests := []test{
		{input: "sun", output: "C", want: sunFunFactsCelcius.Sun},
		{input: "sun", output: "K", want: sunFunFactKelvin.Sun},
		{input: "sun", output: "F", want: sunFunfactFeherenheit.Sun},
		{input: "luna", output: "C", want: moonFunFactCelcius.Luna},
		{input: "luna", output: "K", want: moonFunFactKelvin.Luna},
		{input: "luna", output: "F", want: moonFunFactFeherenheit.Luna},
		{input: "terra", output: "C", want: terraFunFactCelcius.Terra},
		{input: "terra", output: "K", want: terraFunFactKelvin.Terra},
		{input: "terra", output: "F", want: terraFunFactFeherenheit.Terra},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input, tc.output)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
