package server

func Init() error {
	router, err := NewRouter()
	if err != nil {
		return err
	}
	// TODO: listening portの環境変数化
	err = router.Run()
	if err != nil {
		return err
	}
	return nil
}
