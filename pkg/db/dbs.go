package db

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	DbUser     string
	DbType     string
	DbPswd     string
	DbHost     string
	DbPort     string
	DbName     string
	DbProtocol string
	Collection string
	testDBHost string
	testDBName string
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.DbProtocol, c.DbHost, c.DbPort)
}

func (c *Config) getTestDBConnStr() string {
	return c.getDBConnStr(c.testDBName, c.testDBHost, c.DbPort)
}

func (c *Config) getDBConnStr(dbName, dbHost, dbPort string) string {
	return fmt.Sprintf("%s%s:%s", dbName, dbHost, dbPort)
}

func Get() *Config {
	config := &Config{}
	// Read the necessary environment variables
	flag.StringVar(&config.DbUser, "dbuser", os.Getenv("MONGODB_USER"), "DB user name")
	flag.StringVar(&config.DbPswd, "dbpswd", os.Getenv("MONGODB_PASSWORD"), "DB pass")
	flag.StringVar(&config.DbPort, "dbport", os.Getenv("MONGODB_PORT"), "DB port")
	flag.StringVar(&config.DbHost, "dbhost", os.Getenv("MONGODB_HOST"), "DB host")
	flag.StringVar(&config.DbName, "dbname", os.Getenv("MONGODB_DB"), "DB name")
	flag.StringVar(&config.DbType, "dbtype", os.Getenv("DB_TYPE"), "DB  type")
	flag.StringVar(&config.DbProtocol, "dbprotocol", os.Getenv("MONGODB_DB_PROTOCOL"), "Protocol the  db uses")
	flag.StringVar(&config.Collection, "collection", os.Getenv("MONGODB_COLLECTION"), "DB collection")
	flag.StringVar(&config.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&config.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")

	flag.Parse()
	return config
}
