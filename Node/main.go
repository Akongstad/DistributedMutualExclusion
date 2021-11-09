package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	//"github.com/Akongstad/DistributedMutualExclusion/proto"
)

type Node struct {
	port       string
	id         string
	otherPorts []string
}

func (n *Node) listen() {
	lis, err := net.Listen("tcp", n.port)
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	//course_proto.RegisterCourseServiceServer(grpcServer, &s)

	log.Printf("node listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed serve server: %v", err)
	}
}

func main() {
	port := flag.String("P", "Anonymous", "port")
	flag.Parse()
	id := flag.String("I", "Anonymous", "id")
	flag.Parse()
	n := Node{
		port: *port,
		id:   *id,
	}

	n.listen()
}
