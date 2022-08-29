package repos

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MorselShogiew/UsersManager/proto/user"
	"github.com/MorselShogiew/UsersManager/redis"
)

type UserDBRepo interface {
	SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error)
	DeleteUser(rctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	GetUsers(req *pb.GetUsersRequest, stream pb.UserRepo_GetUsersServer) error
}

type userDB struct {
	pb.UnimplementedUserRepoServer
	db    *DBManager
	redis *redis.RedisManager
	p     *KafkaProducer
	ctx   context.Context
}

func (s *userDB) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	user := req.GetData()
	email := user.GetEmail()
	name := user.GetName()
	log.Printf("Save User: %v, %v", name, email)
	var id int32
	sqlStatement := `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id
	`
	err := s.db.db.QueryRow(sqlStatement, email, name).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("Failed user saving: %v", err)
	}
	log.Printf("Created User Id: %v", id)
	user.Id = id
	go s.p.Produce(fmt.Sprintf("Created user with id %d, emain %s, name %s", user.Id, user.Email, user.Name))
	return &pb.SaveUserResponse{Status: pb.Status_Success, Data: user}, nil
}

func (s *userDB) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := req.GetId()
	if id <= 0 {
		return &pb.DeleteUserResponse{Status: pb.Status_RequestError}, nil
	}
	log.Printf("Delete User by id: %v", id)
	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;`
	res, err := s.db.db.Exec(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	c, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return &pb.DeleteUserResponse{Status: pb.Status_Success}, nil
	}
	return &pb.DeleteUserResponse{Status: pb.Status_NothinToDo}, nil
}

func (s *userDB) GetUsers(req *pb.GetUsersRequest, stream pb.UserRepo_GetUsersServer) error {
	redisUsers, ok := s.redis.GetUsersFromRedis()
	if ok {
		log.Print("Retrieve users from Redis")
		for _, u := range *redisUsers {
			stream.Send(&pb.User{Id: u.Id, Name: u.Name, Email: u.Email})
		}
		return nil
	}

	log.Print("Retrieve users from DB")
	rows, err := s.db.db.Query("SELECT id, email, name FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	var users []redis.StoredUser
	for rows.Next() {
		u := &pb.User{}
		if err := rows.Scan(&u.Id, &u.Email, &u.Name); err != nil {
			return fmt.Errorf("Failed while scaning row: %s", err)
		}
		users = append(users, redis.StoredUser{Id: u.Id, Email: u.Email, Name: u.Name})
		stream.Send(u)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("Failed getting users: %s", err)
	}

	log.Print("Store to Redis")
	s.redis.StoreUsersIntoRedis(&users)

	return nil
}

func NewUserDB(db *DBManager, redis *redis.RedisManager, p *KafkaProducer) *userDB {
	s := &userDB{
		db:    db,
		redis: redis,
		p:     p,
		ctx:   context.Background(),
	}
	return s
}
