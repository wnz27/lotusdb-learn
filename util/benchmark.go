/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-07-09 00:54:56
 * @LastEditTime: 2023-07-09 02:16:51
 * @FilePath: /lotusdb-learn/util/benchmark.go
 * @description: type some description
 */

package util

import (
	"os"
	"path/filepath"
)

const (
	DefaultKeyNum = 10000
)

// BenchmarkStore interface for kv store
type BenchmarkStore interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Exist(key []byte) (bool, error)
	Delete(key []byte) error
	Close() error
}

// dirSize calculate path dir's files size
func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// keysNum assign numbers of key, current support max is 200000, default is DefaultKeyNum (10000)
func keysNum(maxNum int) int {
	if maxNum == 0 {
		return DefaultKeyNum
	}
	return maxNum
}

type TestScale struct {
	KeysNum int // numbers of key

	// calculate random size with min and max, key and value
	MinKeySize   int
	MaxKeySize   int
	MinValueSize int
	MaxValueSize int

	Concurrency int    // concurrency run numbers, goroutines in go.
	Path        string // database path
	Compact     bool   // is benchmark compaction. write keys twice and run compaction after
	Collection  bool   // is benchmark collection. write some keys and delete keys, then use collection.
}
