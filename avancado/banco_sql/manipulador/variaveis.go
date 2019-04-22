package manipulador

import "html/template"

//modeloOla armazenam os modelo de página ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloOlaLocal armazenam o modelo de pagina local
var ModeloOlaLocal = template.Must(template.ParseFiles("html/ola.html"))
