package main

import (
	"context"

	"github.com/jsclayton/arr-utils/cmd/arr/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
