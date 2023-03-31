package slug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugifyModule(t *testing.T) {
	t.Run("New function should return a slug", func(t *testing.T) {
		slug := New("Hello World")
		assert.Equal(t, "hello-world", slug)
	})

	t.Run("New function should return a slug with Turkish characters", func(t *testing.T) {
		slug := New("Merhaba Dünya", TR)
		assert.Equal(t, "merhaba-dunya", slug)
	})

	t.Run("Is function should return true", func(t *testing.T) {
		slug := Is("hello-world")
		assert.Equal(t, true, slug)
	})

	t.Run("Is function should return false", func(t *testing.T) {
		slug := Is("hello world")
		assert.Equal(t, false, slug)
	})

	t.Run("Is function should return false if string length greater than 50", func(t *testing.T) {
		slug := Is("hello-world-hello-world-hello-world-hello-world-hello-world-hello-world")
		assert.Equal(t, false, slug)
	})

	t.Run("AddSub function should add a new sub", func(t *testing.T) {
		AddSub(EN, Sub{
			' ': "-",
		})
		slug := New("Hello World")
		assert.Equal(t, "hello-world", slug)
	})
}
