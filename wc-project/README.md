# Word Count (wc) Implementation in Go

This project implements a clone of the Unix `wc` command in Go. It counts lines, words, and characters in text files.

## Features

- Count lines (`-l` flag)
- Count words (`-w` flag)
- Count characters (`-c` flag)
- Support for multiple files
- Handles standard input when no file is specified
- Concurrent processing of multiple files
- Unicode support

## Installation 

To install the wc command, run:

```bash
git clone https://github.com/spundone/go-bootcamp.git
cd wc-project
go build
```

## Usage
Count lines, words, and characters (default)
```bash
./wc file.txt
```

Count only lines
```bash
./wc -l file.txt
```

Count only words
```bash
./wc -w file.txt
```

Count only characters
```bash
./wc -c file.txt
Multiple files
```bash
./wc file1.txt file2.txt
```

Read from standard input
```bash
cat file.txt | ./wc
```
## Examples
Count all metrics in a file
```bash
./wc -l shakespeare.txt
```
Count only lines
```bash
./wc -l example.txt
3 example.txt
```
Multiple files
```bash
./wc .txt
3 6 18 file1.txt
5 10 30 file2.txt
8 16 48 total
```

## Testing
Run all tests
```bash
go test
```
Run tests with verbose output
```bash
go test -v
```
Run tests with coverage
```bash
go test -cover
```

## Implementation Details

- Uses Go's `bufio.Scanner` for efficient file reading
- Implements concurrent file processing using goroutines
- Handles Unicode characters correctly
- Follows Unix `wc` behavior for counting words and lines
- Provides clear error messages for common issues

## Error Handling

The program handles various error conditions:
- Non-existent files
- Permission denied
- Directory instead of file
- Invalid flags
- IO errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by the Unix `wc` command
- Shakespeare texts from [shakespeare-plays-dataset-scraper](https://github.com/ravexina/shakespeare-plays-dataset-scraper)

### Shakespeare Tests

The project includes tests using Shakespeare's plays as test data. To enable these tests:

1. Create a directory for Shakespeare's texts:
```bash
mkdir shakespeare-db
```

2. Download some Shakespeare plays:
```bash
cd shakespeare-db
curl -O https://raw.githubusercontent.com/ravexina/shakespeare-plays-dataset-scraper/master/shakespeare-db/hamlet.txt
curl -O https://raw.githubusercontent.com/ravexina/shakespeare-plays-dataset-scraper/master/shakespeare-db/macbeth.txt
curl -O https://raw.githubusercontent.com/ravexina/shakespeare-plays-dataset-scraper/master/shakespeare-db/romeo-and-juliet.txt
```

## Project Structure

```
wc-project/
├── main.go           # Main program implementation
├── main_test.go      # Basic functionality tests
├── shakespeare_test.go # Shakespeare corpus tests
├── go.mod           # Go module file
├── README.md        # This file
└── shakespeare-db/  # Directory for Shakespeare texts
    ├── hamlet.txt
    ├── macbeth.txt
    └── romeo-and-juliet.txt
```

## Implementation Details

- Uses Go's `bufio.Scanner` for efficient file reading
- Implements concurrent file processing using goroutines
- Handles Unicode characters correctly
- Follows Unix `wc` behavior for counting words and lines
- Provides clear error messages for common issues

## Error Handling

The program handles various error conditions:
- Non-existent files
- Permission denied
- Directory instead of file
- Invalid flags
- IO errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by the Unix `wc` command
- One2N [Go Bootcamp](https://one2n.io/go-bootcamp/go-projects/word-count-in-go)
- Shakespeare texts from [shakespeare-plays-dataset-scraper](https://github.com/ravexina/shakespeare-plays-dataset-scraper)
