package example_impl

import (
	"fmt"

	api "github.com/peterrosell/another-grpc-gw-example/example-api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DocumentServiceServer struct {
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

func (s *DocumentServiceServer) CreateDocument(ctx context.Context, req *api.CreateDocumentRequest) (*api.CreateDocumentResponse, error) {
	fmt.Printf("CreateDocument: %v\n", req.Document)
	fmt.Printf("CreateDocument: id=%s\n", req.Document.Id)

	response := new(api.CreateDocumentResponse)
	response.Document = req.Document

	return response, nil
}

func (s *DocumentServiceServer) ReadDocument(ctx context.Context, req *api.ReadDocumentRequest) (*api.ReadDocumentResponse, error) {
	fmt.Printf("ReadDocument: id=%s\n", req.Id)

	doc := new(api.Document)
	doc.Id = req.Id
	doc.DocDate = "2016-11-05T20:55:00Z"
	doc.Number = "2345"
	doc.Name = "Smörrebröd"
	response := new(api.ReadDocumentResponse)
	response.Document = doc

	return response, nil
}

func (s *DocumentServiceServer) DeleteDocument(ctx context.Context, req *api.DeleteDocumentRequest) (*api.DeleteDocumentResponse, error) {
	fmt.Printf("DeleteDocument: id=%s\n", req.Id)
	return new(api.DeleteDocumentResponse), nil
}

func (s *DocumentServiceServer) SearchDocument(ctx context.Context, req *api.SearchDocumentRequest) (*api.SearchDocumentResponse, error) {
	fmt.Printf("SearchDocument: %s -- %s\n", req.Filter.StartDate, req.Filter.EndDate)
	doc := new(api.Document)
	doc.Id = "e6d7c3f6-a391-11e6-9e64-ffd139010bad"
	doc.DocDate = "2016-11-05T20:55:00Z"
	doc.Number = "2345"
	doc.Name = "Smörrebröd"

	response := new(api.SearchDocumentResponse)
	response.Documents = []*api.Document{doc}

	return response, nil
}

func newDocumentServiceServer() *DocumentServiceServer {
	s := new(DocumentServiceServer)
	return s
}

func RegisterDocumentService(s *grpc.Server) {
	api.RegisterDocumentServiceServer(s, newDocumentServiceServer())
}
