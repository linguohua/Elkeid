package plugin

import (
	"context"
	"fmt"

	"github.com/bytedance/Elkeid/agent/proto"
)

func (p *Plugin) Shutdown() {
	// TODO: implementation for windows
}

func Load(ctx context.Context, config proto.Config) (plg *Plugin, err error) {
	return nil, fmt.Errorf("plugin not yet support on windows")
}
