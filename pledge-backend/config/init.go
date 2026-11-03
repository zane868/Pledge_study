package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

//go:embed sepolia.json
var deployments embed.FS

func init() {
	var err error
	Config, err = Load()
	if err != nil {
		panic(err)
	}
}

// Load does not connect to external services, so tools and tests can load configuration safely.
func Load() (*Conf, error) {
	c := &Conf{
		Mysql:        MysqlConfig{Address: "127.0.0.1", Port: "3306", DbName: "pledge", UserName: "pledge", MaxOpenConns: 20, MaxIdleConns: 5, MaxLifeTime: 300},
		Redis:        RedisConfig{Address: "127.0.0.1", Port: "6379", Db: 0, MaxIdle: 5, MaxActive: 20, IdleTimeout: 180},
		Env:          EnvConfig{Port: "8080", Version: "1", Protocol: "http", DomainName: "localhost", TaskDuration: 120},
		DefaultAdmin: DefaultAdminConfig{Username: "admin"},
		Jwt:          JwtConfig{ExpireTime: 3600},
	}
	if filename := os.Getenv("PLEDGE_CONFIG"); filename != "" {
		if _, err := toml.DecodeFile(filename, c); err != nil {
			return nil, fmt.Errorf("load PLEDGE_CONFIG: %w", err)
		}
	}
	raw, err := deployments.ReadFile("sepolia.json")
	if file := os.Getenv("PLEDGE_DEPLOYMENT_FILE"); file != "" {
		raw, err = os.ReadFile(file)
	}
	if err != nil {
		return nil, fmt.Errorf("load deployment: %w", err)
	}
	if err = json.Unmarshal(raw, &c.Deployment); err != nil {
		return nil, fmt.Errorf("decode deployment: %w", err)
	}
	c.TestNet.ChainId = strconv.FormatUint(c.Deployment.ChainID, 10)
	c.TestNet.PledgePoolToken = c.Deployment.Pool
	c.TestNet.BscPledgeOracleToken = c.Deployment.Oracle
	c.TestNet.NetUrl = os.Getenv("SEPOLIA_RPC_URL")
	values := map[string]*string{
		"PLEDGE_MYSQL_HOST": &c.Mysql.Address, "PLEDGE_MYSQL_PORT": &c.Mysql.Port, "PLEDGE_MYSQL_DATABASE": &c.Mysql.DbName,
		"PLEDGE_MYSQL_USER": &c.Mysql.UserName, "PLEDGE_MYSQL_PASSWORD": &c.Mysql.Password,
		"PLEDGE_REDIS_HOST": &c.Redis.Address, "PLEDGE_REDIS_PORT": &c.Redis.Port, "PLEDGE_REDIS_PASSWORD": &c.Redis.Password,
		"PLEDGE_ADMIN_USER": &c.DefaultAdmin.Username, "PLEDGE_ADMIN_PASSWORD": &c.DefaultAdmin.Password,
		"PLEDGE_JWT_SECRET": &c.Jwt.SecretKey, "PLEDGE_API_PORT": &c.Env.Port, "PLEDGE_API_HOST": &c.Env.Host,
	}
	for name, target := range values {
		if value, ok := os.LookupEnv(name); ok {
			*target = value
		}
	}
	for name, target := range map[string]*int{"PLEDGE_REDIS_DB": &c.Redis.Db, "PLEDGE_JWT_TTL": &c.Jwt.ExpireTime} {
		if v, ok := os.LookupEnv(name); ok {
			n, e := strconv.Atoi(v)
			if e != nil || n < 0 {
				return nil, fmt.Errorf("invalid %s", name)
			}
			*target = n
		}
	}
	if c.Deployment.ChainID != 11155111 || !strings.HasPrefix(c.TestNet.PledgePoolToken, "0x") {
		return nil, fmt.Errorf("invalid Sepolia deployment manifest")
	}
	return c, nil
}

func ValidateAPI() error {
	if len(Config.Jwt.SecretKey) < 32 {
		return fmt.Errorf("PLEDGE_JWT_SECRET must contain at least 32 characters")
	}
	if len(Config.DefaultAdmin.Password) < 12 {
		return fmt.Errorf("PLEDGE_ADMIN_PASSWORD must contain at least 12 characters")
	}
	if Config.Jwt.ExpireTime <= 0 {
		return fmt.Errorf("PLEDGE_JWT_TTL must be positive")
	}
	return nil
}

func SupportedChain(id int) bool { return strconv.Itoa(id) == Config.TestNet.ChainId }
