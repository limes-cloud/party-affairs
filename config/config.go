package config

type AuthInfo struct {
	AppId     string
	AppSecret string
}

type Config struct {
	DefaultUserAvatar string
	Service           struct {
		Resource string
	}
	Auth struct {
		YiBan AuthInfo
	}
}
