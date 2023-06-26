package xoss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/conf"
)

var (
	log                       = logger.Logger("xoss")
	ERR_OSS_INSTANCE_IS_EMPTY = errors.New("oss instance is empty")
)

var (
	cli *OssClient
)

type OssClient struct {
	oss *oss.Bucket
	cfg conf.Oss
}

func NewOssClient(cfg conf.Oss) *OssClient {
	cli = &OssClient{cfg: cfg}
	cli.oss, _ = ConnectOss(cfg)
	return cli
}

func GetOss() *oss.Bucket {
	if cli.oss == nil {
		cli.oss, _ = ConnectOss(cli.cfg)
	}
	return cli.oss
}

func ConnectOss(cfg conf.Oss) (*oss.Bucket, error) {

	ossClient, err := oss.New(cfg.EndPoint, cfg.AccessKey, cfg.SecretKey)
	if err != nil {
		return nil, err
	}
	bucket, err := ossClient.Bucket(cfg.Bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
