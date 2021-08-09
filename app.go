package gin_example

import (
	"io"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var AppSet = wire.Struct(new(App), "*")

type App struct {
	Log          *zap.Logger
	TracerCloser io.Closer
	DB           *gorm.DB
}
