package cli

import (
	"context"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/moutend/d1/internal/cli/binding"
	"github.com/moutend/d1/internal/cli/bulkcreate"
	"github.com/moutend/d1/internal/cli/constant"
	"github.com/moutend/d1/internal/cli/count"
	"github.com/moutend/d1/internal/cli/create"
	"github.com/moutend/d1/internal/cli/delete"
	"github.com/moutend/d1/internal/cli/list"
	"github.com/moutend/d1/internal/cli/verify"
	"github.com/moutend/d1/internal/d1"
	"github.com/moutend/d1/internal/transport"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:               "d1",
	Short:             "d1: Cloudflare wrangler d1 alternative",
	SilenceUsage:      true,
	PersistentPreRunE: rootCommandPersistentPreRunE,
}

func rootCommandPersistentPreRunE(cmd *cobra.Command, args []string) error {
	viper.BindEnv("api_token")

	transport := transport.New()

	if debug, _ := cmd.Flags().GetBool("debug"); debug {
		transport.SetLogger(log.New(cmd.OutOrStdout(), "debug: ", 0))
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	d1Client, err := d1.NewClient(
		httpClient,
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		os.Getenv("CLOUDFLARE_API_TOKEN"),
		os.Getenv("CLOUDFLARE_D1_LOCATION"),
	)

	if err != nil {
		return err
	}

	ctx := cmd.Context()
	ctx = context.WithValue(ctx, constant.D1ClientContextKey, d1Client)

	cmd.SetContext(ctx)

	return nil
}

func init() {
	RootCommand.PersistentFlags().BoolP("debug", "d", false, "Enable debug output")

	if info, ok := debug.ReadBuildInfo(); ok {
		RootCommand.Version = info.Main.Version
	} else {
		RootCommand.Version = "undefined"
	}

	RootCommand.AddCommand(verify.Command)
	RootCommand.AddCommand(count.Command)
	RootCommand.AddCommand(list.Command)
	RootCommand.AddCommand(binding.Command)
	RootCommand.AddCommand(bulkcreate.Command)
	RootCommand.AddCommand(create.Command)
	RootCommand.AddCommand(delete.Command)
}
