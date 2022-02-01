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
	flag.StringVar(&dataFlag.HTTPHeader.ApiKey, "apiKey", "", "Данные Header Api-Key")
	flag.StringVar(&dataFlag.HTTPHeader.ClientId, "clientId", "", "Данные Header Client-Id")
	flag.StringVar(&dataFlag.HTTPHeader.URL, "url", "https://api-seller.ozon.ru/v2/category/tree", "URL API")
	flag.StringVar(&dataFlag.HTTPHeader.Method, "method", http.MethodPost, "Метод запроса")
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
