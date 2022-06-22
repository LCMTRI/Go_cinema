package main

import (
	//"Go_cinema/database"
	"Go_cinema/database"
	pb "Go_cinema/protos"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"io"
	"log"
	"net"

	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ============== gRPC client ===============
// ==========================================
var (
	client pb.ComputeServiceClient
)

// =============== gRPC server ==============
// ==========================================
var users []*pb.UserInfoRes

type userServer struct {
	pb.UnimplementedComputeServiceServer
	UserCollection *mongo.Collection
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

func (s *userServer) GetUser(ctx context.Context, in *pb.Id) (*pb.UserInfoRes, error) {
	log.Printf("Received: %v", in)

	res := &pb.UserInfoRes{}
	for _, user := range users {
		if user.Code == in.Value {
			res = user
			break
		}
	}
	return res, nil
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.UserInfoReq) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	in.Code = strconv.Itoa(rand.Intn(10000))

	inserted, _ := s.UserCollection.InsertOne(ctx, in)
	res.Value = inserted.InsertedID.(primitive.ObjectID).Hex()
	addedUser := &pb.UserInfoRes{Id: res.Value, Code: in.Code, Name: in.Name, Email: in.Email, Password: in.Password, WatchedMovies: in.WatchedMovies}
	users = append(users, addedUser)
	return &res, nil
}

func (s *userServer) UpdateUser(ctx context.Context, in *pb.UserInfoReq) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	for _, user := range users {
		if user.Code == in.Code {
			user.Name = in.Name
			user.Email = in.Email
			user.Password = in.Password
			user.WatchedMovies = in.WatchedMovies
			res.Value = 1
			s.UserCollection.UpdateOne(ctx, bson.M{"code": in.Code}, bson.D{{"$set", in}})
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
			s.UserCollection.DeleteOne(ctx, bson.M{"id": user.Id})
			break
		}
	}
	return &res, nil
}

func (s *userServer) GetWatchedMoviesGateway(ctx context.Context, in *pb.Id) (*pb.MovieInfoList, error) {
	log.Printf("Received user id: %v", in)
	req := &pb.MovieList{}
	var watched_movies []string
	var userAv = false
	for _, user := range users {
		if user.Id == in.Value {
			watched_movies = user.WatchedMovies
			userAv = true
			log.Println()
			break
		}
	}
	if !userAv {
		log.Println("Can't find User with given Id")
		return nil, nil
	}
	req.WatchedMovies = append(req.WatchedMovies, watched_movies...)
	//ctx := context.Background()
	stream, err := client.GetWatchedMoviesUser(ctx, req)
	var movies []*pb.MovieInfoRes
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
		movies = append(movies, row)

	}
	res := pb.MovieInfoList{WatchedMovies: movies}
	return &res, nil
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
	userCollection := database.DB.Database("User_service_database").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
		return
	}

	// gRPC Client to Movie gRPC Server
	conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewComputeServiceClient(conn)
	fmt.Println("Welcom to Movie gRPC Server")

	// gRPC Server port :8002
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterComputeServiceServer(s, &userServer{
		UserCollection: userCollection,
	})

	log.Printf("User gRPC Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
