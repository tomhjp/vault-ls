import * as monaco from 'monaco-editor'
import { hclLanguage } from "./monarch"
// import { vsConverted, vsDarkConverted } from "./themes"
// import { loadWASM } from 'onigasm'
// import { Registry } from 'monaco-textmate'
// import { wireTmGrammars } from 'monaco-editor-textmate'

const LANGUAGE_ID = 'vault-agent';
const MODEL_URI = 'file:///model.vault-agent'
const MONACO_URI = monaco.Uri.parse(MODEL_URI);

const initialContents = `// This is a Vault Agent config file
// See documentation here: https://www.vaultproject.io/docs/agent#configuration
// TIP: press Ctrl+Space while editing for autocomplete options

pid_file = "./pidfile"

`;

loadGolangWASM("vault-ls.wasm").then(async () => {
    await clientMain()
}).catch((error: any) => {
    console.log("ERROR", error)
})

// async function loadGrammars(editor: monaco.editor.IStandaloneCodeEditor) {
//     await loadWASM(require('onigasm/lib/onigasm.wasm'))

//     const registry = new Registry({
//         getGrammarDefinition: async (scopeName) => {
//             switch (scopeName) {
//                 case "source.hcl":
//                     return {
//                         format: 'json',
//                         content: await (await fetch(`hcl.tmGrammar.json`)).text()
//                     }
//                 case "source.hcl.vault":
//                     return {
//                         format: 'json',
//                         content: await (await fetch(`vault.tmGrammar.json`)).text()
//                     }
//             }
//         }
//     })

//     // map of monaco language IDs to TextMate scopeNames
//     const grammars = new Map()
//     grammars.set('hcl', 'source.hcl')
//     grammars.set(LANGUAGE_ID, 'source.hcl.vault')
    
//     await wireTmGrammars(monaco, registry, grammars, editor)
// }

function createDocument(model: monaco.editor.IReadOnlyModel) {
    return {
        uri: MODEL_URI,
        languageId: model.getLanguageId(),
        version: model.getVersionId(),
        content: model.getValue(),
    }
}

function getModel(): monaco.editor.IModel {
    return monaco.editor.getModel(MONACO_URI) as monaco.editor.IModel;
}

async function clientMain() {
    // Register the language with Monaco.
    monaco.languages.register({
        id: LANGUAGE_ID,
        extensions: ['.vault-agent'],
    });

    monaco.languages.setMonarchTokensProvider(LANGUAGE_ID, hclLanguage);

    // // monaco's built-in themes aren't powerful enough to handle TM tokens
    // // https://github.com/Nishkalkashyap/monaco-vscode-textmate-theme-converter#monaco-vscode-textmate-theme-converter
    // monaco.editor.defineTheme('vs-converted', vsConverted as monaco.editor.IStandaloneThemeData);
    // monaco.editor.defineTheme('vs-dark-converted', vsDarkConverted as monaco.editor.IStandaloneThemeData);

    // Create the editor.
    monaco.editor.create(document.getElementById("container")!, {
        model: monaco.editor.createModel(initialContents, LANGUAGE_ID, MONACO_URI),
        glyphMargin: true,
        lightbulb: {
            enabled: true
        },
        theme: 'vs',
    });

    // await loadGrammars(editor);

    var select = document.getElementById('themeselect') as HTMLSelectElement;
	var currentTheme = 'vs';
	select.onchange = function () {
		currentTheme = select.options[select.selectedIndex].value;
		monaco.editor.setTheme(currentTheme);
	};

    // Document syncing calls.
    wasmTextDocumentDidOpen(createDocument(getModel()));
    getModel().onDidChangeContent((event) => {
        wasmTextDocumentDidChange(MODEL_URI, event.versionId, event.changes);
    });

    // Completion plumbing.
    monaco.languages.registerCompletionItemProvider(LANGUAGE_ID, {
        provideCompletionItems(model, position, context, token): monaco.languages.CompletionList {
            let result = wasmProvideCompletionItems(createDocument(model), position);
            for (var item of result.items) {
                if (item.command) {
                    // vault-ls -> monaco conversion
                    item.command.id = item.command.command;
                }
            }
            return {
                suggestions: result.items,
                incomplete: result.isIncomplete,
            };
        },

        resolveCompletionItem(item, token): monaco.languages.CompletionItem | monaco.Thenable<monaco.languages.CompletionItem> {
            // vault-ls handles this in the `insertTextFormat` field, but we seem to have a mismatch on the client
            // side and monaco wants the equivalent information in the `insertTextRules` field instead.
            // TODO: Either get to the bottom of that, or transform it in the original provideCompletionItems call instead.
            item.insertTextRules = monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet;
            return item;
        }
    });
}
