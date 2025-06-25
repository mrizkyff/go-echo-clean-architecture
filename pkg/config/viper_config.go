package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func NewRemoteConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	// TODO: Uncomment this to use consul
	// v.AddRemoteProvider("consul", "localhost:8500", "pendekin/config")
	// v.SetConfigType("yaml")
	// logrus.Info("*** Connecting to consul server ***")
	// err := v.ReadRemoteConfig()
	// if err != nil {
	// 	logrus.Errorf(" *** Error reading configuration %s ***", err.Error())
	// }

	err := v.ReadInConfig()
	if err != nil {
		logrus.Errorf(" *** Error reading configuration %s ***", err.Error())
	}

	logrus.Info("*** Success to retrieve config from consul ***")
	return v
}
