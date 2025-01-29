package clients

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Valentin-Foucher/social-stats/common/commonutils"
)

type HttpClient struct {
	inner      *http.Client
	baseUrl    string
	maxRetries int
}

func getRequest(
	url string,
	method string,
	body io.Reader,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	if _, found := headers["Content-Type"]; !found {
		headers["Content-Type"] = "application/json"
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}

	for k, v := range queryParameters {
		q := request.URL.Query()
		for _, e := range v {
			q.Add(k, e)
		}
		request.URL.RawQuery = q.Encode()
	}
	return request, nil
}

func (client *HttpClient) request(
	method string,
	path string,
	statusCodes []int,
	body io.Reader,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", client.baseUrl, path)

	request, err := getRequest(
		url,
		method,
		body,
		headers,
		queryParameters,
	)

	if err != nil {
		return nil, err
	}

	triesCount := 0
	var response *http.Response
	for {
		response, err = client.inner.Do(request)

		if !commonutils.SliceContains(statusCodes, response.StatusCode) {
			body, _ := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("unexpected status code (%d): %s", response.StatusCode, string(body))
		}
		if err == nil {
			break
		}
		if triesCount > client.maxRetries {
			return nil, fmt.Errorf("too many retries, status code (%d): %s", response.StatusCode, err)
		}
		triesCount += 1
	}

	return response, nil
}

func (client *HttpClient) get(
	path string,
	statusCodes []int,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Response, error) {
	return client.request("GET", path, statusCodes, nil, headers, queryParameters)
}

func (client *HttpClient) post(
	path string,
	statusCodes []int,
	body io.Reader,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Response, error) {
	return client.request("POST", path, statusCodes, body, headers, queryParameters)
}

func (client *HttpClient) put(
	path string,
	statusCodes []int,
	body io.Reader,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Response, error) {
	return client.request("PUT", path, statusCodes, body, headers, queryParameters)
}

func (client *HttpClient) delete(
	path string,
	statusCodes []int,
	headers map[string]string,
	queryParameters map[string][]string,
) (*http.Response, error) {
	return client.request("DELETE", path, statusCodes, nil, headers, queryParameters)
}
