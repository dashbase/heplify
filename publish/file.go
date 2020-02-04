package publish

import (
	"github.com/negbie/logp"
)

type FileOutputer struct {
}

func (fo *FileOutputer) Output(msg []byte) {
	/* 	jsonPkt, err := json.MarshalIndent(msg, "", "  ")
	   	if err != nil {
	   		logp.Err("json %v", err)
	   		return
	   	}
		   logp.Info("%s", jsonPkt)
	*/
	h, err := DecodeHEP(msg)
	if err == nil {
		logp.Info("%s\n", h.toLogString())
	} else {
		logp.Warn("%s", err)
	}
}

func NewFileOutputer() (*FileOutputer, error) {
	fo := &FileOutputer{}
	return fo, nil
}
