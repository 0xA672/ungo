package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("=== ungo: Go toolchain uninstaller ===")

	goroot, err := findGOROOT()
	if err != nil {
		fmt.Printf("Error: unable to locate Go installation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found Go installation at: %s\n", goroot)

	if !isValidGOROOT(goroot) {
		fmt.Printf("Error: %s does not appear to be a valid Go installation (bin/go or pkg directory missing)\n", goroot)
		os.Exit(1)
	}

	fmt.Print("Are you sure you want to delete this Go installation? This cannot be undone! (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	if confirm != "yes" && confirm != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	fmt.Printf("Removing %s ...\n", goroot)
	if err := os.RemoveAll(goroot); err != nil {
		fmt.Printf("Error: failed to remove directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Go installation directory removed successfully.")

	printPATHWarning(goroot)
}

func findGOROOT() (string, error) {
	if goroot := os.Getenv("GOROOT"); goroot != "" {
		return goroot, nil
	}

	if goPath, err := exec.LookPath("go"); err == nil {
		cmd := exec.Command(goPath, "env", "GOROOT")
		out, err := cmd.Output()
		if err == nil {
			return strings.TrimSpace(string(out)), nil
		}
	}

	return guessDefaultGOROOT(), nil
}

func guessDefaultGOROOT() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/local/go"
	case "linux":
		return "/usr/local/go"
	case "windows":
		return `C:\Go`
	default:
		return "/usr/local/go"
	}
}

func isValidGOROOT(path string) bool {
	binGo := filepath.Join(path, "bin", "go")
	if runtime.GOOS == "windows" {
		binGo += ".exe"
	}
	if _, err := os.Stat(binGo); err == nil {
		return true
	}

	_, errPkg := os.Stat(filepath.Join(path, "pkg"))
	_, errSrc := os.Stat(filepath.Join(path, "src"))
	return errPkg == nil || errSrc == nil
}

func printPATHWarning(goroot string) {
	fmt.Println()
	fmt.Println("Warning: you must manually remove the following directory from your PATH environment variable:")
	fmt.Printf("  %s\n", filepath.Join(goroot, "bin"))
	fmt.Println()
	fmt.Println("Recommended steps:")
	switch runtime.GOOS {
	case "windows":
		fmt.Println("  1. Open System Properties -> Advanced -> Environment Variables")
		fmt.Println("  2. Remove the above path from the PATH variable")
		fmt.Println("  3. Restart your command prompt")
	default:
		fmt.Println("  1. Edit ~/.bashrc, ~/.zshrc, or /etc/profile")
		fmt.Println("  2. Remove the line that adds the Go binary to PATH (e.g., export PATH=.../go/bin:$PATH)")
		fmt.Println("  3. Run 'source ~/.bashrc' or restart your terminal")
	}
	fmt.Println()
	fmt.Println("Uninstall complete. Go has been removed from your system.")
}
