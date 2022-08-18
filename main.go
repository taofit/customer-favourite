package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Customer struct {
	name  string
	candy string
	eaten uint
}

type TopCustomerFavourite struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    uint   `json:"totalSnacks"`
}

func findMaxEatenNum(candy map[string]uint) (max uint) {
	max = 0
	for _, num := range candy {
		if num >= max {
			max = num
		}
	}
	return
}

func getFavouriteSnack(candyToEaten map[string]uint) (favouriteSnackSlice []string) {
	max := findMaxEatenNum(candyToEaten)
	for name, num := range candyToEaten {
		if num == max {
			favouriteSnackSlice = append(favouriteSnackSlice, name)
		}
	}
	return
}

func findFavouriteSnack(customer []Customer) (favouriteSnack string) {
	candyToEaten := make(map[string]uint)

	for _, customerInfo := range customer {
		candyToEaten[customerInfo.candy] += customerInfo.eaten
	}
	favouriteSnackSlice := getFavouriteSnack(candyToEaten)

	if len(favouriteSnackSlice) > 1 {
		sort.Strings(favouriteSnackSlice)
		favouriteSnack = strings.Join(favouriteSnackSlice, ", ")
	}
	if len(favouriteSnackSlice) == 1 {
		favouriteSnack = favouriteSnackSlice[0]
	}
	return
}

func findTotal(customer []Customer) (total uint) {
	for _, customerInfo := range customer {
		total += customerInfo.eaten
	}
	return
}

func findTopCustomerFavourite(customers []Customer) (topCustomerFavourite []TopCustomerFavourite) {
	nameToCustomer := make(map[string][]Customer)
	for _, customer := range customers {
		nameToCustomer[customer.name] = append(nameToCustomer[customer.name], customer)
	}

	for name, customer := range nameToCustomer {
		favouriteSnack := findFavouriteSnack(customer)
		total := findTotal(customer)
		topCustomerFavourite = append(topCustomerFavourite, TopCustomerFavourite{name, favouriteSnack, total})

	}
	sort.SliceStable(topCustomerFavourite, func(i, j int) bool {
		return topCustomerFavourite[i].TotalSnacks > topCustomerFavourite[j].TotalSnacks
	})
	return
}

func main() {
	customers := []Customer{
		{"Annika", "Geisha", 100},
		{"Jonas", "Geisha", 200},
		{"Jonas", "Kexchoklad", 100},
		{"Aadya", "Nötchoklad", 2},
		{"Jonas", "Nötchoklad", 3},
		{"Jane", "Nötchoklad", 17},
		{"Annika", "Geisha", 100},
		{"Jonas", "Geisha", 700},
		{"Jane", "Nötchoklad", 4},
		{"Aadya", "Center", 7},
		{"Jonas", "Geisha", 900},
		{"Jane", "Nötchoklad", 1},
		{"Jonas", "Kexchoklad", 12},
		{"Jonas", "Plopp", 40},
		{"Jonas", "Center", 27},
		{"Aadya", "Center", 2},
		{"Annika", "Center", 8},
	}
	topCustomerFavourite := findTopCustomerFavourite(customers)
	res, err := json.MarshalIndent(topCustomerFavourite, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
