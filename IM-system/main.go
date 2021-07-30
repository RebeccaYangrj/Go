package main

//用于创建接口以及调用start
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
