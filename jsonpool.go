package kingdee

import (
	"bytes"
	"encoding/json"
	"sync"
)

// SafeJSONPool 线程安全的 JSON 序列化池
type SafeJSONPool struct {
	pool sync.Pool
}

// NewSafeJSONPool 创建新的线程安全 JSON 池
func NewSafeJSONPool() *SafeJSONPool {
	return &SafeJSONPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

// Marshal 序列化任意对象为 JSON 字符串
func (p *SafeJSONPool) Marshal(v interface{}) (string, error) {
	buf := p.getBuffer()
	defer p.putBuffer(buf)

	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(v); err != nil {
		return "", err
	}

	return p.bufferToString(buf), nil
}

// MarshalIndent 带缩进的序列化
func (p *SafeJSONPool) MarshalIndent(v interface{}, prefix, indent string) (string, error) {
	buf := p.getBuffer()
	defer p.putBuffer(buf)

	data, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}

	buf.Write(data)
	return buf.String(), nil
}

// getBuffer 从池中获取 buffer
func (p *SafeJSONPool) getBuffer() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

// putBuffer 将 buffer 放回池中
func (p *SafeJSONPool) putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	p.pool.Put(buf)
}

// bufferToString 将 buffer 内容转换为字符串并处理换行符
func (p *SafeJSONPool) bufferToString(buf *bytes.Buffer) string {
	data := buf.Bytes()
	if len(data) > 0 && data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}
