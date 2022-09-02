package xstrings

import "strings"

// TrimPrefixs 去除字符串开头
func TrimPrefixs(s string, prefixs ...string) string {
	for _, prefix := range prefixs {
		s = strings.TrimPrefix(s, prefix)
	}
	return s
}

// TrimSuffixs 去除字符串结尾
func TrimSuffixs(s string, suffixs ...string) string {
	for _, suffix := range suffixs {
		s = strings.TrimSuffix(s, suffix)
	}
	return s
}

// ContainssOr 一个存在即可
func ContainssOr(s string, substr ...string) bool {
	for _, sub := range substr {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// ContainssAnd 必须都存在
func ContainssAnd(s string, substr ...string) bool {
	for _, sub := range substr {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// TrimSpacAndToLower 去除空格和转为小写
func TrimSpacAndToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// TrimSpacAndToTitle 去除空格和转为大写
func TrimSpacAndToTitle(s string) string {
	return strings.ToTitle(strings.TrimSpace(s))
}
