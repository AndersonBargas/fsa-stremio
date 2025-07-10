package configs

type BasicConfigs interface {
	Port() string
	DBPath() string
}
