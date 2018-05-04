package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type configurations struct {
	domain              string
	serviceAccount      string
	requiredTargetGroup string
	servicePassword     string
	serviceTable        string
	userTable           string
	incidentTable       string
	maintenanceTable    string
	env                 string
	adderToken          string
	ticketURL           string
	sender              string
	emailDomain         string
	searchDomain        string
	supportnumber       string
	externalUserTable   string
}

func (c *configurations) setConfigs() {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "LOCAL"
	}
	fmt.Println("Working in: ", mode, " mode")
	viper.SetConfigName(mode)
	viper.AddConfigPath("configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("No configuration file loaded")
	}
	(*c).domain = viper.GetString("domain")
	(*c).serviceAccount = viper.GetString("serviceAccount")
	(*c).requiredTargetGroup = viper.GetString("requiredTargetGroup")
	(*c).servicePassword = os.Getenv("SERVICEPASSWORD")
	(*c).adderToken = os.Getenv("ADDERTOKEN")
	(*c).serviceTable = viper.GetString("serviceTable")
	(*c).userTable = viper.GetString("userTable")
	(*c).incidentTable = viper.GetString("incidentTable")
	(*c).maintenanceTable = viper.GetString("maintenanceTable")
	(*c).env = viper.GetString("env")
	(*c).ticketURL = viper.GetString("trackerURL")
	(*c).sender = viper.GetString("sender")
	(*c).emailDomain = viper.GetString("emailDomain")
	(*c).searchDomain = viper.GetString("searchDomain")
	(*c).supportnumber = viper.GetString("supportnumber")
	(*c).externalUserTable = viper.GetString("externalUserTable")
}

var currentConfig configurations
