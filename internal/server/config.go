package server

type Config struct {
	AppName string `json:"app_name"`
	Host    string `json:"host"`
	Port    string `json:"port"`
}
