package connections

func CleanUp() {
	DB.Close()
}
