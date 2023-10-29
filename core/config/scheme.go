package config

import (
	"fmt"
	"time"
)

type Configuration struct {
	Docker   Docker   `mapstructure:"docker"`
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Token    Token    `mapstructure:"token"`
	Redis    Redis    `mapstructure:"redis"`
	Email    Email    `mapstructure:"email"`
}

type Docker struct {
	Redis    RedisD   `mapstructure:"redis"`
	Postgres Postgres `mapstructure:"postgres"`
}

type RedisD struct {
	Image string `mapstructure:"image"`
}
type Postgres struct {
	Container string `mapstructure:"container"`
	Image     string `mapstructure:"image"`
}

type Database struct {
	Driver   string `mapstructure:"driver"`
	Engine   string `mapstructure:"engine"`
	Name     string `mapstructure:"name"`
	SslMode  string `mapstructure:"ssl_mode"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Token struct {
	SecretKey   string        `mapstructure:"secret_key"`
	ExpDuration time.Duration `mapstructure:"exp_duration"`
}

type Email struct {
	SenderName     string        `mapstructure:"sender_name"`
	SenderAddress  string        `mapstructure:"sender_address"`
	SenderPassword string        `mapstructure:"sender_password"`
	ExpDuration    time.Duration `mapstructure:"exp_duration"`
}

func (c *Database) GetDBSource() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.Username, c.Password, c.Name, c.Port, c.SslMode)
}

func (r *Redis) GetRedisAddress() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}
