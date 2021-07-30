package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

//Server的类型
type Server struct {
	IP   string
	Port int
	//增加online user的map表格//同步机制
	OnlineMap map[string]*User
	maplock   sync.RWMutex

	//消息广播
	Message chan string
}

//一个函数可以在主函数中调用来创建struct
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: (make(map[string]*User)),
		Message:   make(chan string),
	}
	return server
}

//监听user消息一旦有消息就发送给全部的user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//将msg发送给全部的在线的User
		this.maplock.Lock()
		for _, cil := range this.OnlineMap {
			cil.C <- msg
		}
		this.maplock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := " [" + user.Addr + "] " + user.Name + " : " + msg
	this.Message <- sendMsg
}

//处理链接的业务
func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("Successfully Connected")
	//先建立user
	user := NewUser(conn, this)
	//先给user上线
	user.OnLine()

	isLive := make(chan bool)

	//接受用户端传递的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.OffLine()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}
			//提取用户的消息并且除去最后的“\n”
			msg := string(buf[:n-1])
			//针对用户的消息进行广播
			user.DoMessage(msg)
			isLive <- true
		}

	}()

	//当前handler阻塞
	//用一个select的方式，如果用户10秒钟不运行的话就将用户踢出
	for {
		select {
		//如果islive里面有消息的话就往管道里面存
		case <-isLive:
			//如果被触发不做任何事情，等待select被重置即可

		//如果islive有消息select被触发了那么里面的case就会重置，但是如果没有被触发那么10秒钟后这个case就会被触发
		//用timer做一个定时器，定时触发
		case <-time.After(time.Second * 300):
			user.SendMsg("Kick off")
			//销毁用户的channel
			close(user.C)
			//关闭连接
			conn.Close()
			//退出当前的handler
			return // runtime.Goexit()
		}

	}

}

//启动server 服务
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	//用一个defer去close整个lister
	defer listener.Close()
	//启动监听message的goroutine
	go this.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		//handler 用一个go去并发一些
		go this.Handler(conn)
	}

}
