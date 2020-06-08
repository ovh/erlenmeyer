package cmd

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ovh/erlenmeyer/middlewares"
	"github.com/ovh/erlenmeyer/proto/graphite"
	"github.com/ovh/erlenmeyer/proto/influxdb"
	"github.com/ovh/erlenmeyer/proto/opentsdb"
	"github.com/ovh/erlenmeyer/proto/prom"
	remoteRead "github.com/ovh/erlenmeyer/proto/prom_remote"
	"github.com/ovh/erlenmeyer/proto/warp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/willf/pad"
)

func init() {
	cobra.OnInitialize(initConfig)

	// Application flags
	// Config (--config) flag expect the configuration file name
	RootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.erlenmeyer.yaml)")

	// Verbose (--verbose) flag to activate verbose mode
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// Listen (--listen) flag to so set the proxy listen URL
	RootCmd.Flags().StringP("listen", "l", "127.0.0.1:8080", "listen address")

	// Bind persistent / local flags from cobra to viper
	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	if err := viper.BindPFlags(RootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Bind environment variables
	viper.SetEnvPrefix("erlenmeyer")
	viper.AutomaticEnv()

	// Set config search path
	viper.AddConfigPath("/etc/erlenmeyer/")
	viper.AddConfigPath("$HOME/.erlenmeyer")
	viper.AddConfigPath(".")

	// Load config
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	viper.SetDefault("timeunit", "us")
	viper.SetDefault("prometheus.fillprevious.period", "5 m")

	// Load user defined config
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "erlenmeyer",
	Short: "A proxy that translates queries to Warp10",
	Run: func(cmd *cobra.Command, args []string) {
		// Use echo router
		r := echo.New()
		addr := viper.GetString("listen")

		// Disable echo logger
		r.Logger.SetOutput(ioutil.Discard)

		// Rewrite path
		r.Pre(middleware.Rewrite(map[string]string{
			"/warp10/*": "/warp/$1",
		}))

		// Enable echo middlewares
		r.Use(middleware.MethodOverride())
		r.Use(middleware.Secure())
		r.Use(middleware.Recover())
		r.Use(middlewares.Gzip())

		// Enable custom middlewares
		r.Use(middlewares.CORS())
		r.Use(middlewares.Logger())

		// Expose metrics on /metrics using prometheus
		r.Any("/metrics", echo.WrapHandler(promhttp.Handler()))
		r.Any("/", func(ctx echo.Context) error {
			return ctx.NoContent(http.StatusOK)
		})

		// tokens to deny
		tokens := viper.GetStringSlice("deny.tokens")

		// Register opentsdb handlers
		openTSDB := opentsdb.NewOpenTSDB()
		gOpenTSDB := r.Group("/opentsdb", middlewares.Protocol("opentsdb"), middlewares.Deny(tokens))
		gOpenTSDB.Any("/api/query*", middlewares.Native(openTSDB.HandleQuery))
		gOpenTSDB.Any("/api/query/last*", middlewares.Native(openTSDB.HandleQueryLast))
		gOpenTSDB.Any("/api/suggest*", middlewares.Native(openTSDB.HandleSuggest))
		gOpenTSDB.Any("/api/aggregators*", middlewares.Native(openTSDB.HandleAggregators))
		gOpenTSDB.Any("/api/search/lookup*", middlewares.Native(openTSDB.HandleLookup))
		gOpenTSDB.Any("/api/config/filters*", middlewares.Native(openTSDB.HandleConfigFilters))

		// Register prometheus query language
		promQL := prom.NewPromQL()
		gPromQL := r.Group("/prometheus", middlewares.Protocol("prometheus"), middlewares.Deny(tokens))
		gPromQL.Any("/api/v1/query_range*", middlewares.Native(promQL.QueryRange))
		gPromQL.Any("/api/v1/query*", middlewares.Native(promQL.InstantQuery))
		gPromQL.Any("/api/v1/series*", middlewares.Native(promQL.FindAndDeleteSeries))
		gPromQL.Any("/api/v1/label/:label/values*", promQL.FindLabelsValues)
		gPromQL.Any("/remote_read*", remoteRead.HandlerBuilder())

		// Register graphite query language
		gGraphite := r.Group("/graphite", middlewares.Protocol("graphite"), middlewares.Deny(tokens))
		gGraphite.Any("/render*", middlewares.Native(graphite.Render))
		gGraphite.Any("/metrics*", middlewares.Native(graphite.Find))
		gGraphite.Any("/metrics/find*", middlewares.Native(graphite.Find))
		gGraphite.Any("/metrics/expand*", middlewares.Native(graphite.Expand))
		gGraphite.Any("/metrics/index.json*", middlewares.Native(graphite.Index))

		// Register influx query language
		i := influxdb.NewInfluxDB()
		gInfluxDB := r.Group("/influxdb", middlewares.Protocol("influxdb"), middlewares.Deny(tokens))
		gInfluxDB.Any("/query*", middlewares.Native(i.Query))

		// Register warp handler
		gWarp := r.Group("/warp", middlewares.Protocol("warp10"))
		gWarp.Any("/api/v0/exec", warp.Exec)
		gWarp.Any("/api/v0/*", middlewares.ReverseWithConfig(middlewares.ReverseConfig{
			URL: viper.GetString("warp_endpoint") + "/api/v0",
		}))

		// Display routes mapping
		for _, route := range r.Routes() {
			log.Debugf("%s %s ---> %s", pad.Right(route.Method, 7, " "), pad.Right(route.Path, 50, " "), route.Name)
		}

		// Setup http server using native server
		server := &http.Server{
			Handler: r,
			Addr:    addr,
		}

		// Start the http server in a go routine in order to
		// handle system signal
		go func() {
			log.Infof("Start erlenmeyer on %s", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

		/// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, os.Interrupt)

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Cannot gracefully shutdown erlenmeyer")
		}
	},
}
