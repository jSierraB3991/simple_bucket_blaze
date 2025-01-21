package bucketblazeservice

import (
	"encoding/base64"
	"fmt"

	bucketblazeclient "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_client"
)

type BucketBlazeService struct {
	apiKey     string
	keyId      string
	urlApi     string
	httpClient *bucketblazeclient.HttpClient
}

func New(apkiKey, keyId, urlApi string) *BucketBlazeService {
	return &BucketBlazeService{
		apiKey:     apkiKey,
		keyId:      keyId,
		urlApi:     urlApi,
		httpClient: bucketblazeclient.NewHttpClient(),
	}
}

func (s *BucketBlazeService) getAuthorizationData() string {
	datakey := fmt.Sprintf("%s:%s", s.keyId, s.apiKey)
	dataBytes := []byte(datakey)
	dataString := base64.StdEncoding.EncodeToString(dataBytes)
	return "Basic " + dataString
}
