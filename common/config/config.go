package config

import "os"

type Config struct {
	APIAddr        string
	FEAddr         string
	MongoDBURI     string
	JWTSecret      string
	RedisAddr      string
	RedisPass      string
	GmailEmail     string
	GmailEmailPass string
	PayOS          string
	PayOSClientID  string
	PayOSApiKey    string
	PayOSChecksum  string
	CloudinaryURL  string
}

func LoadConfig() *Config {
	cfg := &Config{
		APIAddr:        getEnv("API_ADDR"),
		FEAddr:         getEnv("FE_ADDR"),
		MongoDBURI:     getEnv("MONGODB_URI"),
		JWTSecret:      getEnv("JWT_SECRET"),
		RedisAddr:      getEnv("REDIS_ADDR"),
		RedisPass:      getEnv("REDIS_PASSWORD"),
		GmailEmail:     getEnv("GMAIL_EMAIL"),
		GmailEmailPass: getEnv("GMAIL_EMAIL_PASSWORD"),
		PayOS:          getEnv("PAYOS"),
		PayOSClientID:  getEnv("PAYOS_CLIENT_ID"),
		PayOSApiKey:    getEnv("PAYOS_API_KEY"),
		PayOSChecksum:  getEnv("PAYOS_CHECKSUM"),
		CloudinaryURL:  getEnv("CLOUDINARY_URL"),
	}
	return cfg
}

func getEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return ""
}
