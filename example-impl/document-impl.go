package myapi_impl

import (
	"fmt"

	"github.com/peterrosell/grpc-gw-example/api"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/peterrosell/another-gprc-gw-example/example-api"
)


type biServer struct {
}



/*
func (s *DocumentServiceServer) FetchStatistics(ctx context.Context, req *my.FilterRequest) (*my.BiResponse, error) {
	fmt.Printf("FetchStatistics: %s -- %s\n", req.Sub.StartDate, req.Sub.EndDate)
	a := new(my.Article)
	a.Id = "e6d7c3f6-a391-11e6-9e64-ffd139010bad"
	a.DocDate = "2016-11-05T20:55:00Z"
	a.Number = "2345"
	a.Name = "Smörrebröd"

	response := new(my.BiResponse)
	response.Articles = []*my.Article{a}

	return response, nil
}
*/

func (s *DocumentServiceServer) CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	return nil,nil
}

func (s *DocumentServiceServer) ReadDocument(context.Context, *ReadDocumentRequest) (*ReadDocumentResponse, error) {
	return nil,nil
}

func (s *DocumentServiceServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil,nil
}

func (s *DocumentServiceServer) SearchDocument(context.Context, *SearchDocumentRequest) (*SearchDocumentResponse, error) {
	return nil,nil
}

func newDocumentServiceServer() *DocumentServiceServer {
	s := new(biServer)
	return s
}

func RegisterDocumentService(s *grpc.Server) {
	RegisterDocumentServiceServer(s, newBiServer())
}
