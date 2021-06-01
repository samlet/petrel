package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// maxNetworkRetriesDelay and minNetworkRetriesDelay defines sleep time in milliseconds between
// tries to send HTTP request again after network failure.
const maxNetworkRetriesDelay = 5000 * time.Millisecond
const minNetworkRetriesDelay = 500 * time.Millisecond

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func InvokeService(srv string, reqstr string) (string, error) {
	url := fmt.Sprintf("https://localhost:8443/rest/services/%s", srv)

	payload := strings.NewReader(reqstr)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Alfin-Version", APIVersion)
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJ1c2VyTG9naW5JZCI6ImFkbWluIiwiaXNzIjoiQXBhY2hlT0ZCaXoiLCJleHAiOjE2MzQ1OTc3NTcsImlhdCI6MTYxNjU5Nzc1N30.Luuf2bK7ZJ8KE_CtsA3iPZ189i-Qbm2qK5r5VfeQcJqIyTKy4DHf2fBAp37W8OtU6SIplwCdnbTMtHuCZ5h8cA")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	return string(body), nil
}

func Post(srv string, objmap map[string]interface{}) (string, error) {
	out, err := json.Marshal(objmap)
	if err != nil {
		return "", err
	}
	return InvokeService(srv, string(out))
}

type AlfinParams struct {
	Params `json:"-"`
}

//func (c *AlfinParams) GetParams() *AlfinParams {
//	//return &Params{c.Metadata, c.Payload}
//	return c
//}

//type PA map[string]string
//func NewParams(input interface{}) (*AlfinParams, error) {
//	payload, err:=json.Marshal(input)
//	if err != nil {
//		return nil, err
//	}
//	return &AlfinParams{Params{PA{}, string(payload)}}, nil
//}

type AlfinBackend struct {
	Logger              *zap.Logger
	networkRetriesSleep bool
	httpClient          *http.Client
}

func NewAlfinBackend() *AlfinBackend {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &AlfinBackend{Logger: logger,
		networkRetriesSleep: false,
		httpClient:          http.DefaultClient,
	}
}

// UnmarshalJSONVerbose unmarshals JSON, but in case of a failure logs and
// produces a more descriptive error.
//func (s *AlfinBackend) UnmarshalJSONVerbose(statusCode int, body []byte, v interface{}) error {
//	err := json.Unmarshal(body, v)
//	if err != nil {
//		// If we got invalid JSON back then something totally unexpected is
//		// happening (caused by a bug on the server side). Put a sample of the
//		// response body into the error message so we can get a better feel for
//		// what the problem was.
//		bodySample := string(body)
//		if len(bodySample) > 500 {
//			bodySample = bodySample[0:500] + " ..."
//		}
//
//		// Make sure a multi-line response ends up all on one line
//		bodySample = strings.Replace(bodySample, "\n", "\\n", -1)
//
//		newErr := fmt.Errorf("Couldn't deserialize JSON (response status: %v, body sample: '%s'): %v",
//			statusCode, bodySample, err)
//		s.Logger.Error(newErr.Error(), zap.Int("status", statusCode))
//		return newErr
//	}
//
//	return nil
//}

// sleepTime calculates sleeping/delay time in milliseconds between failure and a new one request.
func (s *AlfinBackend) sleepTime(numRetries int) time.Duration {
	// We disable sleeping in some cases for tests.
	if !s.networkRetriesSleep {
		return 0 * time.Second
	}

	// Apply exponential backoff with minNetworkRetriesDelay on the
	// number of num_retries so far as inputs.
	delay := minNetworkRetriesDelay + minNetworkRetriesDelay*time.Duration(numRetries*numRetries)

	// Do not allow the number to exceed maxNetworkRetriesDelay.
	if delay > maxNetworkRetriesDelay {
		delay = maxNetworkRetriesDelay
	}

	// Apply some jitter by randomizing the value in the range of 75%-100%.
	jitter := rand.Int63n(int64(delay / 4))
	delay -= time.Duration(jitter)

	// But never sleep less than the base sleep seconds.
	if delay < minNetworkRetriesDelay {
		delay = minNetworkRetriesDelay
	}

	return delay
}

func (c *AlfinBackend) Call(method, path, key string, params ParamsContainer, v LastResponseSetter) error {
	srv := path
	url := fmt.Sprintf("https://localhost:8443/rest/services/%s", srv)

	//payload := strings.NewReader(params.GetParams().Payload)
	payload, err := json.Marshal(params)
	c.Logger.Debug("payload: " + string(payload))
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(method, url, bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Alfin-Version", APIVersion)
	if key == "" {
		key = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJ1c2VyTG9naW5JZCI6ImFkbWluIiwiaXNzIjoiQXBhY2hlT0ZCaXoiLCJleHAiOjE2MzQ1OTc3NTcsImlhdCI6MTYxNjU5Nzc1N30.Luuf2bK7ZJ8KE_CtsA3iPZ189i-Qbm2qK5r5VfeQcJqIyTKy4DHf2fBAp37W8OtU6SIplwCdnbTMtHuCZ5h8cA"
	}
	req.Header.Add("Authorization", "Bearer "+key)

	// add all metadata
	for k, v := range params.GetParams().Metadata {
		req.Header.Add(k, v)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	c.Logger.Debug("Response", zap.String("body", string(resBody)))
	//err = c.UnmarshalJSONVerbose(res.StatusCode, resBody, v)
	v.SetLastResponse(newAPIResponse(res, resBody))

	typedResult, err := WrapTypedResult(string(resBody), &v)
	if typedResult != nil {
		v.SetLastTypedResponse(typedResult)
	}

	return err
}
