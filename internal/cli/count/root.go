package count

import (
	"fmt"

	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "count",
	Short: "Count D1 databases",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	d1Client, ok := cmd.Context().Value(constant.D1ClientContextKey).(*d1.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	count, err := d1Client.Count(cmd.Context())

	if err != nil {
		return err
	}

	cmd.Println(count)

	return nil
}
