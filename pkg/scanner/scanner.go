package scanner

import (
    "fmt"
    "os/exec"
)

// ScanImage uses Trivy to scan a Docker image for known vulnerabilities.
func ScanImage(imageName string) error {
    cmd := exec.Command("trivy", "image", imageName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error scanning image %s: %v\n%s", imageName, err, output)
    }
    fmt.Println(string(output)) // Print the scan results
    return nil
}
