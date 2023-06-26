//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-example" mock_$GOPACKAGE/mock_$GOFILE
package message

import (
	"context"

	"github.com/sivchari/chat-example/pkg/domain/entity"
)

type Repository interface {
	Insert(ctx context.Context, message *entity.Message) error
	SelectByRoomID(ctx context.Context, roomID string) ([]*entity.Message, error)
}
