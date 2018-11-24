package models

import (
	"time"
)

type (
	Mult struct {
		Name string
		Len  string
	}
	Typer struct {
		Status int
		Key    string
	}
	Evento struct {
		Serial      int
		Evento      int
		Ocorrencia  int
		Tempototal  int
		Tempoinicio time.Time
		Tempofinal  time.Time
		Gpsinicio   string
		Gpsfinal    string
		Telemetria  string
		Name        string
	}
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	JwtToken struct {
		ID        string `json:"_id"`
		Image     string `json:"image"`
		Nome      string `json:"nome"`
		Email     string `json:"email"`
		Matricula string `json:"matricula"`
		Token     string `json:"token"`
	}
	Environment struct {
		Appname      string `json:"appname"`
		Appenv       string `json:"appenv"`
		Appport      string `json:"appport"`
		Dbconnection string `json:"dbconnection"`
		Dbhost       string `json:"dbhost"`
		Dbport       string `json:"dbport"`
		Dbdatabase   string `json:"dbdatabase"`
		Dbusername   string `json:"dbusername"`
		Dbpassword   string `json:"dbpassword"`
	}
	Exception struct {
		Message string `json:"message"`
	}
)
