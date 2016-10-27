package googleClient

import "go-http-client"

type googleClientFactory struct {
	clientFactory *client.CClientFactory
}

func(gf *googleClientFactory) newGoogleClientFactory(builder *googleClientFactoryBuilder)(*googleClientFactory) {
	parentClientFactory := builder.CBuilder.Build()
	return &googleClientFactory{parentClientFactory}
}

func(gf *googleClientFactory) newGoogleClient() (*GoogleClient) {
	return ((*GoogleClient)(nil)).newClient(gf)
}


type googleClientFactoryBuilder struct {
	*client.CBuilder
}

func(gfBuilder *googleClientFactoryBuilder) NewBuilder() (builder *googleClientFactoryBuilder) {
	parentBuilder := ((*client.CBuilder)(nil)).NewBuilder()
	builder = &googleClientFactoryBuilder{parentBuilder}
	return
}

func(gfBuilder *googleClientFactoryBuilder) Build() (factory *googleClientFactory) {
	factory = ((*googleClientFactory)(nil)).newGoogleClientFactory(gfBuilder)
	return
}




