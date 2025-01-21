package bucketblazeresponse

import "time"

type AutorizeAccountResponse struct {
	AccountID                         string                         `json:"accountId"`
	APIInfo                           AutorizeAccountResponseAppInfo `json:"apiInfo"`
	ApplicationKeyExpirationTimestamp *time.Time                     `json:"applicationKeyExpirationTimestamp"`
	AuthorizationToken                string                         `json:"authorizationToken"`
}

type AutorizeAccountResponseAppInfo struct {
	StorageAPI AutorizeAccountResponseAppInfoStorageApi `json:"storageApi"`
}

type AutorizeAccountResponseAppInfoStorageApi struct {
	AbsoluteMinimumPartSize int      `json:"absoluteMinimumPartSize"`
	APIURL                  string   `json:"apiUrl"`
	BucketID                any      `json:"bucketId"`
	BucketName              any      `json:"bucketName"`
	Capabilities            []string `json:"capabilities"`
	DownloadURL             string   `json:"downloadUrl"`
	InfoType                string   `json:"infoType"`
	NamePrefix              any      `json:"namePrefix"`
	RecommendedPartSize     int      `json:"recommendedPartSize"`
	S3APIURL                string   `json:"s3ApiUrl"`
}
