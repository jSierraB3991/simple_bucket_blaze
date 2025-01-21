package bucketblazeservice

import (
	bucketblazelibs "github.com/jSierraB3991/simple_bucket_blaze/domain/bucket_blaze_libs"
	bucketblazerequest "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_request"
	bucketblazeresponse "github.com/jSierraB3991/simple_bucket_blaze/infrastructure/bucket_blaze_response"
)

func (s *BucketBlazeService) Autorize() (*bucketblazeresponse.AutorizeAccountResponse, error) {
	authData := s.getAuthorizationData()

	headers := []bucketblazerequest.HeaderRequest{
		{Key: "Authorization", Value: authData},
	}
	var data bucketblazeresponse.AutorizeAccountResponse
	err := s.httpClient.Get(s.urlApi, bucketblazelibs.AUTORIZE_URL, &data, headers)
	if err != nil {
		return nil, err
	}
	return &data, err
}
