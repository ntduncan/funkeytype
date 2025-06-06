package system

var configPath = "~/.config/funkeytype/config.json"

// @TODO: Updates new config settings
func SaveConfig(config Config) error {
	return nil
}

// Creates new .config/funkeytype path
// Initialize new .config/funkeytype/config.json?
func initConfig() error {
	return nil
}

func LoadConfig() (Config, error) {
	var config Config
	//Read COnfig
	//IF no config, initConfig
	return config, nil
}
