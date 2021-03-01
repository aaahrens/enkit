package machinist

import (
	"encoding/json"
	"fmt"
	"github.com/enfabrica/enkit/lib/srand"
	"github.com/enfabrica/enkit/lib/token"
	machinist "github.com/enfabrica/enkit/machinist/rpc"
	"math/rand"
	"net"
)

func NewServerRequest() *ServerRequest {
	return &ServerRequest{}
}

type Server struct {
	Config  ServerFlagSet
	Encoder token.BinaryEncoder
	Req     *ServerRequest
}

func NewServer(req *ServerRequest) *Server {
	return &Server{
		Req: req,
	}
}
func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s Server) Poll(server machinist.Controller_PollServer) error {
	request, err := server.Recv()
	if err != nil {
		return err
	}
	ping := request.GetPing()
	fmt.Println(ping.Payload)
	response := machinist.PollResponse{
		Resp: &machinist.PollResponse_Pong{
			Pong: &machinist.ActionPong{
				Payload: []byte("hello world"),
			},
		},
	}
	return server.Send(&response)
}

func (s Server) Upload(server machinist.Controller_UploadServer) error {
	panic("implement me")
}

func (s Server) Download(request *machinist.DownloadRequest, server machinist.Controller_DownloadServer) error {
	panic("implement me")
}

type invitationToken struct {
	Addresses []string
	Port      int
}

func (s Server) GenerateInvitation() ([]byte, error) {
	nats, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var attachedIpAddresses []string
	for _, nat := range nats {
		addresses, err := nat.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addresses {
			if tcpNat, ok := addr.(*net.IPNet); ok {
				attachedIpAddresses = append(attachedIpAddresses, tcpNat.IP.String())
			}
			//not sure if should try and handle non ipAddr interfaces
		}
	}
	i := invitationToken{
		Port:      s.Config.Port,
		Addresses: attachedIpAddresses,
	}
	jsonString, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	encodedToken, err := s.Encoder.Encode(jsonString)
	if err != nil {
		return nil, err
	}
	return encodedToken, nil
}

type ServerRequest struct {
	Port    int
	Encoder token.BinaryEncoder
}

func (rs *ServerRequest) WithPort(p int) *ServerRequest {
	rs.Port = p
	return rs
}

func (rs ServerRequest) Start() (*Server, error) {
	return &Server{}, nil
}

func (rs *ServerRequest) UseEncoder(encoder token.BinaryEncoder) *ServerRequest {
	en, err := token.NewSymmetricEncoder(rand.New(srand.Source))
	if err != nil {
		panic(err)
	}
	rs.Encoder = en
	return rs
}

func (s *Server) Nodes() {

}
