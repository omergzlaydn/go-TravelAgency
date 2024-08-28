package main

import "fmt"

// Servis türlerini temsil eden yapılar
type LuxService struct {
	seating          string
	foodAndBeverage  string
	climateControl   string
	entertainment    string
	additionalItems  string
	restStopService  string
}

type NormalService struct {
	seating          string
	foodAndBeverage  string
	climateControl   string
	entertainment    string
	restStopDuration string
}

// Lüks Sefer için fonksiyon
func (lux LuxService) DescribeService() {
	fmt.Println("Lüks Sefer Hizmeti:")
	fmt.Println("Oturma Düzeni:", lux.seating)
	fmt.Println("Yemek ve İçecek Servisi:", lux.foodAndBeverage)
	fmt.Println("Kişisel Klima:", lux.climateControl)
	fmt.Println("Eğlence Hizmetleri:", lux.entertainment)
	fmt.Println("Ekstra Ürünler:", lux.additionalItems)
	fmt.Println("Dinlenme Tesisi Hizmetleri:", lux.restStopService)
	fmt.Println()
}

// Normal Servis için fonksiyon
func (normal NormalService) DescribeService() {
	fmt.Println("Normal Servis Hizmeti:")
	fmt.Println("Oturma Düzeni:", normal.seating)
	fmt.Println("Yemek ve İçecek Servisi:", normal.foodAndBeverage)
	fmt.Println("Kişisel Klima:", normal.climateControl)
	fmt.Println("Eğlence Hizmetleri:", normal.entertainment)
	fmt.Println("Dinlenme Tesisi Mola Süresi:", normal.restStopDuration)
	fmt.Println()
}

func main() {
	// Lüks Sefer hizmetinin özelliklerini tanımla
	luxService := LuxService{
		seating:          "1+1 oturma düzeni",
		foodAndBeverage:  "Sınırsız yemek ve içecek servisi",
		climateControl:   "Kişiye özel ayarlanabilir klima",
		entertainment:    "Oyun konsolu (gamepass ve playstation plus dahil)",
		additionalItems:  "Terlik ve battaniye",
		restStopService:  "Ücretsiz dinlenme tesisi yemekleri",
	}

	// Normal Servis hizmetinin özelliklerini tanımla
	normalService := NormalService{
		seating:          "2+1 oturma düzeni",
		foodAndBeverage:  "İçecek servisi",
		climateControl:   "Klimalı otobüsler",
		entertainment:    "Kişiye özel televizyonlar ve dahili film, TV kanalları",
		restStopDuration: "15 dakikalık dinlenme tesisi molaları",
	}

	// Servisleri tanımla ve açıkla
	fmt.Println("AnadoluBaykuşları AŞ olarak iki farklı hizmet seçeneği sunuyoruz:\n")
	luxService.DescribeService()
	normalService.DescribeService()
}
