package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/LqtCode/kubectl-clog/pkg/plugin"
)

func InitAndExecute() {
	if err := plugin.RunPlugin(context.TODO()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
