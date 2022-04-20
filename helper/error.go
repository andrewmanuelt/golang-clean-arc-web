package helper

func ShowError(err error) {
	if err != nil {
		panic(err)
	}
}
