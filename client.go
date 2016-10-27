package client

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"go-http-client/authorizationProvider"
	"go-http-client/entityMapper"
)

type CClient struct {
	client *http.Client
	url    *url.URL
	authorizationProvider authorizationProvider.AuthorizationProvider
	entityMapper entityMapper.EntityMapper
}


func(cli *CClient) newClient(clientFactory *CClientFactory)(*CClient) {
	client := &CClient{}
	client.url = clientFactory.baseUrl
	client.client = clientFactory.client
	client.authorizationProvider = clientFactory.authorizationProvider
	client.entityMapper = clientFactory.entityMapper
	return client
}

func(cli *CClient) setAuthorization(authorization string) {
	cli.authorizationProvider = authorizationProvider.SimpleAuthorizationProvider{Authorization:authorization}
}

func(cli *CClient) newClientWithBaseUrl(clientFactory CClientFactory, baseUrl *url.URL)(*CClient){
	client := &CClient{}
	client.url = baseUrl
	client.client = clientFactory.client
	return client
}

func (cli *CClient) Get(path ...string)(*CRequest) {
	return cli.newGetRequest(path...)

}

func (cli *CClient) Post(path ...string)(*CRequest) {
	return cli.newPostRequest(path...)
}

func (cli *CClient) Delete(path ...string)(*CRequest) {
	return cli.newDeleteRequest(path...)
}

func (cli *CClient) Put(path ...string)(*CRequest) {
	return cli.newPutRequest(path...)
}

func (cli *CClient) execute(request *http.Request)(*CResponse, error) {
	cResponse := &CResponse{}
	if cli.authorizationProvider != nil {
		request.Header.Set("Authorization", cli.authorizationProvider.GetAuthorization())
	}
	fmt.Println("request url >>>>>", request.URL.String())
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(request.Body)
	//s := buf.String()
	//fmt.Println("request body >>>>>", s)
	fmt.Println("request header >>>>>", request.Header)

	response, error := cli.client.Do(request)
	cResponse.Response = response


	if error != nil {
		fmt.Println("response error <<<<<", error.Error())
		return cResponse, error
	}
	fmt.Println("response status code <<<<<", cResponse.StatusCode)
	body, error := ioutil.ReadAll(cResponse.Body)
	if error != nil {
		fmt.Println("response error <<<<<", error.Error())
		return cResponse, error
	}
	cResponse.Payload = string(body)
	fmt.Println("response header <<<<<", response.Header)
	fmt.Println("response body <<<<<", cResponse.Payload)
	return cResponse, nil
}

func (cli *CClient) newGetRequest(path ...string)(*CRequest) {
	//fmt.Println("URL", this.url)
	getRequest, _ := http.NewRequest("GET", "", nil)
	request := &CRequest{client: cli, request: getRequest, uriBuilder: cli.url}
	return request.Path(path...)
}

func (cli *CClient) newPostRequest(path ...string)(*CRequest) {
	//fmt.Println("URL", this.url)
	postRequest, _ := http.NewRequest("POST", "", nil)
	request := &CRequest{client: cli, request: postRequest, uriBuilder: cli.url}
	return request.Path(path...)
}

func (cli *CClient) newDeleteRequest(path ...string)(*CRequest) {
	//fmt.Println("URL", this.url)
	deleteRequest, _ := http.NewRequest("DELETE", "", nil)
	request := &CRequest{client: cli, request: deleteRequest, uriBuilder: cli.url}
	return request.Path(path...)
}

func (cli *CClient) newPutRequest(path ...string)(*CRequest) {
	//fmt.Println("URL", this.url)
	putRequest, _ := http.NewRequest("PUT", "", nil)
	request := &CRequest{client: cli, request: putRequest, uriBuilder: cli.url}
	return request.Path(path...)
}





