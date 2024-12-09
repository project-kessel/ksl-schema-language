package main

import (
	"strings"
	"unicode"

	"github.com/project-kessel/ksl-schema-language/pkg/ksl"
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

type SourceFile struct {
	Text string
}

var files = map[string]*SourceFile{}

func NewSourceFile(text string) *SourceFile {
	return &SourceFile{
		Text: text,
	}
}

func main() {
	// This increases logging verbosity (optional)
	commonlog.Configure(1, nil)
	_ = files
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textDocumentCompletion,
		TextDocumentDidOpen:    textDocumentDidOpen,
		TextDocumentDidChange:  textDocumentDidChange,
		TextDocumentDidClose:   textDocumentDidClose,
	}

	server := server.NewServer(&handler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider.TriggerCharacters = append(capabilities.CompletionProvider.TriggerCharacters, " ")

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
	file, ok := files[params.TextDocument.URI]
	if !ok {
		return nil, nil
	}

	index := params.Position.IndexIn(file.Text)
	fragment := file.Text[:index]
	fragment = strings.TrimRightFunc(fragment, unicode.IsLetter)
	tokens, rule, err := ksl.GetNextTokens(strings.NewReader(fragment))
	if err != nil {
		return nil, err
	}

	completionItems := []protocol.CompletionItem{}

	for _, token := range tokens {
		if token.Literal != nil {
			completionItems = append(completionItems, literalCompletion(*token.Literal))
		} else {
			switch token.Name {
			case "ACCESS":
				for _, kw := range []string{"public", "internal"} {
					completionItems = append(completionItems, literalCompletion(kw))
				}
				if rule == "relation" {
					completionItems = append(completionItems, literalCompletion("private"))
				}
			case "CARDINALITY":
				for _, kw := range []string{"AtMostOne", "ExactlyOne", "AtLeastOne", "Any"} {
					completionItems = append(completionItems, literalCompletion(kw))
				}
			}
		}
	}

	return completionItems, nil
}

func literalCompletion(name string) protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword

	return protocol.CompletionItem{
		Label:            name,
		Kind:             &kind,
		Detail:           &name,
		InsertText:       &name,
		CommitCharacters: []string{" "},
	}
}

func textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	uri := params.TextDocument.URI
	sourceFile := SourceFile{
		Text: params.TextDocument.Text,
	}
	files[uri] = &sourceFile
	return nil
}

func textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	sourceFile, ok := Get(params.TextDocument.URI)
	if !ok {
		return nil
	}

	sourceFile.ApplyChanges(params.ContentChanges)

	return nil
}

func textDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	Close(params.TextDocument.URI)
	return nil
}

func (s *SourceFile) ApplyChanges(changes []interface{}) {
	for _, change := range changes {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			startIndex, endIndex := c.Range.IndexesIn(s.Text)
			s.Text = s.Text[:startIndex] + c.Text + s.Text[endIndex:]
		case protocol.TextDocumentContentChangeEventWhole:
			s.Text = c.Text
		}
	}
}

func Get(URI string) (*SourceFile, bool) {
	f, ok := files[URI]
	return f, ok
}

func Close(URI string) {
	delete(files, URI)
}
