package main

import (
	"DealGenerator/src/model"
	"DealGenerator/src/util"
	"encoding/json"
	"math"
	"math/rand"
	"os"
)

const (
	Home_File_Dir = "/Users/Nino/.go/src/DealGenerator"
	Generated_Deals_File = Home_File_Dir + "/generated-deals.json"
	Locations_File = Home_File_Dir + "/locations.json"
	Categories_Local_File = Home_File_Dir + "/categories-local.json"
	Categories_Shopping_File = Home_File_Dir + "/categories-shopping.json"
	Categories_Travel_File = Home_File_Dir + "/categories-travel.json"
	Merchants_Local_File = Home_File_Dir + "/merchants-local.json"
	Merchants_Shopping_File = Home_File_Dir + "/merchants-shopping.json"
	Merchants_Travel_File = Home_File_Dir + "/merchants-travel.json"
	Data_Local_File = Home_File_Dir + "/data-local.json"
	Data_Shopping_File = Home_File_Dir + "/data-shopping.json"
	Data_Travel_File = Home_File_Dir + "/data-travel.json"
	Max_Price_Limit = 30000
	Deals_To_Generate float64 = 1000
)


func main() {

	var dealsPerVertical = math.Ceil(Deals_To_Generate / 3)

	locations := getLocations()
	merchantsTravel, merchantsShopping, merchantsLocal := getMerchants()
	categoriesTravel, categoriesShopping, categoriesLocal := getCategories()
	dataTravel, dataShopping, dataLocal := getData()

	var maxDeals = int(dealsPerVertical)
	var dealsTravel, dealsShopping, dealsLocal = make([]model.Deal, maxDeals), make([]model.Deal, maxDeals), make([]model.Deal, maxDeals)

	generateDeals(maxDeals, dealsTravel, categoriesTravel, locations, merchantsTravel, dataTravel)
	generateDeals(maxDeals, dealsShopping, categoriesShopping, locations, merchantsShopping, dataShopping)
	generateDeals(maxDeals, dealsLocal, categoriesLocal, locations, merchantsLocal, dataLocal)

	var allDeals = append(append(dealsTravel, dealsShopping...), dealsLocal...)
	saveDealsAsJsonFile(allDeals)
}

func saveDealsAsJsonFile(allDeals []model.Deal) {
	f, err := os.Create(Generated_Deals_File)
	util.CheckErr(err)
	defer f.Close()
	f.WriteString("[\n")
	for i, deal := range allDeals {
		dealJson, _ := json.Marshal(deal)
		f.WriteString("\t")
		_, err = f.Write(dealJson)
		util.CheckErr(err)
		if len(allDeals) - 1 == i {
			f.WriteString("\n")
		} else {
			f.WriteString(",\n")
		}
	}
	f.WriteString("]")
	f.Sync()
}

func generateDeals(maxDeals int, dealsArray []model.Deal, categories []model.Category, locations []model.Location, merchants []model.Merchant, data []model.Data) {
	for i := 0; i < maxDeals; i++ {
		dat := data[rand.Intn(len(data))]
		merchant := merchants[rand.Intn(len(merchants))]
		category := categories[rand.Intn(len(categories))]
		location := locations[rand.Intn(len(locations))]

		dealsArray[i] = model.Deal{
			ID:       rand.Intn(maxDeals * 3),
			Image:    dat.Image,
			Title:    dat.Title,
			Price:    rand.Intn(Max_Price_Limit),
			Merchant: merchant.Name,
			Category: category.Name,
			Location: location.Name,
		}
	}
}

func getCategories() (categoriesTravel, categoriesShopping, categoriesLocal []model.Category) {
	rawTravel := util.ReadJsonFile(Categories_Travel_File)
	rawShopping := util.ReadJsonFile(Categories_Shopping_File)
	rawLocal := util.ReadJsonFile(Categories_Local_File)

	json.Unmarshal(rawTravel, &categoriesTravel)
	json.Unmarshal(rawShopping, &categoriesShopping)
	json.Unmarshal(rawLocal, &categoriesLocal)

	return
}

func getMerchants() (merchantTravel, merchantShopping, merchantLocal []model.Merchant) {
	rawTravel := util.ReadJsonFile(Merchants_Travel_File)
	rawShopping := util.ReadJsonFile(Merchants_Shopping_File)
	rawLocal := util.ReadJsonFile(Merchants_Local_File)

	json.Unmarshal(rawTravel, &merchantTravel)
	json.Unmarshal(rawShopping, &merchantShopping)
	json.Unmarshal(rawLocal, &merchantLocal)

	return
}

func getData() (dataTravel, dataShopping, dataLocal []model.Data) {
	rawTravel := util.ReadJsonFile(Data_Travel_File)
	rawShopping := util.ReadJsonFile(Data_Shopping_File)
	rawLocal := util.ReadJsonFile(Data_Local_File)

	json.Unmarshal(rawTravel, &dataTravel)
	json.Unmarshal(rawShopping, &dataShopping)
	json.Unmarshal(rawLocal, &dataLocal)

	return
}

func getLocations() []model.Location {
	raw := util.ReadJsonFile(Locations_File)
	var locations []model.Location
	json.Unmarshal(raw, &locations)
	return locations
}
