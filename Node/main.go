package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Akongstad/DistributedMutualExclusion/proto"
	"google.golang.org/grpc"
)

type STATE int

var wait *sync.WaitGroup

type Node struct {
	name         string
	id           int
	state        STATE
	timestamp    int
	ports        []string
	replyCounter int
	queue  		Queue
	protoNode    proto.Node
	proto.UnimplementedExclusionServiceServer
}

const (
	RELEASED STATE = iota
	WANTED
	HELD
)

func (n *Node) ReceiveRequest(ctx context.Context, requestMessage *proto.RequestMessage) error {
	log.Printf("Received request from: %s", requestMessage.User.Name)
	if n.state == HELD || (n.state == WANTED && n.timestamp < int(requestMessage.Timestamp)) {

		return nil
	} else {

		return nil
	}
}

func (n *Node) AccessCritical(ctx context.Context, requestMessage *proto.RequestMessage) (proto.ReplyMessage, error) {
	log.Printf("%s, %d, Stamp: %d Requesting access to critical", n.name, n.id, n.timestamp)
	n.state = WANTED
	n.MessageAll(ctx, requestMessage)

	return proto.ReplyMessage{}, nil
}

func (n *Node) ReceiveReply(ctx context.Context, replyMessage *proto.ReplyMessage) error {

	n.replyCounter++

	if n.replyCounter == len(n.ports)-1 {
		n.EnterCriticalSection()
	}

	return nil
}

func (n *Node) EnterCriticalSection() {
	log.Printf("%s entered the critical section", n.name)
	time.Sleep(5 * time.Second)

	n.LeaveCriticalSection()
}

func (n *Node) LeaveCriticalSection() {
	log.Printf("%s is leaving the critical section", n.name)
	n.state = RELEASED
	n.timestamp++
	n.replyCounter = 0
}

func (n *Node) MessageAll(ctx context.Context, msg *proto.RequestMessage) error {
	for i := 0; i < len(n.ports); i++ { //eller kigge på Serf/Consul til at gruppere serverne
		if i != n.id {
			conn, err := grpc.Dial(":" + n.ports[i])
			defer conn.Close()
			if err != nil {
				log.Fatalf("Failed to listen port: %v", err)
			}
			client := proto.NewExclusionServiceClient(conn)

			client.ReceiveRequest(ctx, msg)
		}
	}
	return nil
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

/*func (n *Node) updateClock(otherTimestamp int) {
	if n.timestamp < otherTimestamp {
		return otherTimestamp +1
	} else {
		return n.timestamp
	}
}*/

func main() {
	name := flag.String("N", "Anonymous", "name")
	flag.Parse()
	id := flag.Int("I", 0, "id")
	flag.Parse()

	n := Node{
		name:      *name,
		id:        *id,
		state:     RELEASED,
		protoNode: proto.Node{Id: int32(*id), Name: *name},
	}

	file, _ := os.Open("/ports.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n.ports = append(n.ports, scanner.Text())
	}

	n.listen()

	// OKAY HVORDAN GØR MAN DET HER
	wait.Add(1)
	//go listenToOtherPorts()

	go func() {
		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			if strings.ToLower(scanner.Text()) == "exit" {
				os.Exit(1)
			} else if strings.ToLower(scanner.Text()) == "request access" {
				request := proto.RequestMessage{User: &n.protoNode, Timestamp: int32(n.timestamp)}
				reply, err := n.AccessCritical(context.Background(), &request)
				if err != nil {
					log.Fatalf("Failed to Request %v, %d", err, reply.Timestamp)
					break
				}
			}
		}
	}()
}

/* func listenToOtherPorts(){
	defer wait.Done()

} */
