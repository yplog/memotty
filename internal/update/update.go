package update

import (
	"os"
	"os/exec"
)

func Run() error {
	cmd := exec.Command("bash", "-c", "curl -fsSL https://raw.githubusercontent.com/yplog/memotty/main/scripts/install.sh | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
