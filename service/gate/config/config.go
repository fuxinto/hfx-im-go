package config

type Config struct {
	Log    Log
	Server Server
}

type Server struct {
	Port         string
	MaxConnCount int
}

//
//type IMRpc struct {
//	Addr    string
//	Network string
//}

type Log struct {
	PathName    string
	ServiceName string
}
