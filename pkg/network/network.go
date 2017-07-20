package network

import (
	"../node"
	"log"
)

type Network struct {
	Nodes        []node.Node
	Name         string
	AccessTokens []string
}

var (
	Local Network
)

func InitLocal() {
	Local = Network{
		Name: "local",
	}

	log.Println("Initialized local network as", Local.Name)
}

func (n Network) PushNode(node node.Node) {
	n.Nodes = append(n.Nodes, node)
}

func (n Network) PushAccessToken(token string) {
	n.AccessTokens = append(n.AccessTokens, token)
}
