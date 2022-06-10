package main

import (
	pb "Go_cinema/protos"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ============== gRPC client ===============
// ==========================================
var client pb.ComputeServiceClient

func runGetMovies(client pb.ComputeServiceClient) []*pb.MovieInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var Movies []*pb.MovieInfo

	req := &pb.Empty{}
	stream, err := client.GetMovies(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
		}
		Movies = append(Movies, row)
	}
	return Movies
}

func runGetMovie(client pb.ComputeServiceClient, movieid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: movieid}
	res, err := client.GetMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
	}
	log.Printf("MovieInfo: %v", res)
}

func runCreateMovie(client pb.ComputeServiceClient, isbn string,
	title string, firstname string, lastname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.MovieInfo{Isbn: isbn, Title: title,
		Director: &pb.Director{Firstname: firstname, Lastname: lastname}}
	res, err := client.CreateMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateMovie(_) = _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateMovie Id: %v", res)
	} else {
		log.Printf("CreateMovie Failed")
	}
}

func runUpdateMovie(client pb.ComputeServiceClient, movieid string,
	isbn string, title string, firstname string, lastname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.MovieInfo{Id: movieid, Isbn: isbn,
		Title: title, Director: &pb.Director{Firstname: firstname, Lastname: lastname}}
	res, err := client.UpdateMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateMovie(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateMovie Success")
	} else {
		log.Printf("UpdateMovie Failed")
	}
}

func runDeleteMovie(client pb.ComputeServiceClient, movieid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: movieid}
	res, err := client.DeleteMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteMovie(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteMovie Success")
	} else {
		log.Printf("DeleteMovie Failed")
	}
}

// func getMoviesForUser(c echo.Context) error {
// 	var Movies = runGetMovies(client)
// 	json.NewEncoder(c.Response().Writer).Encode(Movies)
// 	return c.String(http.StatusOK, "")
// }

// =============== gRPC server ==============
// ==========================================
var users []*pb.UserInfo

type userServer struct {
	pb.UnimplementedComputeServiceServer
}

func (s *userServer) GetUsers(in *pb.Empty, stream pb.ComputeService_GetUsersServer) error {
	log.Printf("Received: %v", in)
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *userServer) GetUser(ctx context.Context, in *pb.Id) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.UserInfo{}
	for _, user := range users {
		if user.GetId() == in.GetValue() {
			res = user
			break
		}
	}
	return res, nil
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.UserInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(10000))
	in.Id = res.GetValue()
	users = append(users, in)
	return &res, nil
}

func (s *userServer) UpdateUser(ctx context.Context, in *pb.UserInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for index, user := range users {
		if user.GetId() == in.GetId() {
			users = append(users[:index], users[index+1:]...)
			in.Id = user.GetId()
			users = append(users, in)
			res.Value = 1
			break
		}
	}
	return &res, nil
}

func (s *userServer) DeleteUser(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for index, user := range users {
		if user.GetId() == in.GetValue() {
			users = append(users[:index], users[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func main() {
	// gRPC Server port :8002
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterComputeServiceServer(s, &userServer{})

	log.Printf("User gRPC Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// gRPC Client to Movie gRPC Server
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewComputeServiceClient(conn)
	fmt.Println("Welcom to Movie gRPC Server")

	// runGetMovies(client)
	// runGetMovie(client, "1")
	// runCreateMovies(client, "0868968108", "White Chicks", "Adam", "Sandlers")
	// runUpdateMovies(client, userid, isbn, title, firstname, lastname)
	// runDeleteMovies(client, userid)
}
