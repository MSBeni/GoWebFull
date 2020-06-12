package main

import(
	"github.com/spf13/cobra"
	"time"
)

func printTimeCmd() *cobra.Command{
	return &cobra.Command{
		Use: "current time",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey Gophrs, the current time is", prettyTime)
			return nil
		},
	}
}