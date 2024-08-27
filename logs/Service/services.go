package Service

import (
	"logs/logger"
)


var loger = logger.NewLogger()

func NewService(){
	loger.Info().Msg("no info")
}