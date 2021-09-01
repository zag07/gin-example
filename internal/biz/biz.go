// biz 应该依赖于接口，而不是具体实现
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBlogUseCase)
