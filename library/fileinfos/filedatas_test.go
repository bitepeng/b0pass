package fileinfos

import (
	"log"
	"testing"
)

func TestDataInit(t *testing.T) {
	DataSet("data_path","appsss")
	DataInit("data_path","data_text")
	log.Println(DataGet("data_path"))
}