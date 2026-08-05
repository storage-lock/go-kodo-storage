package kodo_storage

// Kodo Storage 七牛云对象存储
// 七牛云 Kodo 兼容 S3 协议，可以复用 S3 实现

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// KodoStorage 七牛云 Kodo 对象存储实现
// Kodo 兼容 S3 协议，复用 S3 存储实现
type KodoStorage struct {
	*s3_storage.S3Storage
}

var _ storage.Storage = &KodoStorage{}

// NewKodoStorage 创建七牛云 Kodo 存储引擎
func NewKodoStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*KodoStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &KodoStorage{
		S3Storage: s3Storage,
	}, nil
}

const KodoStorageName = "kodo-storage"

func (x *KodoStorage) GetName() string {
	return KodoStorageName
}

// Kodo 与 S3 能力完全相同：
// - CAS: 通过条件 PUT 实现
// - ReliableTime: 通过外部 TimeProvider 满足
// - 不支持 AtomicDelete，释放走墓碑路径
func (x *KodoStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}