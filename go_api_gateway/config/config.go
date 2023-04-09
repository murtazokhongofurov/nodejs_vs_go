package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string 
	ServiceName string 
	HttpPort string 
	DriverServiceHost string
	DriverServicePort string
}

func Load() Config {
	if err := godotenv.Load("/config/.env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "go_api_gateway"))
	c.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))
	c.DriverServiceHost = cast.ToString(getOrReturnDefaultValue("DRIVER_SERVICE_HOST", "localhost"))
	c.DriverServicePort = cast.ToString(getOrReturnDefaultValue("DRIVER_SERVICE_PORT", "2001"))
	c.HttpPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8181"))

	return c

}


func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
