package impl

import (
	"github.com/souxls/dvs_api/internal/app/bll"
	"github.com/souxls/dvs_api/internal/app/bll/impl/internal"
	"go.uber.org/dig"
)

// Inject 注入bll实现
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo IDemo) {
//   })
func Inject(container *dig.Container) error {
	_ = container.Provide(internal.NewDemo)
	_ = container.Provide(func(b *internal.Demo) bll.IDemo { return b })
	return nil
}
