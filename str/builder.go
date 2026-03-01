package str

import "strings"

type Builder struct {
	items []string
}

// NewBuilder creates a new string Builder
func NewBuilder() *Builder {
	return &Builder{items: make([]string, 0)}
}

// Add adds a string to the string Builder
func (b *Builder) Add(item string) {
	b.items = append(b.items, item)
}

// Build builds the full string parts, joined by the separator
func (b *Builder) Build(separator string) string {
	return strings.Join(b.items, separator)
}
