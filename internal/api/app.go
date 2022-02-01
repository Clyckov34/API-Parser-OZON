package api

import (
	"fmt"
	"log"
)

func Run(fl *AppFlag) {
	var client Client = &Headers{
		ApiKey:   fl.Header.ApiKey,
		ClientId: fl.Header.ClientId,
		URL:      fl.Header.URL,
		Method:   fl.Header.Method,
	}

	body, err := client.RequestClient()
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.DecoderBodyJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	dataGoods := &DataCategories
	client.ConvertData(*response)

	if fl.ExpFile == true {
		client.CreateJsonFile(*dataGoods)
	}

	for k, v := range *dataGoods {
		fmt.Printf("%v) %v - %v - %v\n", k+1, v.Id, v.Name, v.Parents)
	}

}
