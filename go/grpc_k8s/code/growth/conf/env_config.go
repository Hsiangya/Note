package conf

import (
	"encoding/json"
	"log"
	"os"
)

var GlobalConfig *ProjectConfig

const envConfigName = "USER_GROWTH_CONFIG"

type ProjectConfig struct {
	Db struct {
		Engine          string
		Username        string
		Password        string
		Host            string
		Port            int
		Database        string
		Charset         string
		ShowSql         bool
		MaxIdleConns    int
		MaxOpenConns    int
		CoonMaxLifetime int
	}
	Cache struct{}
}

func LoadConfigs() {
	LoadEnvConfig()
}
func LoadEnvConfig() {
	pc := &ProjectConfig{}

	if strConfigs := os.Getenv(envConfigName); len(strConfigs) > 0 {
		if err := json.Unmarshal([]byte(strConfigs), pc); err != nil {
			log.Fatalf("conf.LoadEnvConfig(%s) error=%s\n", envConfigName, err.Error())
			return
		}
	}

	if pc == nil || pc.Db.Username == "" {
		log.Fatalf("empty os.Getenv config", envConfigName)
		return
	}

	GlobalConfig = pc
}
