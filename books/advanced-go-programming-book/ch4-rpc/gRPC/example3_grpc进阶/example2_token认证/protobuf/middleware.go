package proto

import "golang.org/x/net/context"

type Authentication struct {
	Login string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string)(map[string]string, error, ){
	return map[string]string{"login":a.Login, "password":a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
