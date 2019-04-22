package manipulador

import "html/template"

//modelos armazenam os modelos de pagina que seram executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
