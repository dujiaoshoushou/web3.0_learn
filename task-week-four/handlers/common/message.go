package common

import "github.com/spf13/viper"

var JwtSecret = []byte(viper.GetString("jwt.jwt_secret"))
