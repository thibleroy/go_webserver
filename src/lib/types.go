package lib

type IEnvironment struct {
	WebServerPort int
	MongoURL string
	MongoPort int
	JwtSecret string
}

type IError struct {
	Error   error
	Message string
	Code    int
}
