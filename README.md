<h1 align="center">
  <p align="center">Books Loan</p>
</h1>


Serviço para emprêstimos de livros utilizado GraphQL, escrito em GoLang e utilizando PostgreSQL.

### Especificações 
#### Bibliotecas:
- [fiber](github.com/gofiber/fiber)
- [gqlgen](github.com/99designs/gqlgen)
- [gorm](https://github.com/go-gorm/gorm)
- [logrus](github.com/sirupsen/logrus)
- [govalidator](github.com/asaskevich/govalidator)
- [gMediator](github.com/VitorEmanoel/gMediator)
- [menv](github.com/VitorEmanoel/menv)

#### Bibliotecas para testes:
- [testify](github.com/stretchr/testify)
- [go-sqlmock](github.com/DATA-DOG/go-sqlmock)
- [apitest](github.com/steinfletcher/apitest)

#### Estrutura de dados
![alt data-structure](./book_loans.png)

### Instalação

#### Docker-compose
Utilizar Docker Compose CLI para subir serviços.
(Precisa das portas 8080 e 5432 disponível)
```shell
docker-compose up -d
```


### Testes
Para executar os testes utilizar o comando
```shell
go test ./...
```
