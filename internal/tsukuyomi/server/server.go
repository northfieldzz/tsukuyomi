package server

import (
	"fmt"
	"os"
)

func Init() error {
	router, err := NewRouter()
	if err != nil {
		return err
	}
	err = router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		return err
	}
	return nil
}
