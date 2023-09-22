package pkg

import (
	"github.com/go-toast/toast"
	"log"
	"os/exec"
	"runtime"
)

func Notify(message string) {
	switch runtime.GOOS {
	case "windows":
		notification := toast.Notification{
			Message: message,
			//Icon:    "path/to/icon.png",  // 可选项
			Actions: []toast.Action{},
		}
		err := notification.Push()
		if err != nil {
			log.Fatalf("failed to send notification: %v", err)
		}
	case "darwin":
		cmd := exec.Command("osascript", "-e", `display dialog "`+message+`"`)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to display dialog: %v", err)
		}
	case "linux":
		cmd := exec.Command("notify-send", "My App", message)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to send notification: %v", err)
		}
	default:
		log.Printf("Unsupported platform: %s", runtime.GOOS)
	}
}
