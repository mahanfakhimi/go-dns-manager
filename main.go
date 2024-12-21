package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type DNSProvider string

const (
	Shecan     DNSProvider = "shecan"
	Electro    DNSProvider = "electro"
	DNS403     DNSProvider = "403"
	Google     DNSProvider = "google"
	Cloudflare DNSProvider = "cloudflare"
	OpenDNS    DNSProvider = "opendns"
	None       DNSProvider = "none"
)

var (
	DNSAddresses = map[DNSProvider][]string{
		Shecan:     {"178.22.122.100", "185.51.200.2"},
		Electro:    {"78.157.42.100", "78.157.42.101"},
		DNS403:     {"10.202.10.202", "10.202.10.102"},
		Google:     {"8.8.8.8", "8.8.4.4"},
		Cloudflare: {"1.1.1.1", "1.0.0.1"},
		OpenDNS:    {"208.67.222.222", "208.67.220.220"},
		None:       {},
	}

	DNSMenu = map[int]DNSProvider{
		1: Shecan,
		2: Electro,
		3: DNS403,
		4: Google,
		5: Cloudflare,
		6: OpenDNS,
		0: None,
	}
)

func runAsAdmin() {
	if len(os.Args) > 1 && os.Args[1] == "--admin" {
		return
	}

	cmd := exec.Command("powershell", "-Command", "Start-Process", os.Args[0], "-ArgumentList", "--admin", "-Verb", "runAs")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to run as administrator:", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func applyDNS(provider DNSProvider) {

	if provider == None {
		clearDNSSettings()
	} else {
		setDNSSettings(provider, DNSAddresses[provider])
	}
}

func setDNSSettings(provider DNSProvider, addresses []string) {
	if len(addresses) < 2 {
		fmt.Println("Invalid DNS configuration.")
		return
	}

	primary := addresses[0]
	secondary := addresses[1]

	primaryCmd := exec.Command("netsh", "interface", "ip", "set", "dns", "Wi-Fi", "static", primary)
	secondaryCmd := exec.Command("netsh", "interface", "ip", "add", "dns", "Wi-Fi", secondary)

	if err := primaryCmd.Run(); err != nil {
		fmt.Println("Error setting primary DNS:", err)
		return
	}

	if err := secondaryCmd.Run(); err != nil {
		fmt.Println("Error adding secondary DNS:", err)
		return
	}

	clearScreen()

	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("DNS successfully set to %s.\n", provider)
	color.Unset()
}

func clearDNSSettings() {
	cmd := exec.Command("netsh", "interface", "ip", "set", "dns", "Wi-Fi", "dhcp")

	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing DNS settings:", err)
	}

	clearScreen()

	color.Set(color.FgGreen, color.Bold)
	fmt.Println("DNS settings cleared successfully.")
	color.Unset()
}

func showMenu() {
	color.Set(color.FgBlue, color.Bold)
	fmt.Println("1: shecan")
	fmt.Println("2: electro")
	fmt.Println("3: 403")
	fmt.Println("4: google")
	fmt.Println("5: cloudflare")
	fmt.Println("6: opendns")
	fmt.Println("0: none")
	color.Unset()

	fmt.Print("Select DNS from the list: ")
}

func clearScreen() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	runAsAdmin()

	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil || choice < 0 || choice > 6 || input[0] == '0' && len(input) > 1 {
			clearScreen()

			color.Set(color.FgRed, color.Bold)
			fmt.Println("Invalid input. Please enter a valid number.")
			color.Unset()

			continue
		}

		selectedDNS, exists := DNSMenu[choice]

		if !exists {
			clearScreen()

			color.Set(color.FgRed, color.Bold)
			fmt.Println("Invalid choice. Please try again.")
			color.Unset()
		}

		applyDNS(selectedDNS)
	}
}
