"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.deactivate = exports.activate = void 0;
const vscode = require("vscode");
const node_1 = require("vscode-languageclient/node");
let client;
function activate(context) {
    // This line of code will only be executed once when your extension is activated
    // TODO: Start server exe and communicate with it
    let serverExe = '/home/naughtix/Projects/ksl-schema-language/internal/ksl-lsp/ksl-lsp';
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
exports.activate = activate;
function deactivate() {
    if (!client) {
        vscode.window.showWarningMessage("No active language server to stop.");
        return undefined;
    }
    vscode.window.showInformationMessage("Stopping language server.");
    return client.stop();
}
exports.deactivate = deactivate;
//# sourceMappingURL=extension.js.map