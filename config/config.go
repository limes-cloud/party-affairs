package config

type AuthInfo struct {
	AppId     string
	AppSecret string
}

type Config struct {
	DefaultUserAvatar string
	Service           struct {
		Resource   string
		UserCenter string
	}
	Auth struct {
		YiBan AuthInfo
	}
}
