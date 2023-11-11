package binding

import (
	"encoding/json"
	"os"

	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "binding",
	Short: "Generate D1 bindings for wrangler.toml",
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

	var results []d1.ListResult

	if err := json.Unmarshal(input, &results); err != nil {
		return err
	}
	for i, result := range results {
		cmd.Println("[[d1_databases]]")
		cmd.Printf("binding = %q\n", result.Name)
		cmd.Printf("database_name = %q\n", result.Name)
		cmd.Printf("database_id = %q\n", result.Uuid)

		if i < len(results)-1 {
			cmd.Printf("\n")
		}
	}

	return nil
}
