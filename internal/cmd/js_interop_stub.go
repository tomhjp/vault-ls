//go:build !(js && wasm)

package cmd

func registerGlobals(_ *wasmService) {
	panic("not supported")
}
