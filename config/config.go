package config

var (
	DatabaseURL      string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
)

func LoadConfig() {
	//TODO: refine this to load configuration from config.local.json
	DatabaseURL = "mymongo:27017"
	//DatabaseURL = "127.0.0.1:27017"
	DatabaseName = "mytraining"
	DatabaseUsername = ""
	DatabasePassword = ""
}
