package list

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"

	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/d1"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List D1 databases",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	d1Client, ok := cmd.Context().Value(constant.D1ClientContextKey).(*d1.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	results, err := d1Client.List(cmd.Context())

	if err != nil {
		return err
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].CreatedAt.Before(results[j].CreatedAt)
	})

	slices.Reverse(results)

	data, err := json.MarshalIndent(results, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}
