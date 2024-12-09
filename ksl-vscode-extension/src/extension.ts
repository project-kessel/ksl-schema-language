import * as vscode from 'vscode'

import {
    LanguageClient,
    LanguageClientOptions,
    ServerOptions,
  } from 'vscode-languageclient/node';

let client: LanguageClient;

export function activate(context: vscode.ExtensionContext) {

    // This line of code will only be executed once when your extension is activated
    // TODO: Start server exe and communicate with it
    let serverExe = 'ksl-lsp';
    let ServerOptions: ServerOptions = {
        command: serverExe
    }

    let clientOptions: LanguageClientOptions = {
        // Register the server for KSL source files
        documentSelector: [
            {
                scheme: "file", language: "ksl"
            }
        ],
        synchronize: {
            fileEvents: vscode.workspace.createFileSystemWatcher("**/*.ksl"), // Watch files with your language's extension
        },

    }

    vscode.window.showInformationMessage("Starting language server")
    client = new LanguageClient("KSL", ServerOptions, clientOptions);
    client.info("Client created!")
    let t = client.start();
    t.then(_ => client.info("Connected!"), err => client.error(err))
}

export function deactivate(): Thenable<void> | undefined {
    if (!client) {
      vscode.window.showWarningMessage("No active language server to stop.")
      return undefined;
    }
    vscode.window.showInformationMessage("Stopping language server.")
    return client.stop();
  }