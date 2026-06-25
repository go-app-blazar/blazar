package matchmedia

// Debug toggles debug logging for this package.
func Debug(debug bool) {
	debugMatchMedia = debug
}

// debugMatchMedia toggles debug logging for MatchMedia.
var debugMatchMedia bool = false

// DebugMatchMedia toggles debug logging for MatchMedia.
func DebugMatchMedia(debug bool) {
	debugMatchMedia = debug
}
