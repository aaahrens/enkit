package machinist_test

import (
	"fmt"
	"github.com/enfabrica/enkit/lib/khttp/ktest"
	"github.com/enfabrica/enkit/machinist"
	"testing"
)

func TestRunServerAndGenerateInvite(t *testing.T) {
	netAddr, err := ktest.AllocatePort()
	if err != nil {
		t.Fatal(err.Error())
	}
	serverReq := machinist.NewServerRequest().
		WithPort(netAddr.Port).
		UseEncoder(nil)

	server := machinist.NewServer(serverReq)
	go server.Start()
	defer server.Stop()

	inviteString, err := server.GenerateInvitation()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(inviteString)
	//debug inviteString
	nodeRequest := machinist.NewNodeRequest().
		WithToken(string(inviteString))

	node := machinist.NewNode(nodeRequest)
	go node.Start()
	defer node.Stop()

	currentNodeList := server.Nodes()
	for _, node := currentNodeList {
		fmt.Println(node.Name)
		fmt.Println(node.Tags)
		fmt.Println(node.locations)
		fmt.Println(node.Port)
	}
}

func TestJoinNodes(t *testing.T) {

}

func TestJoinNodesAndKill(t *testing.T) {

}

func TestJoinNodesAndRejoin(t *testing.T) {

}
