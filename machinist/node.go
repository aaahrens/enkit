package machinist

type Node struct {
	Req *NodeRequest
}

func NewNode(req *NodeRequest) *Node {
	return &Node{
		Req: req,
	}
}
func (receiver Node) Start() {

}

func (receiver Node) Stop() {

}

type NodeRequest struct {
	Token string
}

func NewNodeRequest() *NodeRequest {
	return &NodeRequest{}
}

func (nr *NodeRequest) WithToken(t string) *NodeRequest {
	nr.Token = t
	return nr
}
