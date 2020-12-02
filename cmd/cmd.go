package cmd

import (
	"net/http"
	"os"

	"github.com/shuwenhe/shuwen-checklist/config"
	db "github.com/shuwenhe/shuwen-checklist/database"
	"github.com/shuwenhe/shuwen-checklist/model"
	"github.com/shuwenhe/shuwen-checklist/router"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error { // Run but returns an error
		_, err := db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		db.DB.AutoMigrate(&model.Todo{}) // struct -> table

		r := router.SetupRouter() // setup router
		r.Run()

		return http.ListenAndServe(viper.GetString("addr"), nil) // listen and serve
	}

	return rootCmd.Execute()

}
