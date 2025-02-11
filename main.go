package main

import (
	"context"
	"fmt"

	"github.com/kgateway-dev/kgateway/v2/pkg/utils/probes"
)

func main() {
	fmt.Println("Testing kgateway import")
	ctx := context.Background()
	probes.StartLivenessProbeServer(ctx)
	<-ctx.Done()
}
