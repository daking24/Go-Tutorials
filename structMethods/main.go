package main


type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	authenticate := "Authorization: Basic " + a.username+ ":" + a.password
	return authenticate
}