package main

import (
	"net"
	"strings"
)

//build a user type
type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn

	server *Server
}

//创建一个user的API
func NewUser(conn net.Conn, server *Server) *User {
	//得到user的addr
	userAdd := conn.RemoteAddr().String()
	//建立一个user 并且最后可以返回*User类的地址
	user := &User{
		Name: userAdd,
		Addr: userAdd,
		C:    make(chan string),
		Conn: conn,
		//user需要将用户连接到对应的server上去
		server: server,
	}
	//启动监听当前userchannel消息的构成
	go user.ListenMessage()

	return user
}

// 用户上线业务，上线需要将用户存到onlinemap里面去，然后告知大家上线了
func (this *User) OnLine() {
	//用户上线，将用户添加到OnlineMap中
	this.server.maplock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.maplock.Unlock()

	//广播当前用户上线了
	this.server.BroadCast(this, "Online")
}

// 用户下线业务，需要将存在onlinemap里面的用户给删除
func (this *User) OffLine() {
	//用户上线，将用户添加到OnlineMap中
	this.server.maplock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.maplock.Unlock()

	//广播当前用户上线了
	this.server.BroadCast(this, "Offline")
}

// 建立一个func来sendmsg
func (this *User) SendMsg(msg string) {
	this.Conn.Write([]byte(msg))
}

// 用户发消息业务，需要读取用户的消息，然后输入channel然后再输出给别的用户
// 如果用户希望查找当前在线的用户名称，用户发送who，server向用户发送onlinemap里面的所有消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//当处理map相关的信息的时候需要lock和unlock一下
		this.server.maplock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ": is online\n"
			//向当前的用户传递消息
			this.SendMsg(onlineMsg)
		}
		this.server.maplock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//处理用户更改姓名的需求
		//接受新的名字
		newname := strings.Split(msg, "|")[1]

		//判断newname是否存在
		_, ok := this.server.OnlineMap[newname]
		if ok {
			this.SendMsg("This name has been used. \n")
		} else {
			this.server.maplock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newname] = this
			this.server.maplock.Unlock()

			this.Name = newname
			this.SendMsg("Your username has changed:" + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {

		//得到对方用户名
		//检查用户名是否为空
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("Wrong format, please use\"to|user|message\"format\n")
			return
		}
		//检查用户名是否在user中有对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("We don't have this user\n")
			return
		}
		//得到对方channel
		//向对方发消息
		remoteMessage := strings.Split(msg, "|")[2]
		if remoteMessage == "" {
			this.SendMsg("No content, please resend")
			return
		}
		remoteUser.SendMsg(this.Name + "is speaking to you: " + remoteMessage)

	} else {
		this.server.BroadCast(this, msg)
	}

}

//监听当前user channel的消息，一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		//从channel里面读数据
		msg := <-this.C
		//将message发出去 转化为二进制
		this.Conn.Write([]byte(msg + "\n"))
	}

}
