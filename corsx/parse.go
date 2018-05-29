package corsx

import (
	"strconv"

	"github.com/ory/go-convenience/stringsx"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func ParseOptions() cors.Options {
	allowCredentials, _ := strconv.ParseBool(viper.GetString("CORS_ALLOWED_CREDENTIALS"))
	debug, _ := strconv.ParseBool(viper.GetString("CORS_DEBUG"))
	maxAge, _ := strconv.Atoi(viper.GetString("CORS_MAX_AGE"))
	return cors.Options{
		AllowedOrigins:   stringsx.Splitx(viper.GetString("CORS_ALLOWED_ORIGINS"), ","),
		AllowedMethods:   stringsx.Splitx(viper.GetString("CORS_ALLOWED_METHODS"), ","),
		AllowedHeaders:   stringsx.Splitx(viper.GetString("CORS_ALLOWED_HEADERS"), ","),
		ExposedHeaders:   stringsx.Splitx(viper.GetString("CORS_EXPOSED_HEADERS"), ","),
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
		Debug:            debug,
	}
}
