package main

import (
	"flag"
	"log"
	"myapiproject/internal/api"
	"net/http"
	"os"
)

var dataFlag api.AppFlag

func init() {
	flag.StringVar(&dataFlag.Header.ApiKey, "apiKey", "", "Данные Header Api-Key")
	flag.StringVar(&dataFlag.Header.ClientId, "clientId", "", "Данные Header Client-Id")
	flag.StringVar(&dataFlag.Header.URL, "url", "https://api-seller.ozon.ru/v2/category/tree", "URL API")
	flag.StringVar(&dataFlag.Header.Method, "method", http.MethodPost, "Метод запроса")
	flag.BoolVar(&dataFlag.ExpFile, "expFile", false, "Выгружает json файл товаров")
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("ошибка: Не указаны флаги")
	} else {
		flag.Parse()
		api.Run(&dataFlag)
	}
}
