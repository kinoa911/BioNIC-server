package config

type DBConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:     "root",      //os.Getenv("DB_USER"),
		Password: "password",  //os.Getenv("DB_PASSWORD"),
		Driver:   "mysql",     //os.Getenv("DB_DRIVER"),
		Name:     "test",      //os.Getenv("DB_NAME"),
		Host:     "localhost", //os.Getenv("DB_HOST"),
		Port:     "3306",      //os.Getenv("DB_PORT"),
	}
}
