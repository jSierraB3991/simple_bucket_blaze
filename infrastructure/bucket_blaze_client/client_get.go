package bucketblazeclient

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	bucketblazerequest "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_request"
	bucketblazeresponse "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_response"
)

func (HttpClient) Get(urlBase, uri string, result interface{}, headers []bucketblazerequest.HeaderRequest) error {
	req, err := http.NewRequest("GET", urlBase+uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for _, v := range headers {
			req.Header.Add(v.Key, v.Value)
		}
	}

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		//body, _ := io.ReadAll(resp.Body)

		//fmt.Println(resp)
		//fmt.Println(string(body))
		return json.NewDecoder(resp.Body).Decode(&result)
	} else {
		var responseEror bucketblazeresponse.InvalidRequestPayResponse
		err = json.NewDecoder(resp.Body).Decode(&responseEror)
		if err != nil {
			return err
		}
		log.Println(responseEror)
		return errors.New(responseEror.CodeError)
	}

}
