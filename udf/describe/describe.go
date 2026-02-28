package describe

func ViaMap[R any](dict map[string]R) func(string) (R, bool) {
	return func(key string) (R, bool) {
		value, ok := dict[key]
		return value, ok
	}
}
