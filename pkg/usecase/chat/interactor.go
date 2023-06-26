//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-example" mock_$GOPACKAGE/mock_$GOFILE
package chat

import (
	"context"

	"github.com/sivchari/chat-example/pkg/domain/entity"
	messagerepository "github.com/sivchari/chat-example/pkg/domain/repository/message"
	roomrepository "github.com/sivchari/chat-example/pkg/domain/repository/room"
	"github.com/sivchari/chat-example/pkg/ulid"
)

type Interactor interface {
	CreateRoom(ctx context.Context, name string) (*entity.Room, error)
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
	ListRoom(ctx context.Context) ([]*entity.Room, error)
	GetPass(ctx context.Context) (string, error)
	SendMessage(ctx context.Context, roomID, text string) error
	ListMessage(ctx context.Context, roomID string) ([]*entity.Message, error)
}

type interactor struct {
	ulidGenerator     ulid.ULIDGenerator
	roomRepository    roomrepository.Repository
	messageRepository messagerepository.Repository
}

func New(
	ulidGenerator ulid.ULIDGenerator,
	roomRepository roomrepository.Repository,
	messageRepository messagerepository.Repository,
) Interactor {
	return &interactor{
		ulidGenerator,
		roomRepository,
		messageRepository,
	}
}

func (i *interactor) CreateRoom(ctx context.Context, name string) (*entity.Room, error) {
	return nil, nil
}

func (i *interactor) ListRoom(ctx context.Context) ([]*entity.Room, error) {
	return nil, nil
}

func (i *interactor) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	return nil, nil
}

func (i *interactor) GetPass(_ context.Context) (string, error) {
	id, err := i.ulidGenerator.Generate()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (i *interactor) SendMessage(ctx context.Context, roomID, text string) error {
	return nil
}

func (i *interactor) ListMessage(ctx context.Context, roomID string) ([]*entity.Message, error) {
	return nil, nil
}
