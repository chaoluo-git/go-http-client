package client

import (
	"net/url"
	"net/http"
	"go-client/client/authorizationProvider"
)

type CClientFactory struct {
	client *http.Client
	baseUrl *url.URL
	authorizationProvider authorizationProvider.AuthorizationProvider
	entityMapper EntityMapper
}


func(f *CClientFactory) newCClientFactory(builder *CBuilder)(*CClientFactory) {
	clientFactory := &CClientFactory{}
	clientFactory.client = &http.Client{}
	clientFactory.baseUrl = builder.BaseUrl
	clientFactory.authorizationProvider = builder.AuthorizationProvider
	clientFactory.entityMapper = builder.EntityMapper
	return clientFactory
}


type CBuilder struct {
	BaseUrl *url.URL
	AuthorizationProvider authorizationProvider.AuthorizationProvider
	EntityMapper EntityMapper
}

func(builder *CBuilder) NewBuilder()(*CBuilder) {
	jsonMapper := JsonMapper{}
	return &CBuilder{EntityMapper: jsonMapper}
}

func(builder *CBuilder) Builder()(*CClientFactory) {
	return (*CClientFactory)(nil).newCClientFactory(builder)
}


