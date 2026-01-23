package config

import "log"

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}
