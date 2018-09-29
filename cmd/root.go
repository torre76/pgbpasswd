package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/torre76/pgbpasswd/encrypt"
)

// Target file where to append built password
var targetFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pgbpasswd <login> <password>",
	Short: "Create MD5 encrypted password to be used with PgBouncer",
	Long: `pgbpasswd is used to create PostgreSQL compliant MD5 encrypted password
that will be used with PgBouncer to establish a connection to the source 
database server.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var encPassword string
		if len(args) == 2 {
			encPassword = encrypt.PgMd5HashedPassword(args[0], args[1])
		} else {
			encPassword = ""
		}
		var boldGreenColor = color.New(color.FgGreen, color.Bold).SprintfFunc()
		fmt.Println("MD5 Hash Password for PgBouncer is \"" + boldGreenColor(encPassword) + "\".\n")

		if targetFile == "" {
			fmt.Println("The line to append to PgBouncer user file is:")
			fmt.Println(boldGreenColor(fmt.Sprintf(
				"\"%s\" \"%s\"\n",
				strings.Replace(args[0], "\"", "\"\"", -1),
				strings.Replace(encPassword, "\"", "\"\"", -1),
			)))
		} else {
			if err := writeToFile(targetFile, args[0], encPassword); err != nil {
				fmt.Println("Error: " + err.Error() + ".")
				fmt.Println("")
			} else {
				fmt.Println("Login and encoded password has been written to \"" + boldGreenColor(targetFile) + "\".")
				fmt.Println("This file can be used as PgBouncer user file.")
				fmt.Println("")
			}

		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetFile, "file", "f", "", "PgBouncer user file where to append built hash")
}

// writeToFile append the PostgreSQL login and hashed password to a file that can be used
// as authorization file for PgBouncer
func writeToFile(targetFile string, login string, hashedPassword string) error {
	var boldRedColor = color.New(color.FgRed, color.Bold).SprintfFunc()
	// Check if this file exists
	var file *os.File
	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
		var fileErr error
		if file, fileErr = os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0644); fileErr != nil {
			return fmt.Errorf("unable to create file \"%s\"", boldRedColor(targetFile))
		}
	} else {
		var fileErr error
		if file, fileErr = os.OpenFile(targetFile, os.O_WRONLY|os.O_APPEND, 0644); fileErr != nil {
			return fmt.Errorf("unable to open file \"%s\" for writing", boldRedColor(targetFile))
		}
	}

	// Here file has been opened
	defer file.Close()

	_, err := file.WriteString(fmt.Sprintf(
		"\"%s\" \"%s\"\n",
		strings.Replace(login, "\"", "\"\"", -1),
		strings.Replace(hashedPassword, "\"", "\"\"", -1),
	))

	if err != nil {
		return fmt.Errorf("unable to write to \"%s\"", boldRedColor(targetFile))
	}

	return nil
}
