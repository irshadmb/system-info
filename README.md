# System Information Collector

A Go application that collects detailed system information using Windows Management Instrumentation (WMI) and outputs it in JSON format.

## Features

- Operating System information (version, architecture, memory)
- CPU details (cores, speed, load)
- Disk drive information (size, interface type)
- Network adapter details (name, MAC address, speed)

## Prerequisites

- Go 1.21 or higher
- Windows operating system
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/irshadmb/system-info.git
cd system-info
```

2. Install dependencies:
```bash
go mod download
```

## Usage

Build and run the application:

```bash
go build ./cmd/main.go
./main
```

The application will output system information in JSON format.

## Project Structure

```
system-info/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── collectors/       # System information collectors
│   └── models/          # Data structures
├── go.mod               # Go module file
└── README.md            # This file
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
