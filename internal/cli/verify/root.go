package verify

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "verify",
	Short: "Verify API token",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	client, ok := cmd.Context().Value(constant.D1ClientContextKey).(*d1.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	result, err := client.Verify(cmd.Context())

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}
