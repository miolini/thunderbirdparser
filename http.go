package thunderbirdparser

import "io/ioutil"
import "net/http"

func httpGet(pageUrl string, httpClient *http.Client) (data []byte, err error) {
	var response *http.Response
	response, err = httpClient.Get(pageUrl)
	if err == nil {
		defer response.Body.Close()
		data, err = ioutil.ReadAll(response.Body)
	}
	return
}