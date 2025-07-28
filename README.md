<h1>Projeto Wonit</h1>
Este projeto consiste em uma aplicação de integração com o sistema XContact da Wonit, desenvolvida como parte de um desafio prático após a certificação em Golang. A aplicação realiza o cadastro, edição e exclusão de ramais, além de gerenciar a autenticação utilizando cache de token. Os endpoints estão documentados e disponíveis via Swagger.

## Como executar:

<h4>1. Clone o repositório com:</h4>
- git clone https://github.com/GustavoDeolaWonit/ProjetoGustavoWonit.git
  
<h4>2. Acesse a pasta do projeto:</h4>
- cd ProjetoGustavoWonit

<h4>3. Execute o projeto:</h4>
- go run main.go
OBS: A aplicação estará disponível na porta configurada (por padrão, `localhost:8080`).

<h4>4. Acesse a documentação Swagger:</h4>
- http://localhost:8080/swagger/index.html

## Funcionalidades

- ✅ Cadastro de ramais
- ✏️ Edição de ramais
- ❌ Remoção de ramais
- 🔐 Gerenciamento de autenticação com cache de token
- 📄 Documentação dos endpoints via Swagger

## Estrutura do Projeto

- `controller/` → Define os handlers HTTP dos endpoints
- `service/` → Contém a lógica de negócio da aplicação
- `repositories/` → Responsável por chamadas HTTP externas e persistência
- `model/` → Define os modelos de entidade
- `dto/` → Define os modelos de requisição e resposta
- `util/` → Gerenciamento de autenticação e cache de token
- `main.go` → Inicialização da aplicação e rotas

  ## Autenticação

A autenticação com a API XContact é feita via login de supervisor, que retorna um token JWT.
Esse token é armazenado em memória e reutilizado por até 1 hora, evitando múltiplas autenticações desnecessárias.

A função GetToken() gerencia esse processo automaticamente:

Se o token ainda estiver válido, ele é reutilizado.

Se estiver expirado ou ausente, é feito um novo login via POST /api/v4/login/supervisor.

A sincronização de acesso ao token é feita com sync.RWMutex para garantir segurança em ambientes concorrentes.
