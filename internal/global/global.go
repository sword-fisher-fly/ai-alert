package global

import (
	"github.com/sword-fisher-fly/ai-alert/config"
	"github.com/spf13/viper"
)

var (
	Layout    = "2006-01-02 15:04:05"
	Config    config.App
	Version   string
	StSignKey = []byte(viper.GetString("jwt.ai"))
)
