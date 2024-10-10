# Futil

`Futil` is a simple file utility command-line application written in Go. It provides basic file operations such as counting the number of lines in a file and calculating the checksum of files using various algorithms (MD5, SHA1, SHA256). The tool supports reading from both files and standard input.

## How to Build and Run

### Prerequisites

- **Go**: Ensure that Go (1.19 or later) is installed on your machine. You can download it from [here](https://golang.org/dl/).
- **Make**: Ensure that `make` is installed. On Linux and macOS, it is usually pre-installed. For Windows users, you can install `make` via [GNUWin](http://gnuwin32.sourceforge.net/packages/make.htm) or by using a Windows Subsystem for Linux (WSL).

### Build the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/Largeb0525/futil.git
   ```

2. Build the application using make:
   ```bash
   make install
   ```

## Features

1. **Line Count**
   - Count the number of lines in a specified file or standard input.
   - Command:
     ```bash
     futil linecount -f <filename>
     ```
     Example:
     ```bash
     futil linecount -f myfile.txt
     ```

2. **Checksum Calculation**
   - Calculate the MD5, SHA1, or SHA256 checksum of a file or standard input.
   - Command:
     ```bash
     futil checksum -f <filename> --md5|--sha1|--sha256
     ```
     Example:
     ```bash
     futil checksum -f myfile.txt --md5
     ```

3. **Error Handling**
   - Handles various error cases such as non-existent files, directories, and binary files.

## Application Design

The application is designed using the `cobra` package to provide an extensible command-line interface (CLI). Each command (such as `linecount` and `checksum`) is implemented as a separate module, making the tool easily maintainable and scalable for adding new features in the future.

- **Commands**
  - `linecount`: Counts the number of lines in a given file or standard input.
  - `checksum`: Calculates the checksum of a file using the specified algorithm (MD5, SHA1, or SHA256).
  
- **Binary File Detection**
  - For both commands, binary file detection is implemented. If the file is binary, the tool will return an error message:
    ```
    error: Cannot do linecount or checksum for binary file '<filename>'
    ```

## Third-Party Libraries

- [**Cobra**](https://github.com/spf13/cobra) - A library for creating powerful modern CLI applications.

## Known Issues

- **Large File Handling**: Processing very large files (several GBs or more) may result in high memory usage and slow performance. This can be mitigated by optimizing IO operations, but it's recommended to use the tool for small to medium-sized files. Future versions may include better handling for large files.

- **Binary File Detection**: The detection of binary files is done using a simple heuristic (checking for non-printable characters). This method might not be completely accurate for all file types, especially for files containing mixed binary and text content.
