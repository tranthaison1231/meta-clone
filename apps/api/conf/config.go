package conf

type Config struct {
	JwtSecret      string `json:"jwt_secret" env:"JWT_SECRET"`
	TokenExpiresIn int    `json:"token_expires_in" env:"TOKEN_EXPIRES_IN"`
}

var (
	Conf *Config
)

func DefaultConfig() *Config {
	return &Config{
		JwtSecret:      "xxx-yyy-zzz",
		TokenExpiresIn: 48,
	}
}
