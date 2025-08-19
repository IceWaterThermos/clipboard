# Multi-Clipboard Manager 🚀

A Windows clipboard manager that allows you to save and restore text content to/from up to 8 clipboard slots using keyboard shortcuts.

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org/)
[![Platform](https://img.shields.io/badge/Platform-Windows-green.svg)](https://www.microsoft.com/windows)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 📋 Key Features

- **8 Clipboard Slots**: Save with F1~F8 keys, load with 1~8 keys
- **Intuitive Key Combinations**: Simple `Ctrl + Shift` combinations
- **Real-time Monitoring**: Background keyboard input detection
- **Safe Memory Management**: Thread-safe clipboard operations
- **Visual Feedback**: Status display for save/load operations

## 🎯 Usage

### Basic Key Combinations

| Function | Key Combination | Description |
|----------|-----------------|-------------|
| **Save** | `Ctrl + Shift + F1~F8` | Save current clipboard content to slot |
| **Load** | `Ctrl + Shift + 1~8` | Load slot content to clipboard |
| **View Slots** | `Ctrl + Shift + F9` | Show all saved slot contents |
| **Exit** | `Ctrl + C` | Terminate program |

### Slot Layout

```
F1(0) F2(1) F3(2) F4(3) F5(4) F6(5) F7(6) F8(7)
 1(0)  2(1)  3(2)  4(3)  5(4)  6(5)  7(6)  8(7)
```

### Usage Example

1. **Save Text**
   - Copy desired text (`Ctrl + C`)
   - Press `Ctrl + Shift + F1` to save to slot 0

2. **Load Text**
   - Press `Ctrl + Shift + 1` to load slot 0 content to clipboard
   - Paste anywhere (`Ctrl + V`)

## 🔧 Installation and Build

### Requirements

- **Go 1.25 or higher**
- **Windows Operating System**
- **Administrator privileges** (required for system-wide keyboard monitoring)

### Build from Source

```bash
# Clone repository
git clone https://github.com/YOUR_USERNAME/clipboard-manager.git
cd clipboard-manager

# Install dependencies
go mod tidy

# Format code
go fmt ./...

# Build
go build -o clipboard-manager.exe

# Or run directly
go run clipboard_manager.go
```

### Optimized Build

For performance-optimized build:

```bash
# Optimized release build
go build -ldflags="-s -w" -o clipboard-manager.exe

# Size optimization (UPX compression - optional)
# After installing UPX: upx --best clipboard-manager.exe
```

### Build Flags Explanation

- `-ldflags="-s -w"`: Remove debug information for smaller file size
  - `-s`: Remove symbol table
  - `-w`: Remove DWARF debug information

## 🚀 Running the Application

### Run with Administrator Privileges

Administrator privileges are required for system-wide keyboard monitoring:

1. **Run Command Prompt as Administrator**
2. Navigate to program directory
3. Execute: `clipboard-manager.exe`

Or

1. **Right-click the executable**
2. **Select "Run as administrator"**

### Auto-start on Windows Boot (Optional)

To automatically run on Windows startup:

1. Press `Win + R` → type `shell:startup`
2. Copy a shortcut of the executable to the startup folder
3. Set the shortcut properties to "Run as administrator"

## 🔍 Troubleshooting

### Common Issues

| Problem | Cause | Solution |
|---------|-------|----------|
| Key combinations not working | Insufficient privileges | Run as administrator |
| Clipboard read failure | Another program using clipboard | Try again after a moment |
| Program won't start | Missing Go runtime | Verify Go installation |

### Debugging

Check these items when issues occur:

```bash
# Check Go version
go version

# Verify dependencies
go mod verify

# Test build
go build -v
```

## 🏗️ Project Structure

```
clipboard-manager/
├── clipboard_manager.go   # Main source code
├── go.mod                # Go module configuration
├── go.sum                # Dependency checksums
├── README.md             # Korean documentation
├── README_EN.md          # English documentation (this file)
└── CLAUDE.md             # Developer guide
```

## 📦 Dependencies

- `github.com/atotto/clipboard`: Cross-platform clipboard access
- `github.com/go-vgo/robotgo`: System automation (indirect dependency)

## ⚠️ Important Notes

- **Windows Only**: Currently uses Windows API directly, not compatible with other OS
- **Administrator Required**: Essential for system-wide keyboard monitoring
- **Memory Security**: Clipboard content is stored in memory as plain text
- **Unicode Support**: Full support for Unicode text including Korean

## 🔄 Version History

- **v1.0.0**: Initial Release
  - 8-slot clipboard management
  - Windows API-based keyboard monitoring
  - Korean UI support

## 📄 License

This project is distributed under the MIT License. See [LICENSE](LICENSE) file for details.


---

📖 **한국어 문서**: [README.md](README.md)