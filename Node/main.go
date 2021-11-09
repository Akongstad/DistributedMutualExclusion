package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/Akongstad/DistributedMutualExclusion/proto"
	"google.golang.org/grpc"
)

type Node struct {
	name string
	id int
	state int
	timestamp int
	ports []string
	proto.UnimplementedExclusionServiceServer

	/*
		0 = RELEASED
		1 = WANTED
		2 = HELD
	*/
}

func Max(a int32, b int32) int32 {
	if b > a {
		return b
	}
	return a
}

func (n *Node) AccessCritical(ctx context.Context, requestMessage *proto.RequestMessage) (proto.ReplyMessage, error) {
	n.state = 1
	n.MessageAll(requestMessage)
	return proto.ReplyMessage{}, nil	
}

func (n *Node) Reply(ctx context.Context, replyMessage *proto.RequestMessage) (proto.ReplyMessage, error) {
	
	
	
	return proto.ReplyMessage{}, nil
}

func (n *Node) MessageAll(msg *proto.RequestMessage) (*proto.ReplyMessage){
	for i := 0; i < len(n.ports); i++{ //eller kigge på Serf/Consul til at gruppere serverne
		if i!= n.id{
			conn, err := grpc.Dial(":" + n.ports[i])
			conn.

		}
	}	
}

func (n *Node) listen() {
	lis, err := net.Listen("tcp", n.ports[n.id])
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterExclusionServiceServer(grpcServer, n)

	log.Printf("node listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed serve server: %v", err)
	}
}

func main() {
	name := flag.String("N", "Anonymous", "name")
	flag.Parse()
	id := flag.Int("I", 0, "id")
	flag.Parse()
	
	n := Node{
		name: *name,
		id:   *id,
		state: 0,
	}

	file, _ := os.Open("/ports.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		n.ports = append(n.ports, scanner.Text())
	}

	n.listen()

	// OKAY HVORDAN GØR MAN DET HER
	wait.Add(1)
	go listenToOtherPorts()
}

func listenToOtherPorts(){
	defer wait.Done()

}