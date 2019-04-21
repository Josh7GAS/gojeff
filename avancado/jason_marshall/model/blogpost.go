package model

//BlogPost armazena dados de um post no Blog
type BlogPost struct {
	UsuarioID int    `json:"userId"`
	ID        int    `jason:"id"`
	Titulo    string `jason:"id"`
	Texto     string `jason:"body"`
}
