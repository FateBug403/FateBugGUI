package pocs

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	data,err := EmbedReadContentByName("azure-kubernetes-service")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
}