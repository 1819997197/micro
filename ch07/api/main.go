package main

func main() {
	router := initRouter()

	//默认情况下，它使用:8080
	router.Run()
	// router.Run(":8000") 硬编码端口
}
