package main

import "github.com/brianvoe/gofakeit/v6"

func main() {

	println(gofakeit.Name())          // Markus Moen
	println(gofakeit.Email())         // alaynawuckert@kozey.biz
	println(gofakeit.Phone())         // (570)245-7485
	println(gofakeit.BS())            // front-end
	println(gofakeit.BeerName())      // Duvel
	println(gofakeit.Color())         // MediumOrchid
	println(gofakeit.Company())       // Moen, Pagac and Wuckert
	println(gofakeit.HackerPhrase())  // Connecting the array won't do anything, we need to generate the haptic COM driver!
	println(gofakeit.JobTitle())      // Director
	println(gofakeit.CurrencyShort()) // USD
}
