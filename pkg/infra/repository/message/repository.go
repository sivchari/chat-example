package message

import (
	"context"
	"sync"

	"github.com/sivchari/chat-example/pkg/domain/entity"
	"github.com/sivchari/chat-example/pkg/domain/repository/message"
)

type repository struct {
	mapByRoomID map[string][]*entity.Message
	// TODO: sync.RWMutexとの違いを考えて最適化しよう
	mu sync.Mutex
}

func New() message.Repository {
	return &repository{
		mapByRoomID: make(map[string][]*entity.Message, 0),
	}
}

func (r *repository) Insert(_ context.Context, message *entity.Message) error {
	return nil
}

func (r *repository) SelectByRoomID(_ context.Context, roomId string) ([]*entity.Message, error) {
	return nil, nil
}
