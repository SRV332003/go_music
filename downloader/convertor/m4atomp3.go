package convertor

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ossrs/go-oryx-lib/aac"
	"github.com/viert/lame"
)

func M4aToMp3(m4aStream io.ReadCloser, mp3Path string) error {
	m4aBytes, err := io.ReadAll(m4aStream)
	if err != nil {
		return err
	}

	log.Println(m4aBytes[:16])

	var d aac.ADTS

	if d, err = aac.NewADTS(); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return err
	}

	var pcm []byte
	pcm, _, err = d.Decode(m4aBytes)
	if err != nil {
		fmt.Println("decode failed, err is", err)
		return err
	}

	// Create the output file
	mp3File, err := os.OpenFile(mp3Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	// Convert the m4a file to mp3

	wr := lame.NewWriter(mp3File)
	wr.Encoder.SetBitrate(d.ASC().SampleRate.ToHz())
	wr.Encoder.SetQuality(1)

	pcmReader := bytes.NewReader(pcm)
	io.Copy(wr, pcmReader)

	return nil
}
