package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	//"net/http"
	//"time"

	database "Go_cinema/database"

	pb "Go_cinema/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc"
)

// ============== gRPC server ===============
// ==========================================
var (
	movies []*pb.MovieInfoRes
)

type movieServer struct {
	pb.UnimplementedComputeServiceServer
	movieCollection *mongo.Collection
}

func (s *movieServer) GetMovies(in *pb.Empty, stream pb.ComputeService_GetMoviesServer) error {
	log.Printf("Received: %v", in)
	for _, movie := range movies {
		// movieTmp := &pb.MovieInfoRes{Id: movie.Id, Isbn: movie.Isbn, Title: movie.Title,
		// 	Director: &pb.Director{Firstname: movie.Director.Firstname, Lastname: movie.Director.Lastname}}
		if err := stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieServer) GetMovie(ctx context.Context, in *pb.Id) (*pb.MovieInfoRes, error) {
	log.Printf("Received: %v", in)

	res := &pb.MovieInfoRes{}
	for _, movie := range movies {
		if movie.Id == in.Value {
			// movieTmp := &pb.MovieInfoRes{Id: movie.Id, Isbn: movie.Isbn, Title: movie.Title,
			// 	Director: &pb.Director{Firstname: movie.Director.Firstname, Lastname: movie.Director.Lastname}}
			res = movie
			break
		}
	}
	return res, nil
}

func (s *movieServer) CreateMovie(ctx context.Context, in *pb.MovieInfoReq) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	in.Code = strconv.Itoa(rand.Intn(10000))
	inserted, _ := s.movieCollection.InsertOne(ctx, in)
	res.Value = inserted.InsertedID.(primitive.ObjectID).Hex()
	addedMovie := &pb.MovieInfoRes{Id: res.Value, Code: in.Code, Isbn: in.Isbn, Title: in.Title,
		Director: &pb.Director{Firstname: in.Director.Firstname, Lastname: in.Director.Lastname}}
	movies = append(movies, addedMovie)
	return &res, nil
}

func (s *movieServer) UpdateMovie(ctx context.Context, in *pb.MovieInfoReq) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for _, movie := range movies {
		if movie.Code == in.Code {
			movie.Isbn = in.Isbn
			movie.Title = in.Title
			movie.Director.Firstname = in.Director.Firstname
			movie.Director.Lastname = in.Director.Lastname
			res.Value = 1
			s.movieCollection.UpdateOne(ctx, bson.M{"code": in.Code}, bson.D{{"$set", in}})
			break
		}
	}
	return &res, nil
}

func (s *movieServer) DeleteMovie(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for index, movie := range movies {
		if movie.Code == in.Value {
			movies = append(movies[:index], movies[index+1:]...)
			res.Value = 1
			s.movieCollection.DeleteOne(ctx, bson.M{"code": movie.Code})
			break
		}
	}

	return &res, nil
}

func (s *movieServer) GetWatchedMoviesUser(in *pb.MovieList, stream pb.ComputeService_GetWatchedMoviesUserServer) error {
	for _, movieId := range in.WatchedMovies {
		for _, movie := range movies {
			if movieId == movie.Code {
				if err := stream.Send(movie); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return nil
}

func main() {
	// connect to database
	// var URI = "mongodb+srv://lecaominhtri0701:lecaominhtri@cluster0.x7apzya.mongodb.net/?retryWrites=true&w=majority"
	// fmt.Println("Starting the server...")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		fmt.Println("Disconnected")
	// 		panic(err)
	// 	}
	// }()
	database.Connect()
	movieCollection := database.DB.Database("Movie_service_database").Collection("Movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := movieCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal(err)
		return
	}
	// gRPC Server port :8001
	fmt.Printf("movies: %v", movies)
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterComputeServiceServer(s, &movieServer{
		movieCollection: movieCollection,
	})

	log.Printf("Movie gRPC Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
