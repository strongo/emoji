package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestMain_Success(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)
	if err := os.MkdirAll(filepath.Join(tempDir, "go", "emoji"), 0o755); err != nil {
		t.Fatalf("failed to create dirs: %v", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`[{"emoji":"😀","description":"grinning","category":"smileys"}]`))
	}))
	defer ts.Close()

	origURL := gemojiDBJsonURL
	origFatal := logFatal
	defer func() {
		gemojiDBJsonURL = origURL
		logFatal = origFatal
	}()

	gemojiDBJsonURL = ts.URL
	var fatalCalled bool
	logFatal = func(_ ...any) {
		fatalCalled = true
	}

	main()
	if fatalCalled {
		t.Error("expected main() to succeed without calling logFatal")
	}
}

func TestMain_GemojiError(t *testing.T) {
	origURL := gemojiDBJsonURL
	origFatal := logFatal
	defer func() {
		gemojiDBJsonURL = origURL
		logFatal = origFatal
	}()

	gemojiDBJsonURL = "http://127.0.0.1:0/invalid"
	var fatalCalled bool
	logFatal = func(_ ...any) {
		fatalCalled = true
	}

	main()
	if !fatalCalled {
		t.Error("expected logFatal to be called when getGemoji fails")
	}
}

func TestMain_GenerateError(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir) // do not create go directory so Generate fails

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`[{"emoji":"😀","description":"grinning","category":"smileys"}]`))
	}))
	defer ts.Close()

	origURL := gemojiDBJsonURL
	origFatal := logFatal
	defer func() {
		gemojiDBJsonURL = origURL
		logFatal = origFatal
	}()

	gemojiDBJsonURL = ts.URL
	var fatalCalled bool
	logFatal = func(_ ...any) {
		fatalCalled = true
	}

	main()
	if !fatalCalled {
		t.Error("expected logFatal to be called when Generate fails")
	}
}

func TestGetGemoji(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"emoji":"😀","description":"grinning","category":"smileys"}]`))
		}))
		defer ts.Close()

		emojis, err := getGemoji(ts.URL)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(emojis) != 1 || emojis[0].Emoji != "😀" {
			t.Errorf("unexpected emojis: %+v", emojis)
		}
	})

	t.Run("http_error", func(t *testing.T) {
		_, err := getGemoji("http://127.0.0.1:0/invalid")
		if err == nil {
			t.Fatal("expected http error, got nil")
		}
	})

	t.Run("read_body_error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "100")
			w.WriteHeader(http.StatusOK)
			conn, _, err := w.(http.Hijacker).Hijack()
			if err == nil {
				_ = conn.Close()
			}
		}))
		defer ts.Close()

		_, err := getGemoji(ts.URL)
		if err == nil {
			t.Fatal("expected read error, got nil")
		}
	})

	t.Run("json_error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`invalid-json`))
		}))
		defer ts.Close()

		_, err := getGemoji(ts.URL)
		if err == nil {
			t.Fatal("expected unmarshal error, got nil")
		}
	})
}

func TestGoGenerator_GetWriter(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)
	if err := os.MkdirAll(filepath.Join(tempDir, "go"), 0o755); err != nil {
		t.Fatalf("failed to create go dir: %v", err)
	}

	gen := newGoGenerator()

	// Empty ID returns nil, nil
	w, err := gen.getWriter("category", "")
	if w != nil || err != nil {
		t.Errorf("expected nil, nil for empty ID, got w=%v, err=%v", w, err)
	}

	// First call creates writer
	w1, err := gen.getWriter("category", "nature")
	if err != nil || w1 == nil {
		t.Fatalf("expected writer, got: %v", err)
	}

	// Second call returns cached writer
	w2, err := gen.getWriter("category", "nature")
	if err != nil || w2 != w1 {
		t.Fatalf("expected cached writer, got: %v", err)
	}

	// Create error when go directory missing
	t.Chdir(t.TempDir())
	_, err = gen.getWriter("category", "missing")
	if err == nil {
		t.Fatal("expected error when go dir does not exist")
	}
}

func TestGoGenerator_Generate_Comprehensive(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)
	if err := os.MkdirAll(filepath.Join(tempDir, "go", "emoji"), 0o755); err != nil {
		t.Fatalf("failed to create dirs: %v", err)
	}

	emojis := []GemojiEmoji{
		{
			Emoji: "", // skipped
		},
		{
			Emoji:       "🎈",
			Description: "", // skipped in constants
			Category:    "", // default to uncategorized
		},
		{
			Emoji:       "🇹🇷",
			Description: "flag Turkey",
			Category:    "flags",
		},
		{
			Emoji:       "🏘",
			Description: "houses",
			Category:    "places",
		},
		{
			Emoji:       "1️⃣",
			Description: "keycap: 1st  place & fun’s - test. : ; , “quote” (parens) ! # * 2nd 3rd",
			Category:    "symbols",
		},
	}

	gen := newGoGenerator()
	if err := gen.Generate(emojis); err != nil {
		t.Fatalf("unexpected Generate error: %v", err)
	}

	// Verify generated files exist
	for _, expectedFile := range []string{
		"go/category_uncategorized.go",
		"go/category_flags.go",
		"go/category_places.go",
		"go/category_symbols.go",
		"go/emoji_all_map.go",
		"go/emoji/constants.go",
	} {
		if _, err := os.Stat(filepath.Join(tempDir, expectedFile)); err != nil {
			t.Errorf("expected file %s to exist: %v", expectedFile, err)
		}
	}
}

func TestGoGenerator_Generate_Errors(t *testing.T) {
	emojis := []GemojiEmoji{
		{
			Emoji:       "😀",
			Description: "grinning",
			Category:    "smileys",
		},
	}

	t.Run("writeItem_error", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir) // no go/ dir
		gen := newGoGenerator()
		if err := gen.Generate(emojis); err == nil {
			t.Fatal("expected writeItem error")
		}
	})

	t.Run("generateMapFile_error", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)
		if err := os.MkdirAll(filepath.Join(tempDir, "go", "emoji_all_map.go"), 0o755); err != nil {
			t.Fatal(err)
		}
		gen := newGoGenerator()
		if err := gen.Generate(emojis); err == nil {
			t.Fatal("expected generateMapFile error")
		}
	})

	t.Run("generateConstants_error", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)
		if err := os.MkdirAll(filepath.Join(tempDir, "go"), 0o755); err != nil {
			t.Fatal(err)
		}
		// do not create go/emoji dir, so constants.go cannot be created
		gen := newGoGenerator()
		if err := gen.Generate(emojis); err == nil {
			t.Fatal("expected generateConstants error")
		}
	})
}
