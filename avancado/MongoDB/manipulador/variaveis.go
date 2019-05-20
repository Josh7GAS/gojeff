package manipulador

import "html/template"

//ModeloOla armazenam os modelo de página ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazenam o modelo de pagina local
var ModeloLocal = template.Must(template.ParseFiles("html/ola.html"))
