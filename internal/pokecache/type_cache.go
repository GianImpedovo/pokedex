package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	info map[string]cacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		info: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reaploop(interval)
	return c
}

func (c *Cache) reaploop(interval time.Duration) {
	// 1. El Ticker es como una alarma recurrente.
	// Si el intervalo es de 5 minutos, la alarma sonará cada 5 minutos.
	t := time.NewTicker(interval)

	// 2. Este bucle es especial. No corre a toda velocidad.
	// Se queda "esperando" en la línea 'range t.C' hasta que la alarma suena.
	// Cuando la alarma suena, el código de adentro se ejecuta una vez.
	for range t.C {

		// 3. ¡La alarma sonó! Antes de tocar nada, ponemos el candado.
		// Esto evita que alguien intente guardar (Add) o leer (Get)
		// mientras nosotros estamos borrando cosas.
		c.mu.Lock()

		// 4. Revisamos cada una de las entradas que tenemos guardadas.
		for k, v := range c.info {
			// 5. 'time.Since(v.createdAt)' calcula cuánto tiempo ha pasado
			// desde que guardamos este dato hasta el momento actual.
			// Si ese tiempo es mayor al intervalo permitido...
			if time.Since(v.createdAt) > interval {
				// ... ¡entonces el dato es demasiado viejo! Lo borramos.
				delete(c.info, k)
			}
		}

		// 6. Terminamos de limpiar, así que soltamos el candado para que
		// el resto del programa pueda volver a usar la cache.
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.info[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.info[key]; ok {
		return v.val, true
	}
	return nil, false
}
