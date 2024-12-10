"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = require("vscode");
const node_1 = require("vscode-languageclient/node");
let client;
function activate(context) {
    // This line of code will only be executed once when your extension is activated
    // TODO: Start server exe and communicate with it
    let serverExe = context.extensionPath + '/out/ksl-lsp';
    let ServerOptions = {
        command: serverExe
    };
    let clientOptions = {
        // Register the server for KSL source files
        documentSelector: [
            {
                scheme: "file", language: "ksl"
            }
        ],
        synchronize: {
            fileEvents: vscode.workspace.createFileSystemWatcher("**/*.ksl"), // Watch files with your language's extension
        },
    };
    vscode.window.showInformationMessage("Starting language server");
    client = new node_1.LanguageClient("KSL", ServerOptions, clientOptions);
    client.info("Client created!");
    let t = client.start();
    t.then(_ => client.info("Connected!"), err => client.error(err));
}
function deactivate() {
    if (!client) {
        vscode.window.showWarningMessage("No active language server to stop.");
        return undefined;
    }
    vscode.window.showInformationMessage("Stopping language server.");
    return client.stop();
}
//# sourceMappingURL=extension.js.map