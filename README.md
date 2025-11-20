# Api Golang
<br>
📘 Meu Primeiro Projeto de API em Golamg

Este repositório contém meu primeiro projeto de estudo desenvolvendo uma API em Golang (Go). O objetivo principal é aprender os conceitos fundamentais da linguagem, boas práticas de desenvolvimento backend e estruturação de uma API REST.
<br><br>

🚀 Objetivo do Projeto<br>
Aprender a criar uma API utilizando Golang <br>
Entender rotas, handlers e responses <br>
Praticar organização de código e padrões como controllers, repository, usecase e models <br>
Utilizar módulos Go e dependências externas
<br>
Aprender como executar e documentar uma API simples.

<br>

🛠️ Tecnologias Utilizadas <br>
Golang 1.x <br>
Gin / net/http <br>
Go Modules <br>
JSON para comunicação <br>

Banco de dado <br>
(Opcional) SQLite, PostgreSQL, MongoDB <br>
(Opcional) Docker <br>

Nesse projeto utilizei o docoker e docker compose para colocar o postgresql em container e a api rondando.

<br><br>

📂 Estrutura do Projeto (Exemplo) <br>

/project<br>
│── cmd/<br>
│   └── main.go<br>
│── db/<br>
│   ├── conn.go<br>
│── controller/<br>
│   ├── productControler.go<br>
│── model/<br>
│   ├── productModel.go<br>
│   ├── responseModel.go<br>
│── repository/<br>
│   ├── productRepo.go<br>
│── usecase/<br>
│   ├── productUseCase.go<br>
│── go.mod<br>
│── go.sum<br>
└── README.md

<br><br>

▶️ Como Executar o Projeto <br>
Instale o Go em sua máquina https://go.dev/dl/

Clone o repositório: 
````
git clone https://github.com/seu-usuario/seu-repo.git
````

Acesse o diretório:
````
cd seu-repo
````

Baixe as dependências:
```
go mod tidy
```

Execute o servidor:
````
go run cmd/main.go
````

<br>

📌 Rotas da API (Exemplo)

Método	Rota	Descrição <br>
GET	 /ping	         Testa se a API está ativa<br>
GET	 /products	     Lista usuários<br>
POST /product	     Cria um novo produto

<br><br>

🧪 Testes

Caso você utilize testes:

go test ./...

<br>

🔮 Próximos Passos / Roadmap

 1. Implementar Testes unitários
 2. Adicionar testes automatizados
 3. Implementar autenticação bearer token
 4. Melhorar tratamento de erros

<br><br>

📖 Aprendizados

Esse projeto marca o início da minha jornada com Golang.

Aqui aprendi sobre:
<li>Servidor http</li>
<li>Manipulação de JSON</li>
<li>Organização de projeto</li>
<li>Módulos e dependências</li>
<li>Boas práticas de API REST</li>

<br><br>

Ronaldo Torre <br>
Senior Software Engineer - 2025