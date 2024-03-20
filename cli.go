package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var errFailure = errors.New("command failure")

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

	alertCmd := &cobra.Command{
		Use:   "alert [WEBHOOK URL] key1=value1 key2=value2",
		Short: "send an alert to a JSON webhook",
		RunE:  runAlert,
	}

	alertCmd.Flags().StringP("seperator", "s", "=", "A custom seperator to use to split arguments")

	cmd.AddCommand(versionCmd, logCmd, warnCmd, errorCmd, alertCmd)

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

func sendAlert(webhookURL string, content []byte) error {
	var bodyReader io.Reader

	slog.Info("send request to webhook", "url", webhookURL, "body", string(content))
	if len(content) != 0 {
		bodyReader = bytes.NewReader(content)
	}

	res, err := http.Post(webhookURL, "application/json", bodyReader)
	if err != nil {
		slog.Error("alert failed", "response_status", res.Status)
		return err
	}
	return nil
}

func runAlert(cmd *cobra.Command, args []string) error {
	sep, _ := cmd.Flags().GetString("seperator")

	if len(args) == 0 {
		slog.Error("no webhook url specified.")
		return errFailure
	}

	webhookURL := args[0]
	if len(args) == 1 {
		return sendAlert(webhookURL, nil)
	}

	argMap := make(map[string]string)

	for _, a := range args[1:] {
		parts := strings.Split(a, sep)
		if len(parts) != 2 {
			slog.Error("invalid format", "argument", a)
			return errFailure
		}

		argMap[parts[0]] = parts[1]
	}

	jsonContent, err := json.Marshal(argMap)
	if err != nil {
		slog.Error("cannot create JSON from arguments", "error", err)
		return err
	}

	return sendAlert(webhookURL, jsonContent)
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
