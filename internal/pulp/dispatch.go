package pulp

import (
	"context"
	"fmt"
	"strings"
)

func HandleDispatch(ctx context.Context, endpoint string, payload []byte) error {
	if strings.TrimSpace(endpoint) == "" {
		return fmt.Errorf("empty Pulp endpoint")
	}
	// 把入口 ctx 的取消传到出站读循环：值长撤令即在此打断，不再堵到 HTTP 超时。
	return PostOutbound(ctx, endpoint, payload)
}
