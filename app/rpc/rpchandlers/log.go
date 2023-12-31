package rpchandlers

import (
	"github.com/karlsen-network/karlsend/infrastructure/logger"
	"github.com/karlsen-network/karlsend/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
