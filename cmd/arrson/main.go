package main

import (
	"context"

	"github.com/jsclayton/arrson/cmd/arrson/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
