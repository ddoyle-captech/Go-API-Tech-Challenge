package config

/*
ENV: dev
LOG_LEVEL: DEBUG

DATABASE_CONTAINER_NAME: courses-db-container
DATABASE_NAME: coursesDB
DATABASE_USER: courses-db-user
DATABASE_PASSWORD: courses-db-password
DATABASE_HOST: localhost
DATABASE_PORT: 5432
DATABASE_RETRY_DURATION_SECONDS: 3

HTTP_DOMAIN: localhost
HTTP_PORT: :8080
*/

type Config struct {
	Env      string
	LogLevel string

	DBContainerName string
	DBName          string
	DBUser          string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBRetryDuration string

	HTTPDomain string
	HTTPPort   string
}

func Load(path string) (Config, error) {
	return Config{}, nil
}
