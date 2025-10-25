package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func removeBG(w http.ResponseWriter, r *http.Request) {
	// 50 MB limit
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "bad multipart: "+err.Error(), 400)
		return
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file missing: "+err.Error(), 400)
		return
	}
	defer f.Close()
	src, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, "read file: "+err.Error(), 500)
		return
	}

	py := os.Getenv("PYTHON_BIN")
	if py == "" {
		py = "/Users/bimap/Documents/Coding/Project/AngelaClothes/server/python/.venv/bin/python" // be explicit
	}
	script := "/Users/bimap/Documents/Coding/Project/AngelaClothes/server/python/bgrm.py"
	log.Printf("using python=%s script=%s", py, script)

	ctx, cancel := context.WithTimeout(r.Context(), 45*time.Second)
	defer cancel()
	
	cmd := exec.CommandContext(ctx, py, script)
	cmd.Stdin = bytes.NewReader(src)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		log.Printf("rembg err=%v stderr=%q", err, errBuf.String())
		http.Error(w, "rembg failed: "+err.Error()+" | "+errBuf.String(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(200)
	w.Write(out.Bytes())
}

func main() {
	http.HandleFunc("/remove-bg", removeBG)
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("OK")) })
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
