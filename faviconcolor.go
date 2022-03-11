package snowfavicon

import "sync"

var IconColors = []int{0x0090ed, 0xff4f5f, 0x2ac3a2, 0xff7139, 0xa172ff, 0xffa437, 0xff2a8a}

type FaviconColor struct {
	syncLock        *sync.RWMutex
	hardCodedColors map[string]int
}

func NewFaviconColor() *FaviconColor {
	return &FaviconColor{&sync.RWMutex{}, make(map[string]int)}
}

func (favicon *FaviconColor) HardCodeColor(a string, b int) {
	favicon.syncLock.Lock()
	favicon.hardCodedColors[a] = b
	favicon.syncLock.Unlock()
}

func (favicon *FaviconColor) PickColor(a string) int {
	favicon.syncLock.RLock()
	if c, ok := favicon.hardCodedColors[a]; ok {
		favicon.syncLock.RUnlock()
		return c
	}
	favicon.syncLock.RUnlock()

	hash := hashStrForColor(a)
	index := hash % int32(len(IconColors))
	return IconColors[index]
}

func hashStrForColor(a string) (hash int32) {
	r := []rune(a)
	for i := 0; i < len(r); i++ {
		hash += r[i]
	}
	return
}
