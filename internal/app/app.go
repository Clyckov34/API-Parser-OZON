package app

import (
	"fmt"
	"log"
	"myapiproject/internal/api"
	"myapiproject/internal/convert"
	"myapiproject/internal/export"
	"time"
)

func Run(fl *AppFlag) {
	var client = api.API{
		ApiKey:      fl.API.ApiKey,
		ClientId:    fl.API.ClientId,
		ContentType: "application/json",
		URL:         fl.API.URL,
		Method:      fl.API.Method,
		WithTimeout: 15 * time.Second,
	}

	//Запрос в API
	body, err := client.GetRequest()
	if err != nil {
		log.Fatalln(err)
	}

	//Обработка полученных данных
	var goods = convert.Categories{}
	data, err := goods.SetDecoderBodyJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	//Конвертирование данные в определенную структуру
	goods.SetConvert(*data)
	dataGoods := goods.GetGoods()

	//Вывод данных в терминал
	for k, v := range dataGoods {
		fmt.Printf("%v) %v - %v - %v\n", k+1, v.Id, v.Name, v.Parents)
	}

	//Выгрузка данных
	if fl.ExportFile {
		if err := export.GetJsonFile(dataGoods); err != nil {
			log.Println(err)
		}
	}
}
