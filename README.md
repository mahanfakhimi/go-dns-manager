# DNS Manager

DNS Manager is a command-line tool to easily switch between different DNS providers on a Windows machine.

## Features

- Switch between multiple DNS providers including Shecan, Electro, DNS403, Google, Cloudflare, and OpenDNS.
- Clear DNS settings to use DHCP.
- Requires administrative privileges to change DNS settings.

## DNS Providers

- **Shecan**: 178.22.122.100, 185.51.200.2
- **Electro**: 78.157.42.100, 78.157.42.101
- **DNS403**: 10.202.10.202, 10.202.10.102
- **Google**: 8.8.8.8, 8.8.4.4
- **Cloudflare**: 1.1.1.1, 1.0.0.1
- **OpenDNS**: 208.67.222.222, 208.67.220.220
- **None**: Clear DNS settings

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/DNS-Manager.git
   cd DNS-Manager
   ```

2. Build the project:
   ```sh
   go build -o DNS-Manager.exe
   ```

## Usage

1. Run the executable:

   ```sh
   ./DNS-Manager.exe
   ```

2. Follow the on-screen instructions to select a DNS provider or clear DNS settings.

## Requirements

- Go 1.23.1 or later
- Windows operating system

## License

This project is licensed under the MIT License.
