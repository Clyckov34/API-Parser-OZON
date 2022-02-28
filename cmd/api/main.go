package main

import (
	"flag"
	"log"
	"myapiproject/internal/app"
	"net/http"
	"os"
)

var getFlags app.AppFlag

func init() {
	flag.StringVar(&getFlags.API.ApiKey, "apiKey", "", "Данные Header Api-Key")
	flag.StringVar(&getFlags.API.ClientId, "clientId", "", "Данные Header Client-Id")
	flag.StringVar(&getFlags.API.URL, "url", "https://api-seller.ozon.ru/v2/category/tree", "URL API")
	flag.StringVar(&getFlags.API.Method, "method", http.MethodPost, "Метод запроса")
	flag.BoolVar(&getFlags.ExportFile, "export", false, "Выгружает json файл товаров")
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("error: No flags specified")
	} else {
		flag.Parse()
		app.Run(&getFlags)
	}
}
