package config

import (
	"os"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

var SESSION_SECRET = os.Getenv("SESSION_SECRET")
var EVENT_START = os.Getenv("EVENT_START")
var EVENT_END = os.Getenv("EVENT_END")
var POST_EVENT = os.Getenv("POST_EVENT")
var OBJECT_STORE = os.Getenv("OBJECT_STORE") // s3 or native

var NFS_PATH = "/etc/phoros" // don't change unless you know what you're doing

var APP_PORT = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
