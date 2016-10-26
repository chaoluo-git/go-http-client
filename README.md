# go-client
client facotry

example code
```golang
googleAddress, _ := url.Parse("https://www.google.com")
	token := "eyJhbGciOiJSUzI1NiJ9"

	builder := ((*googleClientFactoryBuilder)(nil)).NewBuilder()
	builder.BaseUrl = googleAddress
	simpleProvider := authorizationProvider.SimpleBearerAuthorizationProvider{Token:token}
	builder.AuthorizationProvider = simpleProvider
	response, err := builder.Builder().newGoogleClient().GetGoogle()
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
  ```
