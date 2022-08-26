package v1

import (
	"context"

	proto_user_model "github.com/theartofdevel/grpc-contracts/gen/go/user_service/model/v1"
	proto_user_service "github.com/theartofdevel/grpc-contracts/gen/go/user_service/service/v1"
)

type userServer struct {
	proto_user_service.UnimplementedUserServiceServer
	u *usecases.UserService
}

func NewUserServer(unimplementedUserServiceServer proto_user_service.UnimplementedUserServiceServer) *userServer {
	return &userServer{UnimplementedUserServiceServer: unimplementedUserServiceServer}
}

func (s *userServer) CreateUser(ctx context.Context, req *proto_user_service.CreateUserRequest) (*proto_user_service.GetUsersResponse, error) {
	res := s.u.CreateUser(req)

	return &proto_user_service.GetUsersResponse{
		Users: []*proto_user_model.User{
			res.ToProto(),
		},
	}, nil
}

func (s *userServer) DeleteUser(ctx context.Context, req *proto_user_service.DeleteUserRequest) (*proto_user_service.UpdateUserResponse, error) {
	res := s.u.DeleteUser(req)
	return res, nil
}
func (s *userServer) GetUsers(ctx context.Context, req *proto_user_service.GetUsersRequest) (*proto_user_service.GetUsersResponse, error) {
	res := s.u.GetUser(req)

	return &proto_user_service.GetUsersResponse{
		Users: []*proto_user_model.User{
			res.ToProto(),
		},
	}, nil
}
