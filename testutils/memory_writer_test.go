// +build test
// +build test

package testutils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingleStringStorage(t *testing.T) {
	w := &MemoryWriter{}
	logger := log.New(w, "", 0)
	msg := "Hello bro"

	logger.Println(msg)

	require.Lenf(t, w.Data, 1, "Expected to have one message at writer but got", len(w.Data))

	assert.Equal(t, msg, w.Data[0])
}

func TestMultipleStringStorage(t *testing.T) {
	w := &MemoryWriter{}
	logger := log.New(w, "", 0)

	strs := []string{"Multiple", "Entries", "Here!"}
	for _, s := range strs {
		logger.Println(s)
	}
	striLen := len(strs)
	require.Lenf(t, w.Data, striLen, "Expected to have %d messages at writer but got %d", striLen, len(w.Data))
	assert.ElementsMatch(t, strs, w.Data, "Expected to have the equal slices")

}
