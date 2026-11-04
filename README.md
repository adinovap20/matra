# Sanskrit Prosody Analyzer 📚

A tool to analyze Sanskrit verses, identifying aksharas, matras, and traditional poetic meters (vruttas) in both Sanskrit and Marathi traditions.

## Features ✨

- **CLI**: Command-line interface for quick analysis of text files.
- **Web UI**: Interactive web application for real-time verse analysis.
- **Wasm**: WebAssembly module for client‑side processing.
- **Meter identification**: Recognizes common Sanskrit and Marathi meters (e.g., Indravajrā, Bhujangaprayāta, Vidyunmālā).

## How to Use

### CLI

```bash
go run ./cmd/cli
# or build a binary:
go build ./cmd/cli
./prosody-analyzer -i input.txt -o output.json
```

The CLI reads an input file (default `input.txt`), analyzes each line, and writes a JSON report to the output file (default `output.json`).

### Web UI

Open `web/index.html` in a browser. Paste or type a Sanskrit verse, click **Analyze**, and view the breakdown, aksharas, matras, and identified meter.

### Wasm (optional)

The project includes a WebAssembly build (`wasm/matra.wasm`) for client‑side analysis without a server.

## Project Structure

```
.
├── cmd/
│   ├── cli/          # Command‑line tool
│   └── wasm/         # WebAssembly bindings
├── prosody/          # Core analysis logic (Go)
├── web/              # Front‑end assets
├── go.mod
└── README.md
```

## Contributing

Feel free to open issues or pull requests. Ensure code follows Go conventions and the existing style.

## License

[MIT](LICENSE)