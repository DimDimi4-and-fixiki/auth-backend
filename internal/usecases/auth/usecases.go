package auth

type usecases struct {
	numVerifyClient numVerifyClient
}

func New(c Config) usecases {
	return usecases{
		numVerifyClient: c.NumVerifyClient,
	}
}
