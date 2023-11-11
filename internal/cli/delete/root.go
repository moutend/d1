package delete

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "delete",
	Short: "Delete D1 database",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	d1Client, ok := cmd.Context().Value(constant.D1ClientContextKey).(*d1.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	result, err := d1Client.Delete(cmd.Context(), args[0])

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	if !result.Success {
		os.Exit(1)
	}

	return nil
}
