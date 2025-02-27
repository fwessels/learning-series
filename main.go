package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func object() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("a", 16*10000))
	out.WriteString(strings.Repeat("b", 16*10000))
	out.WriteString(strings.Repeat("c", 16*10000))
	out.WriteString(strings.Repeat("d", 16*10000))
	out.WriteString(strings.Repeat("e", 16*10000))

	os.WriteFile("object", []byte(out.String()), 0644)
}

func object2() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("abcdefghijklmnop", 10000))
	out.WriteString(strings.Repeat("qrstuvwxyzABCDEF", 10000))
	out.WriteString(strings.Repeat("GHIJKLMNOPQRSTUV", 10000))
	out.WriteString(strings.Repeat("WXYZ0123456789!@", 10000))
	out.WriteString(strings.Repeat(`#$%^&*(){}[]:";'`, 10000))

	os.WriteFile("object2", []byte(out.String()), 0644)
}

func small() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("a", 16))
	out.WriteString(strings.Repeat("b", 16))
	out.WriteString(strings.Repeat("c", 16))
	out.WriteString(strings.Repeat("d", 16))
	out.WriteString(strings.Repeat("e", 16))

	os.WriteFile("small", []byte(out.String()), 0644)
}

func reduced() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("a", 16*10000))
	out.WriteString(strings.Repeat("b", 16*10000))
	out.WriteString(strings.Repeat("c", 16*10000))
	out.WriteString(strings.Repeat("d", 16*10000))
	out.WriteString(strings.Repeat("e", 16*10000))
	out.WriteString(strings.Repeat("f", 16*10000))

	os.WriteFile("reduced", []byte(out.String()), 0644)
}

func blocks() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("a", 209716))
	out.WriteString(strings.Repeat("b", 209716))
	out.WriteString(strings.Repeat("c", 209716))
	out.WriteString(strings.Repeat("d", 209716))
	out.WriteString(strings.Repeat("e", 209712))

	out.WriteString(strings.Repeat("f", 16*10000))
	out.WriteString(strings.Repeat("g", 16*10000))
	out.WriteString(strings.Repeat("h", 16*10000))
	out.WriteString(strings.Repeat("i", 16*10000))
	out.WriteString(strings.Repeat("j", 16*10000))

	os.WriteFile("blocks", []byte(out.String()), 0644)
}

func blocks3() {
	out := strings.Builder{}
	out.WriteString(strings.Repeat("a", 209716))
	out.WriteString(strings.Repeat("b", 209716))
	out.WriteString(strings.Repeat("c", 209716))
	out.WriteString(strings.Repeat("d", 209716))
	out.WriteString(strings.Repeat("e", 209712))

	out.WriteString(strings.Repeat("f", 209716))
	out.WriteString(strings.Repeat("g", 209716))
	out.WriteString(strings.Repeat("h", 209716))
	out.WriteString(strings.Repeat("i", 209716))
	out.WriteString(strings.Repeat("j", 209712))

	out.WriteString(strings.Repeat("k", 16*10000))
	out.WriteString(strings.Repeat("l", 16*10000))
	out.WriteString(strings.Repeat("m", 16*10000))
	out.WriteString(strings.Repeat("n", 16*10000))
	out.WriteString(strings.Repeat("o", 16*10000))

	os.WriteFile("blocks3", []byte(out.String()), 0644)
}

func partialGet(start, end int64) {
	var minioClient *minio.Client
	var err error
	if minioClient, err = minio.New("192.168.40.180:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	}); err != nil {
		log.Fatalf("Error establishing MinIO client: %v", err)
	}

	opts := minio.GetObjectOptions{}
	if err = opts.SetRange(start, end); err != nil {
		log.Fatalf("Error setting range: %v", err)
	}

	object, err := minioClient.GetObject(context.Background(), "test", "blocks3", opts)
	if err != nil {
		log.Fatalf("Error getting object: %v", err)
	}
	defer object.Close()

	// 4. Read and process the partial data
	partialData, err := io.ReadAll(object)
	if err != nil {
		log.Fatalf("Error reading partial data: %v", err)
	}

	// 5. Print or otherwise handle the partial data
	fmt.Print(hex.Dump(partialData))
}

func main() {

	object()
	object2()
	small()
	reduced()
	blocks()
	blocks3()

	// partialGet(0, 63) // aaaa
	// partialGet(0x00100000, 0x00100000+63) // fff
	// partialGet(0x00200000, 0x00200000+63)
}
