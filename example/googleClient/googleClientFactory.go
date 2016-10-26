package googleClient

import "go-client/client"

type googleClientFactory struct {
	clientFactory *client.CClientFactory
}

func(gf *googleClientFactory) newGoogleClientFactory(builder *googleClientFactoryBuilder)(*googleClientFactory) {
	parentClientFactory := builder.CBuilder.Builder()
	return &googleClientFactory{parentClientFactory}
}

func(gf *googleClientFactory) newGoogleClient() (*googleClient) {
	return ((*googleClient)(nil)).newClient(gf)
}


type googleClientFactoryBuilder struct {
	*client.CBuilder
}

func(gfBuilder *googleClientFactoryBuilder) NewBuilder() (builder *googleClientFactoryBuilder) {
	parentBuilder := ((*client.CBuilder)(nil)).NewBuilder()
	builder = &googleClientFactoryBuilder{parentBuilder}
	return
}

func(gfBuilder *googleClientFactoryBuilder) Builder() (factory *googleClientFactory) {
	factory = ((*googleClientFactory)(nil)).newGoogleClientFactory(gfBuilder)
	return
}




