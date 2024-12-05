package initializer

func Init() {
	initConfig()
	initDB()
	initRepositories()
	initServices()
}
