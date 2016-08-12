package workers

import (
	"github.com/Focinfi/sakura/config"
	"gopkg.in/go-playground/pool.v3"
)

var defaultQueue = pool.NewLimited(config.Config.DeaultWorkerNum)
