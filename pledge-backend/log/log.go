package log

import "go.uber.org/zap"

// Logs go to stdout/stderr so binaries and containers do not depend on source-tree paths.
var Logger = zap.NewNop()

func init() {
	if value, err := zap.NewProduction(); err == nil {
		Logger = value
	}
}
