package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(ParsePattern("/p/:name/c"), []string{"p", ":name"})
	ok = reflect.DeepEqual(ParsePattern("/p/*"), []string{"p", "*"})
	ok = reflect.DeepEqual(ParsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/b/c")

	fmt.Printf("%#v\n", r.roots["GET"])
	fmt.Printf("matched path: %s, params:%v\n", n.pattern, ps)
}
