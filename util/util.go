package util

// Util test
type Util struct{}

//CheckError test
func (*Util) CheckError(err error) {
	if err != nil {
		panic(err)
	}

}
