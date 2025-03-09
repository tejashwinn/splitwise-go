package types

type Config struct {
	Db     DbConfig
	Server ServerConfig
	Jwt    JwtConfig
}

type JwtConfig struct {
	SecretKey              []byte
	Issuer                 string
	JwtAccessTokenExpMin   int32
	JwtRefreshTokenExpHour int32
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Url string
}
