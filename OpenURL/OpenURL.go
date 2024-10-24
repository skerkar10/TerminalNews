package OpenURL

import (
  "fmt"
  "os/exec"
  "runtime"
)

func OpenLink(URL string) {
  switch runtime.GOOS {
  case "linux":
    exec.Command("xdg-open", URL).Start()
  case "darwin":
    exec.Command("open", URL).Start()
  default:
    fmt.Errorf("\nThis Operating System is not supported yet. Please stay tuned for future support!\n\n")
  }
}
