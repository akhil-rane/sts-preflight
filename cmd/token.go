package cmd

import (
	"github.com/sjenning/sts-preflight/pkg/cmd/token"
	"github.com/sjenning/sts-preflight/pkg/jwt"
	"github.com/spf13/cobra"
)

var tokenConfig token.Config

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Creates a token signed by the RSA private key and validated by the OIDC provider",
	Run: func(cmd *cobra.Command, args []string) {
		jwt.New(tokenConfig)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.PersistentFlags().Int64Var(&tokenConfig.ExpireSeconds, "expire-seconds", 3600, "Token expiration duration in seconds")
}
