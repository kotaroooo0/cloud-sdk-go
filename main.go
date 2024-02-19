package main

import (
	"net/http"

	"github.com/kotaroooo0/cloud-sdk-go/pkg/api"
	"github.com/kotaroooo0/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/kotaroooo0/cloud-sdk-go/pkg/auth"
)

func main() {
	ecApiKey := "your-api"
	ecClient, err := api.NewAPI(api.Config{
		Client:     new(http.Client),
		AuthWriter: auth.APIKey(ecApiKey),
	})
	// if err != nil {
	// 	return nil, err
	// }
	// デプロイメントの一覧を取得
	res, err := deploymentapi.List(deploymentapi.ListParams{API: ecClient})
	if err != nil {
		return nil, err
	}

}
