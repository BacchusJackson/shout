package main

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "shout",
		Short:        "A CI/CD Logging and Alerts Utility",
		Example:      "",
		SilenceUsage: true,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print version and application metadata",
		Run:   runVersion,
	}

	logCmd := &cobra.Command{
		Use:     "log",
		Short:   "Log an info message",
		Example: "shout log 'some message'",
		Run: func(cmd *cobra.Command, args []string) {
			writeMessage(cmd.OutOrStdout(), "LOG", strings.Join(args, " "))
		},
	}

	warnCmd := &cobra.Command{
		Use:     "warn",
		Short:   "Log a warning message",
		Example: "shout warn 'some warning'",
		Run: func(cmd *cobra.Command, args []string) {
			writeMessage(cmd.OutOrStdout(), "WARN", strings.Join(args, " "))
		},
	}

	errorCmd := &cobra.Command{
		Use:     "error",
		Short:   "Log an error message",
		Example: "shout error 'some error'",
		Run: func(cmd *cobra.Command, args []string) {
			writeMessage(cmd.OutOrStdout(), "ERROR", strings.Join(args, " "))
		},
	}

	cmd.AddCommand(versionCmd, logCmd, warnCmd, errorCmd)

	return cmd
}

func writeMessage(w io.Writer, level string, msg string) {
	c := color.New()
	switch level {
	case "LOG":
		c.Add(color.FgGreen)
	case "WARN":
		c.Add(color.FgYellow)
	case "ERROR":
		c.Add(color.FgRed)

	}
	_, _ = fmt.Fprintf(w, "[%s] %s\n", c.Sprint(level), msg)
}

func runVersion(cmd *cobra.Command, _ []string) {
	_, _ = fmt.Fprintf(cmd.OutOrStdout(),
		`CLIVersion:     %s
GitCommit:      %s
Build Date:     %s
GitDescription: %s
Platform:       %s/%s
GoVersion:      %s
Compiler:       %s
`,
		cliVersion,
		gitCommit,
		buildDate,
		gitDescription,
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
		runtime.Compiler,
	)
}
