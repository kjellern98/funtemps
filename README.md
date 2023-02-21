# funtemps

Dette er et Go-program som kan konverterer mellom ulike temperaturenheter, det kan også vise "funFacts" om solen, månen eller jorden, avhengig av brukerens input.

Programmet tar inputArgumenter fra komandolinjen og flagg, for eksempel temperaturen som skal konverteres, temperaturskalaen og utdataformatet. Den behandler deretter inndata og enten konverterer temperaturen eller viser morsomme fakta, avhengig av flaggene som er gitt.

Konverteringen utføres ved hjelp av en konverteringspakke(conv) som inneholder funksjoner for å konvertere temperaturenheter mellom Celsius, Fahrenheit og Kelvin. Temperaturkonverteringen utføres basert på brukerens InputEnhet og ønsket UtEnhet. Hvis en ugyldig kombinasjon av inngangsflagg oppdages, returnerer programmet en feilmelding.

Hvis brukeren spesifiserer "funfacts"-flagget, viser programmet interessante fakta om solen, månen eller jorden. Det er da ulike temperaturfaktaer som vil bli vist med temperaturenheten som blir spesifisert av brukeren.

Til slutt sjekker programmet om brukeren har gitt noen input i det hele tatt og gir et eksempel på hvordan programmet skal brukes. Koden inkluderer også en funksjon for å sjekke om et flagg ble sendt inn av brukeren.

Eksempler på hvordan man kan bruke programmet:

./funtemps -K 273.15 -out C
Output: 273.15K er 0°C

./funtemps -C -89.4 -out F
Output: -89.4°C er -128.92°F

./funtemps -funfacts sun -t C
Output: (2 linjer)
Temperatur på ytre lag av Solen 5 506.85°C.
Temperatur i Solens kjerne er 15 000 000°C.
