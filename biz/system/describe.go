package system

func DescribeByRuleSet(getter func(key string) string) func(string) string {
	return func(s string) string {
		return getter(s)
	}
}
