package dingtalk

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Expirable interface {
	CreatedAt() int64
	ExpiresIn() int
}

type Cache interface {
	Set(data Expirable) error
	Get(data Expirable) error
}

type FileCache struct {
	Path string
}

func NewFileCache(path string) *FileCache {
	return &FileCache{
		Path: path,
	}
}

func (f *FileCache) Set(data Expirable) error {
	bytes, err := json.Marshal(data)
	if err == nil {
		ioutil.WriteFile(f.Path, bytes, 0644)
	}
	return err
}

func (f *FileCache) Get(data Expirable) error {
	bytes, err := ioutil.ReadFile(f.Path)
	if err == nil {
		err = json.Unmarshal(bytes, data)
		if err == nil {
			created := data.CreatedAt()
			expires := data.ExpiresIn()
			if err == nil && time.Now().Unix() > created+int64(expires-60) {
				err = errors.New("Data is already expired")
			}
		}
	}
	return err
}

func sha1Sign(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func hmacSha256Sign(s, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}

// base编码
func base64EncodeStr(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// base解码
func base64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}

func HandJSONTopResponse(responseData interface{}, content []byte) {
	json.Unmarshal(content, responseData)
}

func HandXMLTopResponse(responseData interface{}, content []byte) {
	xml.Unmarshal(content, responseData)
}
