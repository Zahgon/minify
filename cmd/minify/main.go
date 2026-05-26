package main

import (
	"io/fs"
	"log"
	"os"
	"regexp"
	"time"

	min "github.com/tdewolff/minify/v2"
)

// Version is the current minify version.
var Version = "built from source"

var extMap = map[string]string{
	"asp":         "text/asp",
	"css":         "text/css",
	"ejs":         "text/x-ejs-template",
	"gohtml":      "text/x-go-template",
	"handlebars":  "text/x-handlebars-template",
	"htm":         "text/html",
	"html":        "text/html",
	"js":          "application/javascript",
	"json":        "application/json",
	"mjs":         "application/javascript",
	"mustache":    "text/x-mustache-template",
	"php":         "application/x-httpd-php",
	"rss":         "application/rss+xml",
	"svg":         "image/svg+xml",
	"tmpl":        "text/x-template",
	"webmanifest": "application/manifest+json",
	"xhtml":       "application/xhtml+xml",
	"xml":         "text/xml",
}

var (
	help               bool
	hidden             bool
	inplace            bool
	list               bool
	m                  *min.M
	matches            []string
	matchesRegexp      []*regexp.Regexp
	filters            []string
	filtersRegexp      []*regexp.Regexp
	extensions         map[string]string
	recursive          bool
	quiet              bool
	verbose            int
	version            bool
	watch              bool
	sync               bool
	bundle             bool
	preserve           []string
	preserveMode       bool
	preserveOwnership  bool
	preserveTimestamps bool
	preserveLinks      bool
	mimetype           string
	oldmimetype        string
)

type Matches struct {
	matches *[]string
}

func (m Matches) Help() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (m Matches) Scan(name string, s []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Includes struct {
	filters *[]string
}

func (m Includes) Help() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (m Includes) Scan(name string, s []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Excludes struct {
	filters *[]string
}

func (m Excludes) Help() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (m Excludes) Scan(name string, s []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Task is a minify task.
type Task struct {
	root string
	srcs []string
	dst  string
	sync bool
}

// NewTask returns a new Task.
func NewTask(root, input, output string, sync bool) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// in-place

// Loggers.
var (
	Error   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Debug   *log.Logger
)

func main() {
	// os.Exit doesn't execute pending defer calls, this is fixed by encapsulating run()
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

//f.AddOpt(&htmlMinifier.TemplateDelims, "", "html-template-delims", "Set template delimiters explicitly, for example <?,?> for PHP or {{,}} for Go templates") // TODO: fix parsing {{ }} in tdewolff/argp

// stdin

// stdout

// compile matches and regexps

// detect mimetype, mimetype=="" means we'll infer mimetype from file extensions

////////////////

// set output file or directory, empty means stdout

// concatenate

// Task.sync == false because dirDst == false

// make output directory

////////////////

// also handles <?php

// skip change on output

// skip change on output

func minifyWorker(chanTasks <-chan Task, chanFails chan<- int) { _ = "STUB: not implemented"; return }

// compilePattern returns *regexp.Regexp or glob.Glob
func compilePattern(pattern string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileFilter(filename string) bool { _ = "STUB: not implemented"; return false }

func fileMatches(filename string) bool { _ = "STUB: not implemented"; return false }

func createTasks(fsys fs.FS, inputs []string, output string) ([]Task, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// follow and dereference symlinks

// copy symlink as is

// don't filter mimetype

// follow and dereference symlinks

// copy symlink as is

func minify(t Task) bool {
	_ = "STUB: not implemented"
	// synchronizing files that are not minified but just copied to the same directory, no action needed
	return false
}

// rename original when overwriting

// synchronize file

// copy original

// remove original that was renamed, when overwriting files

func retry(attempts int, fn func() error) (err error) { _ = "STUB: not implemented"; return nil }

func preserveAttributes(srcs []string, root, dst string) { _ = "STUB: not implemented"; return }

// only with --bundle we allow 1 < len(srcs)

// make sure we only set attributes on directories and files inside the root destination

// should never occur

// when using --bundle there are no directories to preserve

// go up to but excluding the root path

func formatBytes(size uint64) string { _ = "STUB: not implemented"; return "" }

func formatDuration(dur time.Duration) string { _ = "STUB: not implemented"; return "" }
