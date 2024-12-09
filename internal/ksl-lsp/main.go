package main

import (
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	// Must include a backend implementation
	// See CommonLog for other options: https://github.com/tliron/commonlog
	_ "github.com/tliron/commonlog/simple"
)

const lsName = "KSL"

var (
	version string = "0.0.1"
	handler protocol.Handler
)

func main() {
	// This increases logging verbosity (optional)
	commonlog.Configure(1, nil)

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textDocumentCompletion,
	}

	server := server.NewServer(&handler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	commonlog.NewInfoMessage(0, "Initialized!")
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	completionItems := []protocol.CompletionItem{
		keywordCompletion("version"),
		keywordCompletion("namespace"),
		keywordCompletion("public"),
		keywordCompletion("internal"),
		keywordCompletion("private"),
		keywordCompletion("type"),
		keywordCompletion("relation"),
		keywordCompletion("import"),
		keywordCompletion("extension"),
		keywordCompletion("AtMostOne"),
		keywordCompletion("ExactlyOne"),
		keywordCompletion("AtLeastOne"),
		keywordCompletion("Any"),
		keywordCompletion("as"),
		keywordCompletion("and"),
		keywordCompletion("or"),
		keywordCompletion("unless"),
		keywordCompletion("ALLOW_DUPLICATES"),
	}

	return completionItems, nil
}

func keywordCompletion(name string) protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	return protocol.CompletionItem{
		Label:            name,
		Kind:             &kind,
		Detail:           &name,
		InsertText:       &name,
		CommitCharacters: []string{" "},
	}
}
