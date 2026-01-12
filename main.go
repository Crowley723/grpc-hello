package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-hello/greet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultPort = ":50051"
)

// server implements the GreetService
type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received SayHello request from: %s", req.Name)
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s! Welcome to gRPC!", req.Name),
	}, nil
}

func (s *server) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	log.Printf("Received SayGoodbye request from: %s", req.Name)
	return &pb.GoodbyeResponse{
		Message: fmt.Sprintf("Goodbye, %s! See you soon!", req.Name),
	}, nil
}

func runServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	log.Printf("Server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func runClient(serverAddr string, name string) {
	// Set up a connection to the server
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call SayHello
	log.Printf("Calling SayHello with name: %s", name)
	helloResp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	fmt.Printf("Response: %s\n", helloResp.Message)

	// Call SayGoodbye
	log.Printf("Calling SayGoodbye with name: %s", name)
	goodbyeResp, err := client.SayGoodbye(ctx, &pb.GoodbyeRequest{Name: name})
	if err != nil {
		log.Fatalf("SayGoodbye failed: %v", err)
	}
	fmt.Printf("Response: %s\n", goodbyeResp.Message)
}

func main() {
	mode := flag.String("mode", "server", "Run mode: 'server' or 'client'")
	port := flag.String("port", defaultPort, "Server port (default: :50051)")
	serverAddr := flag.String("addr", "localhost:50051", "Server address for client mode")
	name := flag.String("name", "World", "Name to send in greeting (client mode only)")
	flag.Parse()

	switch *mode {
	case "server":
		runServer(*port)
	case "client":
		runClient(*serverAddr, *name)
	default:
		log.Fatalf("Invalid mode: %s. Use 'server' or 'client'", *mode)
	}
}
