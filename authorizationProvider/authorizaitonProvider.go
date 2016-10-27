package authorizationProvider

type AuthorizationProvider interface {
	GetAuthorization() string
	RefreshToken() bool
}

/**************************************
|   SimpleBearerAuthorizationProvider  |
***************************************/
type SimpleBearerAuthorizationProvider struct {
	Token string
}

func (s SimpleBearerAuthorizationProvider) GetAuthorization() string {
	return "Bearer " + s.Token
}

func (s SimpleBearerAuthorizationProvider) RefreshToken() bool  {
	return true
}

/***************************************
|      SimpleAuthorizationProvider     |
***************************************/
type SimpleAuthorizationProvider struct {
	Authorization string
}

func (s SimpleAuthorizationProvider) GetAuthorization() string {
	return s.Authorization
}

func (s SimpleAuthorizationProvider) RefreshToken() bool  {
	return true
}

