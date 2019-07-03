package configs

import (
	"github.com/gorilla/securecookie"
	"github.com/pelletier/go-toml"
)

var ConfigTree *toml.Tree

var (
	// AES仅支持16,24或32字节的密钥大小。
	//您需要准确提供该密钥字节大小，或者从您键入的内容中获取密钥。
	hashKey  = []byte("local test set hash cookie key")
	blockKey = []byte("local test set block cookie key")
	Sc       = securecookie.New(hashKey, blockKey)
)