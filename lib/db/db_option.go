package db

import "os"

type DbOptions struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Name    string `json:"name"`
	SslMode string `json:"sslmode"`
	Schema  string `json:"schema"`
}

func DefaultDbOptions() *DbOptions {
	return &DbOptions{
		Host:    "localhost",
		Port:    "5432",
		User:    "postgres",
		Pass:    "",
		Name:    "bodycms",
		SslMode: "disable",
		Schema:  "public",
	}
}

func DbOptionsFromEnv() *DbOptions {
	return &DbOptions{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Name:    os.Getenv("DB_NAME"),
		SslMode: os.Getenv("DB_SSLMODE"),
		Schema:  os.Getenv("DB_SCHEMA"),
	}
}

func (o *DbOptions) GetConnectionString() string {
	return "host=" + o.Host + " port=" + o.Port + " user=" + o.User + " password=" + o.Pass + " dbname=" + o.Name + " sslmode=" + o.SslMode + " search_path=" + o.Schema
}

func (o *DbOptions) Assign(other *DbOptions) *DbOptions {
	if other.Host != "" {
		o.Host = other.Host
	}
	if other.Port != "" {
		o.Port = other.Port
	}
	if other.User != "" {
		o.User = other.User
	}
	if other.Pass != "" {
		o.Pass = other.Pass
	}
	if other.Name != "" {
		o.Name = other.Name
	}
	if other.SslMode != "" {
		o.SslMode = other.SslMode
	}
	if other.Schema != "" {
		o.Schema = other.Schema
	}

	return o
}
