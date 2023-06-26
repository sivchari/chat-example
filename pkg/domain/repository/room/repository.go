//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-example" mock_$GOPACKAGE/mock_$GOFILE
package room

import (
	"context"

	"github.com/sivchari/chat-example/pkg/domain/entity"
)

type Repository interface {
	Insert(ctx context.Context, room *entity.Room) error
	Select(ctx context.Context, id string) (*entity.Room, error)
	SelectAll(ctx context.Context) ([]*entity.Room, error)
}
