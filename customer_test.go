package main

import (
	"fmt"
	"testing"
)

type CustomerFavouriteTest struct {
	customers      []Customer
	expectedResult []TopCustomerFavourite
}

func TestCustomerFavourite(t *testing.T) {
	customerSet1 := []Customer{
		{name: "John", candy: "Gum", eaten: 1},
		{name: "Johan", candy: "Gum", eaten: 12},
		{name: "John", candy: "Butterfinger", eaten: 26},
		{name: "Anna", candy: "Gum", eaten: 50},
		{name: "Anna", candy: "Butterfinger", eaten: 3},
		{name: "Johan", candy: "Jelly", eaten: 70},
		{name: "John", candy: "Nestle", eaten: 200},
	}
	expectedResult1 := []TopCustomerFavourite{
		{Name: "John", FavouriteSnack: "Nestle", TotalSnacks: 227},
		{Name: "Johan", FavouriteSnack: "Jelly", TotalSnacks: 82},
		{Name: "Anna", FavouriteSnack: "Gum", TotalSnacks: 53},
	}

	customerSet2 := []Customer{
		{name: "Annika", candy: "Geisha", eaten: 1},
		{name: "Jonas", candy: "Geisha", eaten: 70},
		{name: "Annika", candy: "Kexchoklad", eaten: 96},
		{name: "Annika", candy: "Nötchoklad", eaten: 24},
		{name: "Jonas", candy: "Nötchoklad", eaten: 50},
		{name: "Aadya", candy: "Center", eaten: 208},
		{name: "Aadya", candy: "Nötchoklad", eaten: 208},
		{name: "Annika", candy: "Geisha", eaten: 12},
		{name: "Jonas", candy: "Geisha", eaten: 3},
		{name: "Jane", candy: "Plopp", eaten: 230},
		{name: "Jane", candy: "Center", eaten: 230},
		{name: "Jane", candy: "Geisha", eaten: 27},
		{name: "Aadya", candy: "Nötchoklad", eaten: 102},
		{name: "Aadya", candy: "Center", eaten: 40},
		{name: "Jane", candy: "Geisha", eaten: 190},
		{name: "Jonas", candy: "Nötchoklad", eaten: 200},
		{name: "Annika", candy: "Nötchoklad", eaten: 1},
	}
	expectedResult2 := []TopCustomerFavourite{
		{Name: "Jane", FavouriteSnack: "Center, Plopp", TotalSnacks: 677},
		{Name: "Aadya", FavouriteSnack: "Nötchoklad", TotalSnacks: 558},
		{Name: "Jonas", FavouriteSnack: "Nötchoklad", TotalSnacks: 323},
		{Name: "Annika", FavouriteSnack: "Kexchoklad", TotalSnacks: 134},
	}

	testVals := []CustomerFavouriteTest{
		{customers: customerSet1, expectedResult: expectedResult1},
		{customers: customerSet2, expectedResult: expectedResult2},
		{customers: []Customer{}, expectedResult: []TopCustomerFavourite{}},
	}

	for index, testVal := range testVals {
		t.Run(fmt.Sprintf("test #%v", index), func(subT *testing.T) {
			if t.Failed() {
				subT.SkipNow()
			}
			result := findTopCustomerFavourite(testVal.customers)
			if !compareTopCustomerFavourite(result, testVal.expectedResult) {
				subT.Fatalf("Expected %v, got %v, so failed", testVal.expectedResult, result)
			}
		})
	}
}

func compareTopCustomerFavourite(result []TopCustomerFavourite, expectedResult []TopCustomerFavourite) bool {
	if len(result) != len(expectedResult) {
		return false
	}

	for index, resultVal := range result {
		if resultVal.Name != expectedResult[index].Name {
			return false
		}
		if resultVal.FavouriteSnack != expectedResult[index].FavouriteSnack {
			return false
		}
		if resultVal.TotalSnacks != expectedResult[index].TotalSnacks {
			return false
		}
	}
	return true
}
