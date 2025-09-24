package rum

import (
	"log"
	"time"
)

type HandlerChainExecutor struct {
	ctx       *RumContext
	chain     HandlerChain
	startTime time.Time
	completed bool
}

func NewHandlerChainExecutor(ctx *RumContext, chain HandlerChain) *HandlerChainExecutor {
	return &HandlerChainExecutor{
		ctx:   ctx,
		chain: chain,
	}
}

func (exec *HandlerChainExecutor) Begin() {
	exec.startTime = time.Now()
	if len(exec.chain) > 0 {
		handler := exec.chain[0]
		handler(exec.ctx)
	}
}

func (exec *HandlerChainExecutor) Complete() {
	if !exec.completed {
		latency := time.Since(exec.startTime)
		statusCode := exec.ctx.statusCode
		log.Printf("[RUM] Request completed: status=%d, latency=%v", statusCode, latency)
		exec.completed = true
	}
}
