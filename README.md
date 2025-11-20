#Api Golang
<br>
📘 Meu Primeiro Projeto de API em Golamg

Este repositório contém meu primeiro projeto de estudo desenvolvendo uma API em Golang (Go). O objetivo principal é aprender os conceitos fundamentais da linguagem, boas práticas de desenvolvimento backend e estruturação de uma API REST.
<br><br>

🚀 Objetivo do Projeto

Aprender a criar uma API utilizando Golang

Entender rotas, handlers e responses

Praticar organização de código e padrões como controllers, services e models

Utilizar módulos Go e dependências externas

Aprender como executar, testar e documentar uma API simples

🛠️ Tecnologias Utilizadas

Golang 1.x

Gin / Fiber / net/http (coloque aqui o framework utilizado)

Go Modules

JSON para comunicação

(Opcional) Banco de dados: SQLite, PostgreSQL, MongoDB

(Opcional) Docker

Nesse projeto utilizei o docoker e docker compose para colocar o postgresql em container e a api rondando.


📂 Estrutura do Projeto (Exemplo)
/project<br>
│── cmd/<br>
│   └── main.go<br>
│── internal/<br>
│   ├── controllers/<br>
│   ├── services/<br>
│   ├── models/<br>
│   └── routes/<br>
│── go.mod<br>
│── go.sum<br>
└── README.md

<br><br>

▶️ Como Executar o Projeto

Instale o Go em sua máquina
https://go.dev/dl/

Clone o repositório:

git clone https://github.com/seu-usuario/seu-repo.git


Acesse o diretório:

cd seu-repo


Baixe as dependências:

go mod tidy


Execute o servidor:

go run cmd/main.go

<br><br>

📌 Rotas da API (Exemplo)
Método	Rota	Descrição
GET	 /ping	    Testa se a API está ativa
GET	 /products	Lista usuários
POST /product	Cria um novo produto

<br><br>

🧪 Testes

Caso você utilize testes:

go test ./...

🔮 Próximos Passos / Roadmap

 Implementar autenticação JWT

 Criar conexão com banco de dados real

 Melhorar tratamento de erros

 Adicionar testes automatizados

 Containerizar com Docker


<br><br>

📖 Aprendizados

Esse projeto marca o início da minha jornada com Go.
Aqui aprendi sobre:

Servidores HTTP

Manipulação de JSON

Organização de projeto

Módulos e dependências

Boas práticas de API REST
