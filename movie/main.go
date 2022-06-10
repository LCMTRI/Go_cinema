package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "Go_cinema/protos"

	"google.golang.org/grpc"
)

// ============== gRPC server ===============
// ==========================================
var movies []*pb.MovieInfo

type movieServer struct {
	pb.UnimplementedComputeServiceServer
}

func (s *movieServer) GetMovies(in *pb.Empty, stream pb.ComputeService_GetMoviesServer) error {
	log.Printf("Received: %v", in)
	for _, movie := range movies {
		if err := stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieServer) GetMovie(ctx context.Context, in *pb.Id) (*pb.MovieInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.MovieInfo{}
	for _, movie := range movies {
		if movie.GetId() == in.GetValue() {
			res = movie
			break
		}
	}
	return res, nil
}

func (s *movieServer) CreateMovie(ctx context.Context, in *pb.MovieInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(10000))
	in.Id = res.GetValue()
	movies = append(movies, in)
	return &res, nil
}

func (s *movieServer) UpdateMovie(ctx context.Context, in *pb.MovieInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for index, movie := range movies {
		if movie.GetId() == in.GetId() {
			movies = append(movies[:index], movies[index+1:]...)
			in.Id = movie.GetId()
			movies = append(movies, in)
			res.Value = 1
			break
		}
	}
	return &res, nil
}

func (s *movieServer) DeleteMovie(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for index, movie := range movies {
		if movie.GetId() == in.GetValue() {
			movies = append(movies[:index], movies[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func initMovies() {
	movie1 := &pb.MovieInfo{Id: "1", Isbn: "0593310438", Title: "The Batman", Director: &pb.Director{Firstname: "Matt", Lastname: "Damon"}}
	movie2 := &pb.MovieInfo{Id: "2", Isbn: "3430220302", Title: "Venom", Director: &pb.Director{Firstname: "Zhong", Lastname: "Cina"}}

	movies = append(movies, movie1)
	movies = append(movies, movie2)
}

func main() {
	// gRPC Server port :8001
	initMovies()
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterComputeServiceServer(s, &movieServer{})

	log.Printf("Movie gRPC Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
