<h1>Projeto Wonit</h1>
Este projeto consiste em uma aplicação de integração com o sistema XContact da Wonit, desenvolvida como parte de um desafio prático após a certificação em Golang. A aplicação realiza o cadastro, edição e exclusão de ramais, além de gerenciar a autenticação utilizando cache de token. Os endpoints estão documentados e disponíveis via Swagger.

<h1>Como executar:</h1>

<h3>1. Clone o repositório com:</h3>
- git clone https://github.com/GustavoDeolaWonit/ProjetoGustavoWonit.git
  
<h3>2. Acesse a pasta do projeto:</h3>
- cd ProjetoGustavoWonit

<h3>3. Execute o projeto:</h3>
- go run main.go
OBS: A aplicação estará disponível na porta configurada (por padrão, `localhost:8080`).

<h3>4. Acesse a documentação Swagger:</h3>
- http://localhost:8080/swagger/index.html

## Funcionalidades

- ✅ Cadastro de ramais
- ✏️ Edição de ramais
- ❌ Remoção de ramais
- 🔐 Gerenciamento de autenticação com cache de token
- 📄 Documentação dos endpoints via Swagger
