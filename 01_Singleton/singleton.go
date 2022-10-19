package main

type Singleton struct{}

var singleton = &Singleton{}

func getInstance() *Singleton {
	return singleton
}

//func main() {
//	getInstance()
//	fmt.Println(1)
//	return
//}
