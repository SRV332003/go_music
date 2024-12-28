package player

import (
	"bytes"
	"log"
	"os/exec"
)

func StopIfPlaying() error {
	var buffer bytes.Buffer
	//cat /proc/asound/card*/pcm*/sub0/status | grep RUNNING
	cmd := exec.Command("find", "/proc/asound/card*/pcm*/sub0/status")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error while checking if audio is playing: %s\n%s", err, stderr.String())
		return err
	}
	buffer.Write(output)
	log.Printf("Vault login output: %s", buffer.String())
	return nil
}
