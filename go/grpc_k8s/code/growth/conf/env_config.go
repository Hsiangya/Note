package conf

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	_, filename, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(filename), "../")
	envPath := filepath.Join(rootDir, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Println("Warning: Error loading .env file:", err)
		// 继续执行，因为环境变量可能通过其他方式设置
	}

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
