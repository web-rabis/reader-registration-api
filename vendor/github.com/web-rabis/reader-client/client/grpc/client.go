package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-client/protobuf"
)

type BaseClient struct {
	address  string
	conn     *grpc.ClientConn
	dialOpts []grpc.DialOption

	readerSvc     client.ReaderService
	dictionarySvc client.DictionaryService
}

var _ client.Base = &BaseClient{}

func NewClient(config *model.ConnectionConfig) (client.Base, error) {
	var grpcOpts []grpc.DialOption
	if config.Insecure == true {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &BaseClient{
		address:  config.Address,
		dialOpts: grpcOpts,
	}, nil
}

func (c *BaseClient) Connect() (err error) {
	c.conn, err = grpc.Dial(c.address, c.dialOpts...)
	if err != nil {
		return err
	}

	return err
}

func (c *BaseClient) Close() error {
	return c.conn.Close()
}

func (c *BaseClient) ReaderSvc() client.ReaderService {
	if c.readerSvc == nil {
		c.readerSvc = NewReaderServiceClient(protobuf.NewReaderSvcClient(c.conn))
	}

	return c.readerSvc
}

func (c *BaseClient) DictionarySvc() client.DictionaryService {
	if c.dictionarySvc == nil {
		c.dictionarySvc = NewDictionaryServiceClient(protobuf.NewDictionarySvcClient(c.conn))
	}

	return c.dictionarySvc
}
