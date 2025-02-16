package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// Gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
