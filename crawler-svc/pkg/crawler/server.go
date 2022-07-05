package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	myPb "github.com/udev-21/iman-test-crawler-svc/pkg/crawler/pb"
	"github.com/udev-21/iman-test-crawler-svc/pkg/models"
	"github.com/udev-21/iman-test-crawler-svc/pkg/post-crud/pb"
)

type Server struct {
	PostCRUD pb.PostCRUDServiceClient
}

type wrapper struct {
	sync.Mutex
	faliedPages []int
	donePages   []int
	pages       [][]models.Post
}

type MyStruct struct {
	Posts []models.Post `json:"data"`
}

var container = new(wrapper)

func (s *Server) Start(ctx context.Context, req *myPb.StartRequest) (*myPb.StartResponse, error) {

	go func() {

		for i := 1; i <= 50; i++ {
			res, err := myHttpClient.Get(fmt.Sprintf("https://gorest.co.in/public/v1/posts?page=%d", i))
			if err != nil {
				log.Fatal(err)
			}
			data := MyStruct{}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				container.Lock()
				container.faliedPages = append(container.faliedPages, i)
				container.Unlock()
				log.Fatal(err)
			}
			if len(data.Posts) == 0 {
				return
			}
		outer:
			for _, post := range data.Posts {
				_, err := s.PostCRUD.Create(context.Background(), &pb.CreatePostRequest{
					Post: &pb.Post{
						ID:     post.ID,
						UserID: post.UserID,
						Title:  post.Title,
						Body:   post.Title,
					},
				})
				if err != nil {
					container.Lock()
					container.faliedPages = append(container.faliedPages, i)
					container.Unlock()
					log.Println(err)
					continue outer
				}
			}
			container.Lock()
			container.donePages = append(container.donePages, i)
			container.pages = append(container.pages, data.Posts)
			container.Unlock()

		}

	}()

	return &myPb.StartResponse{
		Data: myPb.StartResponse_statusOK,
	}, nil
}

func (s *Server) Status(context.Context, *myPb.StatusRequest) (*myPb.StatusResponse, error) {
	container.Lock()
	doneCnt := container.donePages
	failCnt := container.faliedPages
	container.Unlock()

	return &myPb.StatusResponse{
		Done:   int32(len(doneCnt)),
		Failed: int32(len(failCnt)),
	}, nil
}

var myHttpClient = http.Client{Timeout: 3 * time.Second}
