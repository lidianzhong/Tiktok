package config

import "os"

var IP = "https://" + os.Getenv("paas_url")
