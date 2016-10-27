package googleClient

import (
	"testing"
	"net/url"
)

var googleClient *GoogleClient

func init() {
	googleAddress, _ := url.Parse("https://www.google.com")
	builder := ((*googleClientFactoryBuilder)(nil)).NewBuilder()
	builder.BaseUrl = googleAddress
	googleClient = builder.Build().newGoogleClient()
}

func TestGetGoogle(t *testing.T) {
	response, error := googleClient.GetGoogle()
	if error != nil {
		t.Error(error)
		return
	}
	t.Log("response body ", response.Payload)
}