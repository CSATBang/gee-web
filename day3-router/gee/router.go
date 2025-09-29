package gee

import "strings"

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg.roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

}
