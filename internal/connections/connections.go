package connections

func CleanUp() {
	DB.Close()
    Redis.Close()
}
