package config

import (
	"flag"
	"os"
)

// EnvironmentVariable -  Type to environment variables.
type EnvironmentVariable string

const (
	//Base - Define data base app. Ex.: MONGO
	Base EnvironmentVariable = "BASE"
	// Port -  Server port application.
	Port     EnvironmentVariable = "PORT"
	BaseName EnvironmentVariable = "BASE_NAME"
	// BaseAddress - Data base IP address or DNS.
	BaseAddress = "BASE_ADDRESS"
	// BasePort - Data base port.
	BasePort = "BASE_PORT"
	// BaseUser - User data base user.
	BaseUser = "BASE_USER"
	// BasePassword - Data base password.
	BasePassword = "BASE_PASSWORD"
)

type configVar struct {
	name  EnvironmentVariable
	value string
	usage string
}

var configVars []configVar = []configVar{
	{name: Base, value: "MONGO", usage: "Define data base app. Ex.: MONGO"},
	{name: Port, value: "8080", usage: "Server port application."},
	{name: BaseName, value: "b2wtest", usage: "Define base name. Ex.: b2wtest"},
	{name: BaseAddress, value: "127.0.0.1", usage: "Data base IP address or DNS. Ex.: 127.0.0.1"},
	{name: BasePort, value: "27017", usage: "Data base port. Ex.: 27017"},
	{name: BaseUser, value: "", usage: "User data base user. Ex.: mongouser"},
	{name: BasePassword, value: "", usage: "Data base password. Ex.: mongopass"},
}

func setVar(envVar EnvironmentVariable, value string) {
	for i := range configVars {
		if configVars[i].name == envVar {
			configVars[i].value = value
		}
	}
}

func getVar(envVar EnvironmentVariable) *configVar {
	for i := range configVars {
		if configVars[i].name == envVar {
			return &configVars[i]
		}
	}
	return nil
}

//FlagParse - Flags parsing and set values.
func FlagParse() {
	var values []*string
	if !flag.Parsed() {
		for i := range configVars {
			values = append(values, flag.String(string(configVars[i].name), configVars[i].value, configVars[i].usage))
		}
		flag.Parse()
	}
	for i := range configVars {
		configVars[i].value = *values[i]
	}
}

// EnvVal - Find to environment variable value
// or return default value of variable.
func EnvVal(variable EnvironmentVariable) string {
	if value := os.Getenv(string(variable)); value != "" {
		setVar(variable, value)
	}
	if conf := getVar(variable); conf != nil {
		return conf.value
	}
	return ""
}
