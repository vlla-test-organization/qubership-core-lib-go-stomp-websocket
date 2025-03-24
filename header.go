package go_stomp_websocket

type Header struct {
	header []string
}

func (h *Header) Contains(key string) (value string, ok bool) {
	var i int
	if i, ok = h.index(key); ok {
		value = h.header[i+1]
	}
	return
}

func (h *Header) index(key string) (int, bool) {
	for i := 0; i < len(h.header); i += 2 {
		if h.header[i] == key {
			return i, true
		}
	}
	return -1, false
}

func (h *Header) Add(key, value string) {
	h.header = append(h.header, key, value)
}

func (h *Header) Get(key string) string {
	value, _ := h.Contains(key)
	return value
}
