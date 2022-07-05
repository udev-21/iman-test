package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/udev-21/iman-test-post-crud-svc/pkg/db"
	"github.com/udev-21/iman-test-post-crud-svc/pkg/models"
	"github.com/udev-21/iman-test-post-crud-svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) Create(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	if in.Post == nil {
		return nil, errors.New("data required")
	}

	newID := in.Post.ID
	if in.Post.ID < 1 {
		lastPost := models.Post{}
		req := s.H.DB.Last(&lastPost)
		if req.Error != nil && !errors.As(req.Error, &sql.ErrNoRows) {
			return nil, errors.New("something went wrong")
		}
		newID = lastPost.ID + 1
	}

	newPost := models.Post{
		ID:     newID,
		UserID: in.Post.UserID,
		Title:  in.Post.Title,
		Body:   in.Post.Body,
	}
	result := s.H.DB.Create(&newPost)
	createPostResponse := &pb.CreatePostResponse{
		ID: newID,
	}

	if result.Error != nil {
		createPostResponse.Status = pb.Response_statusErr
		createPostResponse.ErrorMsg = new(string)
		*createPostResponse.ErrorMsg = "something went wrong while saving to db"
	}
	return createPostResponse, result.Error
}

func (s *Server) Update(ctx context.Context, in *pb.UpdatePostRequest) (*pb.RequestResponse, error) {
	postShouldUpdate := models.Post{
		ID: in.Post.ID,
	}

	res := s.H.DB.First(&postShouldUpdate)
	if res.Error != nil {
		return nil, res.Error
	}
	log.Println("found:", postShouldUpdate.Title)

	postShouldUpdate.Body = in.Post.Body
	postShouldUpdate.Title = in.Post.Title
	postShouldUpdate.UserID = in.Post.UserID
	res = s.H.DB.Save(&postShouldUpdate)
	if res.Error != nil {
		return nil, res.Error
	}
	log.Println("updated:", postShouldUpdate)

	return &pb.RequestResponse{
		Status: pb.Response_statusOK,
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeletePostRequest) (*pb.RequestResponse, error) {
	post := models.Post{ID: in.ID}
	res := s.H.DB.Delete(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pb.RequestResponse{
		Status: pb.Response_statusOK,
	}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {
	post := models.Post{ID: in.ID}
	res := s.H.DB.First(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pb.ReadPostResponse{
		Status: pb.Response_statusOK,
		Post: &pb.Post{
			ID:     post.ID,
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
		},
	}, nil
}

func (s *Server) List(ctx context.Context, in *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	posts := []models.Post{}
	query := s.H.DB.Order("id ASC")
	if in.UserID != nil && *in.UserID > 0 {
		query.Where("user_id = ?", *in.UserID)
	}
	query.Limit(20)
	if in.Page != nil && *in.Page > 0 {
		query.Offset(20 * (int(*in.Page) - 1))
	}
	res := query.Find(&posts)
	if res.Error != nil {
		return nil, res.Error
	}

	var pbPosts []*pb.Post
	for _, post := range posts {
		pbPosts = append(pbPosts, &pb.Post{
			ID:     post.ID,
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
		})
	}
	return &pb.ListPostResponse{
		Status: pb.Response_statusOK,
		Posts:  pbPosts,
	}, nil
}
