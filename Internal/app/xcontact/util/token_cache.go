package util

import (
	"ProjetoGustavo/Internal/app/xcontact/util/auth"
	"fmt"
	"sync"
	"time"
)

var (
	token    string
	expiraEm time.Time
	mu       sync.RWMutex
)

func GetToken() string {
	mu.RLock()
	if time.Now().Before(expiraEm) {
		defer mu.RUnlock()
		fmt.Printf("Existe token, token in-memory: %s", token)
		return token
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	token = auth.GetLogin()
	expiraEm = time.Now().Add(time.Hour)
	fmt.Printf("Não existia, estou renovando: %s", token)
	return token
}
