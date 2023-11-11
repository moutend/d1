package bulkcreate

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "bulkcreate",
	Short: "Create D1 databases",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	input, err := os.ReadFile(args[0])

	if err != nil {
		return err
	}

	var requests []d1.CreateRequest

	if err := json.Unmarshal(input, &requests); err != nil {
		return err
	}

	d1Client, ok := cmd.Context().Value(constant.D1ClientContextKey).(*d1.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	result, err := d1Client.BulkCreate(cmd.Context(), requests)

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
