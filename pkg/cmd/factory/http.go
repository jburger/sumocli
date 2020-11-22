package factory

import (
	"context"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	sl "github.com/wizedkyle/sumocli/api"
)

func NewHttpClient() (*sl.APIClient, context.Context) {
	credentials, endpoint := login.ReadCredentialsAndEndpoint()
	ctx := context.WithValue(context.Background(), sl.ContextBasicAuth, credentials)
	cfg := &sl.Configuration{
		BasePath:      endpoint,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	client := sl.NewAPIClient(cfg)
	return client, ctx
}

//func NewHttpRequest(method string, apiUrl string) (*http.Client, *http.Request) {
//	client := newHttpClient()
//	authToken, endpoint := login.GetCredentialsAndEndpoint()
//	request, _ := http.NewRequest(method, endpoint+apiUrl, nil)
//	request.Header.Add("Authorization", authToken)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("User-Agent", "Sumocli "+build.Version)
//	return client, request
//}
//
//func NewHttpRequestWithBody(method string, apiUrl string, body []byte) (*http.Client, *http.Request) {
//	client := newHttpClient()
//	authToken, endpoint := login.GetCredentialsAndEndpoint()
//	request, _ := http.NewRequest(method, endpoint+apiUrl, bytes.NewBuffer(body))
//	request.Header.Add("Authorization", authToken)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("User-Agent", "Sumocli "+build.Version)
//	return client, request
//}
//
//func HttpError(statusCode int, errorMessage []byte, logger zerolog.Logger) {
//	if statusCode == 400 {
//		var responseError api.ResponseError
//		jsonErr := json.Unmarshal(errorMessage, &responseError)
//		logging.LogError(jsonErr, logger)
//		fmt.Println(responseError.Errors[0].Message)
//	} else if statusCode == 401 {
//		fmt.Println("Unauthorized access please check the user exists,  are valid.")
//	} else if statusCode == 403 {
//		fmt.Println("This operation is not allowed for your account type or the user doesn't have the role capability to perform this action.")
//	} else if statusCode == 404 {
//		fmt.Println("Requested resource could not be found.")
//	} else if statusCode == 405 {
//		fmt.Println("Unsupported method for URL.")
//	} else if statusCode == 415 {
//		fmt.Println("Invalid content type.")
//	} else if statusCode == 429 {
//		fmt.Println("The API request rate is higher than 4 request per second or inflight API requests are higher than 10 requests per second.")
//	} else if statusCode == 500 {
//		fmt.Println("Internal server error.")
//	} else if statusCode == 503 {
//		fmt.Println("Service is currently unavailable.")
//	}
//}
