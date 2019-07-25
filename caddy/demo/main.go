package main

import (
	_ "github.com/BTBurke/caddy-jwt"
	"github.com/caddyserver/caddy/caddy/caddymain"
	_ "github.com/rodriguesdossantosvincent/loginsrv/caddy"
)

func main() {
	caddymain.Run()
}
