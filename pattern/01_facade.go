package pattern

import (
	"fmt"
	"sync"
)

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике. https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Паттерн фасад используется для упрощения взаимодействия со сложными системами, предоставляя упрощённый интерфейс, который скрывает детали внутренней реализации

Плюсы:
	1) Упрощенное использование: один интерфейс для взаимодействия вместо кучи мелких
	2) Низкая связность: взаимодействие зависит только от интерфейса фасада, а не его подсистем.
Можно легко изменять интерфейсы подсистем, не влияя на интерфейс фасада

Минусы:
	1) Ограниченная гибкость. Инкапсуляция логики подсистем может не обеспечивать нужной гибкости в работе (но как будто это ошибка реализации, а минус)


Что в этом примере?
Реализован фасад для работы с базой данных и кэшем. Работа с бд предполагает взаимодействие с кэшем (ускорение работы путем сохранения данных in-memory)
Фасад позволяет инкапсулировать взаимодействие этих подсистем
*/

type Database struct {
	client *DBClient
	cache  *Cache
}

func NewDatabase() *Database {
	return &Database{
		client: NewDBClient(),
		cache:  NewCache(),
	}
}

func (d *Database) Create(id string, data any) error {
	if err := d.client.Insert(id, data); err != nil {
		// тут можно добавить работу с другими подсистемами, например логирование ошибки
		return err
	}
	d.cache.Set(id, data)
	return nil
}

func (d *Database) Update(id string, data any) error {
	if err := d.client.Update(id, data); err != nil {
		// тут можно добавить работу с другими подсистемами, например логирование ошибки
		return err
	}
	d.cache.Set(id, data)
	return nil
}

func (d *Database) Find(id string) (any, error) {
	if v, ok := d.cache.Get(id); ok {
		return v, nil
	}
	data, err := d.client.Select(id)
	if err != nil {
		// тут можно добавить работу с другими подсистемами, например логирование ошибки
		return nil, err
	}
	d.cache.Set(id, data)
	return data, nil
}

func (d *Database) Delete(id string) error {
	if err := d.client.Delete(id); err != nil {
		// тут можно добавить работу с другими подсистемами, например логирование ошибки
		return err
	}
	d.cache.Del(id)
	return nil
}

// DBClient реализует внутренний компонент системы. В данном примере это работа с базой данных
type DBClient struct {
	// тут разные подключения и все связанное с бд
}

func NewDBClient() *DBClient {
	return &DBClient{}
}

func (c *DBClient) Insert(id string, data any) error {
	fmt.Printf("created new row in db by id %s with data %+v\n", id, data)
	return nil
}

func (c *DBClient) Select(id string) (data any, err error) {
	fmt.Printf("selected row from db by id %s\n", id)
	return nil, nil
}

func (c *DBClient) Update(id string, data any) error {
	fmt.Printf("updated row in db by id %s with data %+v\n", id, data)
	return nil
}

func (c *DBClient) Delete(id string) error {
	fmt.Printf("deleted row in db by id %s\n", id)
	return nil
}

// Cache реализует внутренний компонент системы. В данном примере это кэширование данных
type Cache struct {
	// Тут все связанное с кэшем
	mx   sync.RWMutex
	data map[string]any
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (value any, ok bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok = c.data[key]
	return
}

func (c *Cache) Del(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.data, key)
}
