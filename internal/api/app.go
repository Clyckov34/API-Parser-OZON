package api

import (
	"fmt"
	"log"
)

func Run(fl *AppFlag) {
	var client New = &HTTPHeader{
		ApiKey:   fl.HTTPHeader.ApiKey,
		ClientId: fl.HTTPHeader.ClientId,
		URL:      fl.HTTPHeader.URL,
		Method:   fl.HTTPHeader.Method,
	}

	body, err := client.RequestClient()
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.DecoderBodyJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	client.ConvertData(*response)
	data := client.CategoriesGoods()

	for k, v := range data {
		fmt.Printf("%v) %v - %v - %v\n", k+1, v.Id, v.Name, v.Parents)
	}

	if fl.ExpFile {
		if err := client.CreateJsonFile(data); err != nil {
			log.Println(err)
		}
	}
}
