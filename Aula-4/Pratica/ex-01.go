package main

type usuario struct {
	nome      string
	sobrenome string
	idade     uint
	email     string
	senha     string
}

type usuarios []*usuario

func (u *usuario) mudarNome(nome string, sobrenome string) {
	u.nome = nome
	u.sobrenome = sobrenome
}
func (u *usuario) mudarPassword(senha string) {
	u.senha = senha
}
func (u *usuario) mudarEmail(email string) {
	u.email = email
}

func main() {

	u := usuario{
		nome:      "João",
		sobrenome: "Silva",
		idade:     25,
		email:     "blablabla@.gmail.com",
		senha:     "123456",
	}

	u.mudarNome("Pedro", "Silva")
	u.mudarPassword("123456789")
	u.mudarEmail("legal@legal.com")

}
