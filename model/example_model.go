package model

type ExampleRequest struct {
	AppName string `json:"app_name"`
	AppVer  string `json:"app_ver"`
}

type ExampleResponse struct {
	AppName string `json:"app_name"`
	AppVer  string `json:"app_ver"`
}
