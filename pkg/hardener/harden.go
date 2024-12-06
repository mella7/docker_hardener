package hardener

import (
    "fmt"
    "os/exec"
)

// HardenDockerfile applies security best practices to a Dockerfile.
func HardenDockerfile(dockerfilePath string) error {
    // Example: Disable apt-get recommended packages (hardening example)
    err := exec.Command("sed", "-i", "s/RUN apt-get install .*/RUN apt-get install --no-install-recommends/g", dockerfilePath).Run()
    if err != nil {
        return fmt.Errorf("error hardening Dockerfile: %v", err)
    }
    return nil
}

// SetNonRootUser modifies the Dockerfile to use a non-root user.
func SetNonRootUser(dockerfilePath string) error {
    err := exec.Command("sed", "-i", "/USER root/a USER nonrootuser", dockerfilePath).Run()
    if err != nil {
        return fmt.Errorf("error setting non-root user: %v", err)
    }
    return nil
}
