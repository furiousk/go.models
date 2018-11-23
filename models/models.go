package models

import (
	"log"
	"time"
)

type Mult struct {
	Name string
	Len  string
}

type Typer struct {
	Status int
	Key    string
}

type Evento struct {
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

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	ID        string `json:"_id"`
	Image     string `json:"image"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Matricula string `json:"matricula"`
	Token     string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func ReadMapInputs() (inputs map[string]map[int]string) {

	//inputs := map[string]map[int]string{}
	inputs["inputstatus"] = map[int]string{}
	inputs["extinputstt"] = map[int]string{}
	inputs["extinputst2"] = map[int]string{}
	inputs["outputstatus"] = map[int]string{}

	inputs["inputstatus"][15] = "ignicao"
	inputs["inputstatus"][14] = "panicBtn"
	inputs["inputstatus"][13] = "auxiliar4"
	inputs["inputstatus"][12] = "auxiliar1" //sensor de rpm analogico
	inputs["inputstatus"][11] = "antiFurto"
	inputs["inputstatus"][10] = "corteAntGps"
	inputs["inputstatus"][9] = "voltBateria"
	inputs["inputstatus"][8] = "temperatura"
	inputs["inputstatus"][7] = "antiFurtoLgc"
	inputs["inputstatus"][6] = "auxiliar2" //sensor de velociadade analogico
	inputs["inputstatus"][5] = "aceleracaoSts"
	inputs["inputstatus"][4] = "auxiliar3"
	inputs["inputstatus"][3] = "gsmSinalFraco"
	inputs["inputstatus"][2] = "backupBateriaFraco"
	inputs["inputstatus"][1] = "semBackupBateria"
	inputs["inputstatus"][0] = "trocarBateria"

	inputs["extinputstt"][15] = "violacaoPosit"
	inputs["extinputstt"][14] = "inclinacaoExcessiva"
	inputs["extinputstt"][13] = "marchaLenta"
	inputs["extinputstt"][12] = "falhaAlternador" //sensor de rpm analogico
	inputs["extinputstt"][11] = "deteccaoChuva"
	inputs["extinputstt"][10] = "aceleracaoBrusca"
	inputs["extinputstt"][9] = "freadaBrusca"
	inputs["extinputstt"][8] = "identMotoristaSts"
	inputs["extinputstt"][7] = "vdrEmparelhado"
	inputs["extinputstt"][6] = "reservado1" //sensor de velociadade analogico
	inputs["extinputstt"][5] = "reservado2"
	inputs["extinputstt"][4] = "reservado3"
	inputs["extinputstt"][3] = "reservado4"
	inputs["extinputstt"][2] = "reservado5"
	inputs["extinputstt"][1] = "reservado6"
	inputs["extinputstt"][0] = "reservado7"

	inputs["extinputst2"][15] = "jamming"
	inputs["extinputst2"][14] = "excessoVelSensor"
	inputs["extinputst2"][13] = "excessoRpmSensor"
	inputs["extinputst2"][12] = "banguela" //sensor de rpm analogico
	inputs["extinputst2"][11] = "notused1"
	inputs["extinputst2"][10] = "notused2"
	inputs["extinputst2"][9] = "notused3"
	inputs["extinputst2"][8] = "notused4"
	inputs["extinputst2"][7] = "notused5"
	inputs["extinputst2"][6] = "notused6" //sensor de velociadade analogico
	inputs["extinputst2"][5] = "notused7"
	inputs["extinputst2"][4] = "excessoVelCan"
	inputs["extinputst2"][3] = "excessoRotacaoCan"
	inputs["extinputst2"][2] = "banguelaCan"
	inputs["extinputst2"][1] = "excessoTempCan"
	inputs["extinputst2"][0] = "numeroSateliteZero"

	inputs["outputstatus"][15] = "bloqueio"
	inputs["outputstatus"][14] = "sirene"
	inputs["outputstatus"][13] = "auxiliar1"
	inputs["outputstatus"][12] = "reservado1" //sensor de rpm analogico
	inputs["outputstatus"][11] = "reservado2"
	inputs["outputstatus"][10] = "reservado3"
	inputs["outputstatus"][9] = "reservado4"
	inputs["outputstatus"][8] = "reservado5"
	inputs["outputstatus"][7] = "reservado6"
	inputs["outputstatus"][6] = "reservado7" //sensor de velociadade analogico
	inputs["outputstatus"][5] = "reservado8"
	inputs["outputstatus"][4] = "reservado9"
	inputs["outputstatus"][3] = "reservado10"
	inputs["outputstatus"][2] = "reservado11"
	inputs["outputstatus"][1] = "reservado12"
	inputs["outputstatus"][0] = "reservado13"

	log.Println(inputs)
	return
}

func ReadType() (typer map[int]Typer) {

	typer[1] = Typer{Status: 1, Key: "ignicao"}
	typer[14] = Typer{Status: 0, Key: "null"}
	typer[38] = Typer{Status: 2, Key: "ignicao"}
	typer[42] = Typer{Status: 1, Key: "excessoVelSensor"}
	typer[43] = Typer{Status: 1, Key: "excessoRpmSensor"}
	typer[45] = Typer{Status: 1, Key: "excessoVelCan"}
	typer[46] = Typer{Status: 1, Key: "excessoRotacaoCan"}
	typer[51] = Typer{Status: 1, Key: "aceleracaoBrusca"}
	typer[52] = Typer{Status: 1, Key: "marchaLenta"}
	typer[170] = Typer{Status: 1, Key: "excessoVelSensor"}
	typer[171] = Typer{Status: 1, Key: "excessoRpmSensor"}
	typer[173] = Typer{Status: 1, Key: "excessoVelCan"}
	typer[174] = Typer{Status: 1, Key: "excessoRotacaoCan"}
	typer[179] = Typer{Status: 1, Key: "freadaBrusca"}
	typer[180] = Typer{Status: 1, Key: "marchaLenta"}

	return
}
