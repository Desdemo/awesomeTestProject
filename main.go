package main

import (
	"awesomeTestProject/initRouter"
)

func main()  {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
