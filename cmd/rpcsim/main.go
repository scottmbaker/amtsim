package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"amtsim/internal/logging"
	"amtsim/internal/rpc"
)

// TODO: Improve parameter handling. Use cobra/viper or similar.

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rpcsim <command> [flags]")
		fmt.Println("Commands: activate, amtinfo, deactivate")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "activate":
		activateCmd := flag.NewFlagSet("activate", flag.ExitOnError)
		urlFlag := activateCmd.String("u", "", "RPS URL (wss://...)")
		activateCmd.StringVar(urlFlag, "url", "", "RPS URL (wss://...)")
		noVerifyFlag := activateCmd.Bool("n", false, "Skip cert verify")
		profileFlag := activateCmd.String("profile", "", "Profile Name")
		debugFlag := activateCmd.Bool("debug", false, "Enable debug logging")
		actPasswordFlag := activateCmd.String("password", "", "AMT Password")

		if err := activateCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if *debugFlag {
			logging.SetLevel(logging.LEVEL_DEBUG)
		}

		if *urlFlag == "" {
			fmt.Println("-u is required for activation")
			os.Exit(1)
		}

		password := *actPasswordFlag
		if password == "" {
			password = os.Getenv("AMT_PASSWORD")
		}
		if password == "" {
			password = readPassword()
		}

		profile := *profileFlag
		if profile == "" {
			profile = os.Getenv("PROFILE")
		}

		logging.Info("Activating against %s with profile %s", *urlFlag, *profileFlag)
		client := rpc.NewClient(*urlFlag, profile, password, *noVerifyFlag)
		if err := client.Execute("activate"); err != nil {
			logging.Fatal("Activation failed: %v", err)
		}
		logging.Info("Activation Sequence Completed Successfully.")

	case "amtinfo":
		client := rpc.NewClient("", "", "", false)
		if err := client.AmtInfo(); err != nil {
			logging.Fatal("AmtInfo failed: %v", err)
		}

	case "deactivate":
		deactivateCmd := flag.NewFlagSet("deactivate", flag.ExitOnError)
		deacUrlFlag := deactivateCmd.String("u", "", "RPS URL (wss://...)")
		deactivateCmd.StringVar(deacUrlFlag, "url", "", "RPS URL (wss://...)")
		deacLocalFlag := deactivateCmd.Bool("local", false, "Deactivate locally via WSMAN")
		deactivateCmd.BoolVar(deacLocalFlag, "l", false, "Deactivate locally via WSMAN")
		deacNoVerifyFlag := deactivateCmd.Bool("n", false, "Skip cert verify")
		deacDebugFlag := deactivateCmd.Bool("debug", false, "Enable debug logging")
		deacPasswordFlag := deactivateCmd.String("password", "", "AMT Password")

		if err := deactivateCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *deacDebugFlag {
			logging.SetLevel(logging.LEVEL_DEBUG)
		}

		if *deacLocalFlag && *deacUrlFlag != "" {
			fmt.Println("Error: Cannot specify both --local and -u")
			os.Exit(1)
		}

		if !*deacLocalFlag && *deacUrlFlag == "" {
			fmt.Println("Error: Must specify either --local or -u <url>")
			deactivateCmd.Usage()
			os.Exit(1)
		}

		password := *deacPasswordFlag
		if password == "" {
			password = os.Getenv("AMT_PASSWORD")
		}
		if password == "" {
			password = readPassword()
		}

		if *deacLocalFlag {
			client := rpc.NewClient("http://localhost:16992/wsman", "", password, *deacNoVerifyFlag)
			if err := client.Deactivate(); err != nil {
				logging.Fatal("Local Deactivation failed: %v", err)
			}
		} else {
			cmdStr := "deactivate --password " + password
			client := rpc.NewClient(*deacUrlFlag, "", password, *deacNoVerifyFlag)
			if err := client.Execute(cmdStr); err != nil {
				logging.Fatal("Remote Deactivation failed: %v", err)
			}
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func readPassword() string {
	fmt.Print("Enter AMT Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		// Fallback for non-interactive
		reader := bufio.NewReader(os.Stdin)
		pass, _ := reader.ReadString('\n')
		return strings.TrimSpace(pass)
	}
	fmt.Println()
	return string(bytePassword)
}
