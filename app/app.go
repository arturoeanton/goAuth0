package app

import (
	"encoding/gob"

	"github.com/arturoeanton/goAuth0/config"
	"github.com/gorilla/sessions"
)

var (
	Store  *sessions.FilesystemStore
	Config config.Auth0Config
)

func Init(config config.Auth0Config, store string) error {
	Config = config
	Store = sessions.NewFilesystemStore("", []byte(store))
	gob.Register(map[string]interface{}{})
	return nil
}
