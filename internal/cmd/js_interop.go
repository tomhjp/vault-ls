//go:build js && wasm

package cmd

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

func registerGlobals(w *wasmService) {
	global := js.Global()
	global.Set("wasmProvideCompletionItems", js.FuncOf(w.completeJS))
	global.Set("wasmTextDocumentDidOpen", js.FuncOf(w.didOpenJS))
	global.Set("wasmTextDocumentDidChange", js.FuncOf(w.didChangeJS))
	global.Set("wasmTextDocumentDidClose", js.FuncOf(w.didCloseJS))
}

func (w *wasmService) completeJS(this js.Value, args []js.Value) interface{} {
	doc := args[0]
	position := args[1]

	line := monacoToLSCoord(position.Get("lineNumber").Int())
	character := monacoToLSCoord(position.Get("column").Int())

	items, err := w.complete(docFromJSValue(doc), lsp.Position{
		Line:      line,
		Character: character,
	})

	if err != nil {
		fmt.Println("Error generating completion list:", err)
	}

	// syscall/js requires a return value compatible with js.ValueOf()
	// Marshal and unmarshal to convert from Go struct -> map[string]interface{}
	marshalledItems, err := json.Marshal(items)
	if err != nil {
		fmt.Println("Error marshalling items", err)
		return nil
	}
	var itemsAsMap map[string]interface{}
	err = json.Unmarshal(marshalledItems, &itemsAsMap)
	if err != nil {
		fmt.Println("Error unmarshalling items", err)
		return nil
	}
	return itemsAsMap
}

func (w *wasmService) didOpenJS(this js.Value, args []js.Value) interface{} {
	doc := args[0]

	err := w.didOpen(docFromJSValue(doc))
	if err != nil {
		fmt.Println("Error opening model:", err)
	}
	return nil
}

func (w *wasmService) didChangeJS(this js.Value, args []js.Value) interface{} {
	uri := lsp.DocumentURI(args[0].String())
	version := int32(args[1].Int())
	jsChanges := args[2] // array of monaco.editor.IModelContentChange

	var changes []lsp.TextDocumentContentChangeEvent
	for i := 0; i < args[2].Length(); i++ {
		jsChange := jsChanges.Index(i)
		// rangeOffset := jsChange.Get("rangeOffset").Int()
		// if rangeOffset != 0 {
		// 	fmt.Println("WARNING - rangeOffset used but not supported:", rangeOffset)
		// }

		jsRange := jsChange.Get("range")

		changes = append(changes, lsp.TextDocumentContentChangeEvent{
			Range: &lsp.Range{
				Start: lsp.Position{
					Line:      monacoToLSCoord(jsRange.Get("startLineNumber").Int()),
					Character: monacoToLSCoord(jsRange.Get("startColumn").Int()),
				},
				End: lsp.Position{
					Line:      monacoToLSCoord(jsRange.Get("endLineNumber").Int()),
					Character: monacoToLSCoord(jsRange.Get("endColumn").Int()),
				},
			},
			RangeLength: uint32(jsChange.Get("rangeLength").Int()),
			Text:        jsChange.Get("text").String(),
		})
	}

	err := w.didChange(uri, version, changes)
	if err != nil {
		fmt.Println("Error updating model:", err)
	}

	return nil
}

func (w *wasmService) didCloseJS(this js.Value, args []js.Value) interface{} {
	uri := lsp.DocumentURI(args[0].String())

	err := w.didClose(uri)
	if err != nil {
		fmt.Println("Error closing model:", err)
	}

	return nil
}

func docFromJSValue(v js.Value) lsp.TextDocumentItem {
	return lsp.TextDocumentItem{
		URI:        lsp.DocumentURI(v.Get("uri").String()),
		LanguageID: v.Get("languageId").String(),
		Version:    int32(v.Get("version").Int()),
		Text:       v.Get("content").String(),
	}
}

// The minus 1 is pretty weird and I don't understand it. 🤷‍♂️
// If you auto-complete in the editor at the first character of the document,
// you get a position of (2,2), but vault-ls very sensibly considers that as (1,1)
// so we convert to the vault-ls coordinate system here.
// TODO: Do we need to do the reverse translation as well?
func monacoToLSCoord(v int) uint32 {
	return uint32(v - 1)
}
