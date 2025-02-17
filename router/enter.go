package router

import (
	"awesomeProject/router/example"
	"awesomeProject/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
