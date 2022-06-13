package models

import "github.com/luizclaudioholanda/loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {

	db := db.ConectaBancoDeDados()

	result, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for result.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := result.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriaProduto(nome, descricao string, quantidade int, preco float64) {

	db := db.ConectaBancoDeDados()

	insereProduto, err := db.Prepare("INSERT INTO produtos (nome, descricao, quantidade, preco) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereProduto.Exec(nome, descricao, quantidade, preco)

	defer db.Close()

}

func DeleteProduto(idProduto string) {

	db := db.ConectaBancoDeDados()

	deleteProduto, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(idProduto)

	defer db.Close()

}

func BuscaProdutoPorId(idProduto string) Produto {

	db := db.ConectaBancoDeDados()

	buscaProduto, err := db.Query("SELECT * FROM produtos WHERE id = $1", idProduto)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for buscaProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = buscaProduto.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Preco = preco

	}

	defer db.Close()

	return produto

}

func UpdateProduto(id int, nome, descricao string, quantidade int, preco float64) {

	db := db.ConectaBancoDeDados()

	updateProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, quantidade=$3, preco=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, quantidade, preco, id)

	defer db.Close()

}
