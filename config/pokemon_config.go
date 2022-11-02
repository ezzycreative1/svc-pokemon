package config

import (
	"strings"
)

// in mem package variable
var env = "dev"

// call this func for simple feature flag.
// when your feature need to disable in prod
func IsDev() bool {
	return strings.ToLower(env) == "dev"
}

type Database struct {
	Engine          string `json:"engine,omitempty"`
	Host            string `json:"host,omitempty"`
	Port            int    `json:"port,omitempty"`
	Username        string `json:"username,omitempty"`
	Password        string `json:"-"`
	Schema          string `json:"schema,omitempty"`
	MaxIdle         int    `json:"max_idle,omitempty"`
	MaxConn         int    `json:"max_conn,omitempty"`
	ConnMaxLifetime int    `json:"conn_max_lifetime,omitempty"`
	Environment     string `json:"environment,omitempty"`
}

type MySQLConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Schema          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
	Environment     string
}

func (d *Database) DBConfig() *MySQLConfig {
	return &MySQLConfig{
		Host:            d.Host,
		Port:            d.Port,
		User:            d.Username,
		Password:        d.Password,
		Schema:          d.Schema,
		MaxIdleConns:    d.MaxIdle,
		MaxOpenConns:    d.MaxConn,
		ConnMaxLifetime: d.ConnMaxLifetime,
		Environment:     d.Environment,
	}
}

type Redis struct {
	Host         string `json:"host,omitempty"`
	Port         int    `json:"port,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"-"`
	DB           int    `json:"db,omitempty"`
	UseTLS       bool   `json:"use_tls,omitempty"`
	MaxRetries   int    `json:"max_retries"`
	MinIdleConns int    `json:"min_idle_conns"`
	PoolSize     int    `json:"pool_size"`
	PoolTimeout  int    `json:"pool_timeout"`
	MaxConnAge   int    `json:"max_conn_age"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

type PokemonConfig struct {
	HTTPPort int `json:"http_port,omitempty"`
}

type PokemonAuth struct {
	AccessTokenPrivateKey  string `json:"access_token_private_key,omitempty"`
	AccessTokenPublicKey   string `json:"access_token_public_key,omitempty"`
	AccessTokenExpiresIn   string `json:"access_token_expires_in,omitempty"`
	AccessTokenMaxAge      string `json:"access_token_maxage,omitempty"`
	RefreshTokenPrivateKey string `json:"refresh_token_private_key,omitempty"`
	RefreshTokenPublicKey  string `json:"refresh_token_public_key,omitempty"`
	RefreshTokenExpiresIn  string `json:"refresh_token_expires_in,omitempty"`
	RefreshTokenMaxAge     string `json:"refresh_token_maxage,omitempty"`
}
