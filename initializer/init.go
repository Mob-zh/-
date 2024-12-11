package initializer

func Init() {
	initConfig()
	initMysql()
	initRedis()
	initRepositories()
	initServices()
}
