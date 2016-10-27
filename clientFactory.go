package client

import (
	"net/url"
	"net/http"
	"go-http-client/authorizationProvider"
	"go-http-client/entityMapper"
)

type CClientFactory struct {
	client *http.Client
	baseUrl *url.URL
	authorizationProvider authorizationProvider.AuthorizationProvider
	entityMapper entityMapper.EntityMapper
}


func(cf *CClientFactory) newCClientFactory(builder *CBuilder)(*CClientFactory) {
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
	EntityMapper entityMapper.EntityMapper
}

func(builder *CBuilder) NewBuilder()(*CBuilder) {
	jsonMapper := entityMapper.JsonMapper{}
	return &CBuilder{EntityMapper: jsonMapper}
}

func(builder *CBuilder) Build()(*CClientFactory) {
	return (*CClientFactory)(nil).newCClientFactory(builder)
}


