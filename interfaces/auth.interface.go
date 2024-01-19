package interfaces

type AuthDataLayer interface {
	VerifyCredentials(username string, password string) (string, error)
}

type AuthServiceLayer interface {
	Login(username string, password string) (string, error)
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (string, error)
}
