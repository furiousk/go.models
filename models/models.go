package models

import (
	"time"
)

type (
	//Mult ....
	Mult struct {
		Name string
		Len  string
	}
	//Typer ....
	Typer struct {
		Status int
		Key    string
	}
	//Evento ....
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
	//Login ....
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//JwtToken ....
	JwtToken struct {
		ID        string `json:"_id"`
		Image     string `json:"image"`
		Nome      string `json:"nome"`
		Email     string `json:"email"`
		Matricula string `json:"matricula"`
		Token     string `json:"token"`
	}
	//Environment ....
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
		Exturl       string `json:"exturl"`
		Exttoken     string `json:"exttoken"`
	}
	//Exception ....
	Exception struct {
		Message string `json:"message"`
	}
)
