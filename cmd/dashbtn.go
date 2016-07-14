package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mannkind/dashbtn/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const version string = "0.1.0"

var cfgFile string
var port string
var reload = make(chan bool)

// DashBtnCmd - The root command
var DashBtnCmd = &cobra.Command{
	Use:   "dashbtn",
	Short: "Make Amazon Dash buttons do different things using dnsmasq",
	Long:  "Make Amazon Dash buttons do different things using dnsmasq",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			http.HandleFunc("/", handlers.IndexHandler)
			http.HandleFunc("/dash", handlers.DashHandler)

			if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
				log.Fatal("ListenAndServe: ", err)
			}

			select {
			case <-reload:
				continue
			}
		}
	},
}

// Execute - Adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := DashBtnCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.Printf("DashBtn Version: %s", version)

	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Configuration Changed: %s", e.Name)
			reload <- true
		})

		log.Printf("Loading Configuration %s", cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error Loading Configuration: %s ", err)
		}
		log.Printf("Loaded Configuration %s", cfgFile)
	})

	DashBtnCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".dashbtn.yaml", "The path to the configuration file")
	DashBtnCmd.PersistentFlags().StringVarP(&port, "port", "p", "8001", "The port to run the http server on")
}
