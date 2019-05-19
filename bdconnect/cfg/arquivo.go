package cfg

import (
	"os"

	"gopkg.in/ini.v1"
)

//Credencial é uma array que a funcao cfgArquvio devolve
var Credencial = cfgArquivo()

func cfgArquivo() []string {
	var (
		//Credencial é uma array que a função devolves
		credencial []string
		host       string
		port       string
		user       string
		password   string
	)
	arquivo, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")
	if arquivo != nil {

		host = arquivo.Section("client").Key("host").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}

			return in
		})
		credencial = append(credencial, host)

		port = arquivo.Section("client").Key("port").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, port)

		user = arquivo.Section("client").Key("user").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, user)

		password = arquivo.Section("client").Key("password").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, password)

	}

	return credencial
}
