package main

import (
	"github.com/fsnotify/fsnotify"
)

// Watcher is a wrapper for watching file changes in directories.
type Watcher struct {
	watcher    *fsnotify.Watcher
	dirs       map[string]bool
	paths      map[string]bool
	ignoreNext map[string]bool
	recursive  bool
}

// NewWatcher returns a new Watcher.
func NewWatcher(recursive bool) (*Watcher, error) { _ = "STUB: not implemented"; return nil, nil }

// Close closes the watcher.
func (w *Watcher) Close() error { _ = "STUB: not implemented"; return nil }

// IgnoreNext ignores the next change on a path.
func (w *Watcher) IgnoreNext(path string) { _ = "STUB: not implemented"; return }

// AddPath adds a new path to watch.
func (w *Watcher) AddPath(path string) error { _ = "STUB: not implemented"; return nil }

// Run watches for file changes.
func (w *Watcher) Run() chan string { _ = "STUB: not implemented"; return nil }

// prevent reminifying the first time for files with input==output

// check if changed file is being watched (as a file or indirectly in a dir)

// wait to make sure write is finished
