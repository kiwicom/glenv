package glenv

import "github.com/kiwicom/glenv/internal/glenv/log"

// If you run any command with '--debug', the glenv start writing all logs
// into stderr.
//
// I've decided to write own small 'log' package because I don't need all
// functionality of heavy-loggers and I rather choose not to use dependencies
// until they're really needed.
// Anyway, if in nearest future I will need logger library, this package will
// be like interface/abstraction for any logger implementation.
func EnableDebug() {
	log.SetSeverity(5)
}
