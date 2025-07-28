package config

type AppConfig struct {
    RedisHost string
    RedisPort int
    JWTSecret string
    // ...other app-specific configs
}
