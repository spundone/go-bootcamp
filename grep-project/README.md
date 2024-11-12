# Grep Implementation in Go

This project implements a clone of the Unix `grep` command in Go. It searches for patterns in text files and standard input.

## Features

- Basic pattern search in files (`./mygrep pattern file`)
- Case-insensitive search (`-i` flag)
- Recursive directory search (`-r` flag)
- Context lines (`-A`, `-B`, `-C` flags)
- Count matches only (`-c` flag)
- Output to file (`-o` flag)
- Standard input support
- Unicode support

## Usage

```bash
# Basic search in file
./mygrep "pattern" file.txt

# Case-insensitive search
./mygrep -i "pattern" file.txt

# Recursive search in directory
./mygrep -r "pattern" directory/

# Show 2 lines after match
./mygrep -A 2 "pattern" file.txt

# Show 2 lines before match
./mygrep -B 2 "pattern" file.txt

# Show 2 lines before and after match
./mygrep -C 2 "pattern" file.txt

# Count matches only
./mygrep -c "pattern" file.txt

# Output to file
./mygrep "pattern" file.txt -o output.txt

# Read from standard input
echo "some text" | ./mygrep "pattern"
```

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd grep-project

# Build the program
go build -o mygrep
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests verbosely
go test -v ./...
```

## Project Structure

```
grep-project/
├── main.go           # Main program
├── grep.go           # Core grep functionality
├── options.go        # Command line options
├── main_test.go      # Main program tests
├── grep_test.go      # Core functionality tests
├── testdata/        # Test files
│   ├── test1.txt
│   ├── test2.txt
│   └── nested/
│       └── test3.txt
├── go.mod           # Go module file
└── README.md        # This file
```

## Implementation Details

### Core Features

1. **Basic Pattern Search**
   - Exact string matching in files
   - Line-by-line processing
   - Error handling for file operations

2. **Case Insensitive Search**
   - `-i` flag enables case-insensitive matching
   - Unicode-aware case folding

3. **Recursive Search**
   - `-r` flag enables recursive directory traversal
   - Handles nested directories
   - Skips binary files

4. **Context Lines**
   - `-A n`: Shows n lines after match
   - `-B n`: Shows n lines before match
   - `-C n`: Shows n lines before and after match

5. **Count Mode**
   - `-c` flag shows only match count
   - Fast counting without storing matches

6. **Output Redirection**
   - `-o file` writes output to file
   - Error handling for file creation/writing

7. **Standard Input**
   - Reads from stdin when no file specified
   - Handles pipe operations

### Error Handling

- File not found
- Permission denied
- Directory instead of file
- Invalid flags
- Binary file detection
- Output file already exists

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.
```
