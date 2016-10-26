package client

import (
	"net/http"
	"net/url"
	"fmt"
	"go-client/client/authorizationProvider"
)

type CClient struct {
	client *http.Client
	url    *url.URL
	authorizationProvider authorizationProvider.AuthorizationProvider
	entityMapper EntityMapper
}


func(cli *CClient) NewClient(clientFactory *CClientFactory)(*CClient) {
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

func (cli *CClient) execute(request *http.Request)(*http.Response, error) {
	fmt.Println("url", request.URL.String())

	if cli.authorizationProvider != nil {
		request.Header.Set("Authorization", cli.authorizationProvider.GetAuthorization())
	}
	return cli.client.Do(request)
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





