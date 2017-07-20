package node

import (
	"../resp"
	"../token"
	"log"
)

type Node struct {
	Name        string
	Ip          string
	Port        string
	LocalToken  string
	RemoteToken string
}

var (
	this Node
)

func CreateIdentity() *Node {
	this = Node{
		Name: "snwfdhmp",
		Ip:   "127.0.0.1",
		Port: "2048",
	}

	return &this
}

func (n *Node) GetAddr() string {
	return n.Ip + ":" + n.Port
}

func (n *Node) GetURL() string {
	return "http://" + n.Ip + ":" + n.Port
}

func (n *Node) Tap() {
	addr := n.GetURL()
	log.Println("Attempt to tap", n.GetAddr(), "on", addr)

	n.RemoteToken = resp.Body(addr + "/tap")
}

func (n *Node) Claim() {
	n.LocalToken = token.Rand()

	resp.Body(n.GetURL() + "/claim/" + n.RemoteToken + "/" + n.LocalToken + "/" + this.Port + "?name=" + this.Name)
}
