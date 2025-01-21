package bucketblazeservice_test

import (
	"os"
	"testing"

	bucketblazeservice "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_service"
)

func TestTestingConnectApi(t *testing.T) {
	apiKey := os.Getenv("BUCKET_BLAZE_API_KEY")
	keyId := os.Getenv("BUCKET_BLAZE_KEY_ID")
	urlApi := "https://api005.backblazeb2.com"

	service := bucketblazeservice.New(apiKey, keyId, urlApi)
	authResponse, err := service.Autorize()
	if err != nil {
		t.Errorf("I not want error, and recive %s", err.Error())
	}

	if authResponse.AccountID != keyId {
		t.Errorf("I Want %s and recive %s", keyId, authResponse.AccountID)
	}
}
