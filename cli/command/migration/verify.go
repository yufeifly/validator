package migration

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
	"github.com/yufeifly/validator/verify"
)

func newVerifyCommand() *cobra.Command {
	var opts verify.CheckOptions
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "verify the result of migration",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return verify.CheckResult(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.Addr, "addr", "", "", "target address, format: 192.168.0.1:6666")
	flags.StringVarP(&opts.Range, "range", "", "", "key range to validate, [start: end) , format: key1:key10000")
	return cmd
}
