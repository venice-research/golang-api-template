package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func logRequest(r *http.Request) string {
	url := r.URL
	method := r.Method

	message := fmt.Sprintf("%s %s", method, url)

	fmt.Printf("Got request! %s\n", message)

	return message
}

func generateStartupMessage(currentTime time.Time) string {
	startupMessage := "GoLang API Template started at: " + currentTime.Format("01-02-2006")

	return startupMessage
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-api-template",
	Short: "GoLang API Template",
	Long: `An application template for Go that allows a developer to 
	quickly deploy applications to the DigitalOcean Apps platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			logRequest(r)

			_, err := fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)

			if err != nil {
				fmt.Printf("Error writing response to writer. %s\n", err)
			}
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "80"
		}

		bindAddr := fmt.Sprintf(":%s", port)
		lines := strings.Split(generateStartupMessage(time.Now()), "\n")
		fmt.Println()
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()
		fmt.Printf("==> Server listening at %s 🚀\n", bindAddr)

		server := &http.Server{
			Addr:              fmt.Sprintf(":%s", port),
			ReadHeaderTimeout: 3 * time.Second,
		}

		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golang-api-template.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".golang-api-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".golang-api-template")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}
	}
}
