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
	var client = api.HTTPHeader{
		ApiKey:      fl.HTTPHeader.ApiKey,
		ClientId:    fl.HTTPHeader.ClientId,
		ContentType: "application/json",
		URL:         fl.HTTPHeader.URL,
		Method:      fl.HTTPHeader.Method,
		WithTimeout: 15 * time.Second,
	}

	//Запрос в API
	body, err := client.RequestClient()
	if err != nil {
		log.Fatalln(err)
	}

	//Обработка полученных данных
	var goods = convert.Categories{}
	data, err := goods.DecoderBodyJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	//Конвертирование данные в определенную структуру
	goods.ConvertData(*data)
	dataGoods := goods.Goods()

	//Вывод данных в терминал
	for k, v := range dataGoods {
		fmt.Printf("%v) %v - %v - %v\n", k+1, v.Id, v.Name, v.Parents)
	}

	//Выгрузка данных
	if fl.ExpFile {
		if err := export.CreateJsonFile(dataGoods); err != nil {
			log.Println(err)
		}
	}
}
