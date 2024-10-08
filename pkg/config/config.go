package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"text/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	TLSenabled    bool
	Session       *scs.SessionManager
}
