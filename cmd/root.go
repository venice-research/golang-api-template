package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var cfgFile string

func logRequest(r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	fmt.Println("Got request!", method, uri)
}

func generateStartupMessage(time time.Time) string {
	startupMessage := "GoLang API Template started at: " + time.Format("01-02-2006")

	fmt.Println(rand.Int())

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
			fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)
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

		err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
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
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
