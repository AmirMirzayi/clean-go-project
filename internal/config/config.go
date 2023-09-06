package config

type HttpServer struct{
	Address string
	Port int
}

type Config struct{
	HttpServer
}
