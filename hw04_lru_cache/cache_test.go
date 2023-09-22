package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("Превышение размера кеша", func(t *testing.T) {
		c := NewCache(3)

		// заполняем кеш
		inCache := c.Set("aaa", 1)
		require.False(t, inCache)
		inCache = c.Set("bbb", 2)
		require.False(t, inCache)
		inCache = c.Set("ccc", 3)
		require.False(t, inCache)

		// получаем значение и меняем порядок элементов
		val, ok := c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 2, val)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 3, val)

		// добавлем еще один элемент для превышения емкости кэша
		// должно вытолкнуть из кэша значение ааа как находящийся в конце
		inCache = c.Set("ddd", 4)
		require.False(t, inCache)

		// проверяем есть ли в кэше удаленный элемент
		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		// сохранился ли элемент ddd в кеше
		val, ok = c.Get("ddd")
		require.True(t, ok)
		require.Equal(t, 4, val)
	})

	// очистка кэша
	t.Run("Clear cache", func(t *testing.T) {
		c := NewCache(3)
		c.Set("aaa", 1)
		c.Set("bbb", 2)
		c.Set("ccc", 3)
		c.Clear()
		_, ok := c.Get("aaa")
		require.False(t, ok)
		_, ok = c.Get("bbb")
		require.False(t, ok)
		_, ok = c.Get("ccc")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
