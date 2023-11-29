package config

import (
	"fmt"
	"strings"
)

// Config is a top-level Unit config
type Config struct {
	AccessLog    string                 `json:"access_log,omitempty"`
	Applications map[string]Application `json:"applications,omitempty"`
	Listeners    map[string]Listener    `json:"listeners,omitempty"`
}

// Application is a Unit application
type Application struct {
	Type       string   `json:"type,omitempty"`
	WorkDir    string   `json:"working_directory,omitempty"`
	Executable string   `json:"executable,omitempty"`
	User       string   `json:"user,omitempty"`
	Group      string   `json:"group,omitempty"`
	Args       []string `json:"arguments,omitempty"`
	Processes  int      `json:"processes,omitempty"`
	Path       string   `json:"path,omitempty"`
	Module     string   `json:"module,omitempty"`
	Stderr     string   `json:"stderr,omitempty"`
}

// Listener is a Unit listener
type Listener struct {
	Pass string `json:"pass,omitempty"`
}

// Summary returns a pretty-printable summary of the config
func (c Config) Summary() string {
	b := strings.Builder{}

	if len(c.Listeners) > 0 {
		b.WriteString("Listeners:\n")

		for addr := range c.Listeners {
			b.WriteString(fmt.Sprintf("\t%s\n", addr))
		}
	}

	if len(c.Applications) > 0 {
		b.WriteString("Applications:\n")

		for app := range c.Applications {
			b.WriteString(fmt.Sprintf("\t%s\n", app))
		}
	}

	return b.String()
}

// {
// 	"access_log": "/var/log/unit/access.log",
// 	"applications": {
// 	  "nodejsapp": {
// 		"type": "external",
// 		"working_directory": "/www/app/node-app/",
// 		"executable": "app.js",
// 		"user": "www",
// 		"group": "www",
// 		"arguments": [
// 		  "--tmp-files",
// 		  "/tmp/node-cache"
// 		]
// 	  },
// 	  "pythonapp": {
// 		"type": "python 3.11",
// 		"processes": 16,
// 		"working_directory": "/www/app/python-app/",
// 		"path": "blog",
// 		"module": "blog.wsgi",
// 		"user": "www",
// 		"group": "www",
// 		"stderr": "stderr.log",
// 		"isolation": {
// 		  "rootfs": "/www/"
// 		}
// 	  }
// 	},
// 	"routes": {
// 	  "local": [
// 		{
// 		  "action": {
// 			"share": "/www/local/"
// 		  }
// 		}
// 	  ],
// 	  "global": [
// 		{
// 		  "match": {
// 			"host": "backend.example.com"
// 		  },
// 		  "action": {
// 			"pass": "applications/pythonapp"
// 		  }
// 		},
// 		{
// 		  "action": {
// 			"pass": "applications/nodejsapp"
// 		  }
// 		}
// 	  ]
// 	},
// 	"listeners": {
// 	  "127.0.0.1:8080": {
// 		"pass": "routes/local"
// 	  },
// 	  "*:443": {
// 		"pass": "routes/global",
// 		"tls": {
// 		  "certificate": "bundle",
// 		  "conf_commands": {
// 			"ciphersuites": "TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256",
// 			"minprotocol": "TLSv1.3"
// 		  },
// 		  "session": {
// 			"cache_size": 10240,
// 			"timeout": 60,
// 			"tickets": [
// 			  "IAMkP16P8OBuqsijSDGKTpmxrzfFNPP4EdRovXH2mqstXsodPC6MqIce5NlMzHLP",
// 			  "Ax4bv/JvMWoQG+BfH0feeM9Qb32wSaVVKOj1+1hmyU8ORMPHnf3Tio8gLkqm2ifC"
// 			]
// 		  }
// 		},
// 		"forwarded": {
// 		  "client_ip": "X-Forwarded-For",
// 		  "recursive": false,
// 		  "source": [
// 			"192.0.2.0/24",
// 			"198.51.100.0/24"
// 		  ]
// 		}
// 	  }
// 	},
// 	"settings": {
// 	  "http": {
// 		"body_read_timeout": 30,
// 		"discard_unsafe_fields": true,
// 		"header_read_timeout": 30,
// 		"idle_timeout": 180,
// 		"log_route": true,
// 		"max_body_size": 8388608,
// 		"send_timeout": 30,
// 		"server_version": false
// 	  }
// 	}
//   }
