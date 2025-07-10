package configs

import "os"

type basicConfigsEnv struct {
	dbPath string

	port string
}

func NewBasicConfigsEnv() (*basicConfigsEnv, error) {
	dbPath := os.Getenv("FSA_DB_PATH")
	port := os.Getenv("FSA_PORT")

	c := &basicConfigsEnv{
		dbPath: dbPath,
		port:   port,
	}

	return c, nil
}

func (c *basicConfigsEnv) DBPath() string {
	return c.dbPath
}

func (c *basicConfigsEnv) Port() string {
	return c.port
}
