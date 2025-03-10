package utils

import "github.com/coding/src/domain"

func RemoveNode(s []*domain.Node, i int) []*domain.Node {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
