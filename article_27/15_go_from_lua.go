// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Demonstrační příklad číslo 15:
//    Volání Go funkcí z Lua.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "go_from_lua.lua"

func Hello(L *lua.LState) int {
	log.Println("Hello from Go!")
	return 0
}

func main() {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	err := luaVM.DoFile(LuaSource)
	if err != nil {
		log.Fatal(err)
	}

	err = luaVM.CallByParam(lua.P{
		Fn:      luaVM.GetGlobal("call_go"),
		NRet:    0,
		Protect: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
