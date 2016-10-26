package googleClient

import (
	"net/http"
	"go-client/client"
)


type googleClient struct {
	*client.CClient
}

func(g *googleClient) newClient(factory *googleClientFactory)(*googleClient){
	parentClient := (*client.CClient)(nil).NewClient(factory.clientFactory)
	return &googleClient{parentClient}
}


func(g *googleClient) GetGoogle()(*http.Response, error) {
	return g.Get().Execute()
}