
# Building an LSP

- LSP (usually) communicates over stdin and stdout but you can use anything

 The protocol are messages that look like:
```
Content-Length: ...\r\n
\r\n
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "textDocument/completion",
  "params": {
    ...
  }
}
```
- Many editors support stdin and stdout and TCP
