### 1. Configuração do Docker (se necessário)

Se você precisar configurar o Docker para rodar o PostgreSQL:

- **Ajuste do arquivo `docker-compose.yml`:**
  - No arquivo `docker-compose.yml`, altere os parâmetros para as suas credenciais de acesso ao banco:
    ```yaml
    environment:
      - POSTGRES_USER=seu_usuario
      - POSTGRES_PASSWORD=sua_senha
      - POSTGRES_DB=crud  # NÃO ALTERE O NOME DO BANCO!
    ```

- **Alteração no arquivo `main.go`:**
  - Na linha 152 do arquivo `main.go`, ajuste as credenciais do banco de dados:
    ```go
    db, err = sql.Open("postgres", "postgres://seu_usuario:sua_senha@postgres/crud?sslmode=disable")
    ```

### 2. Executando o Projeto

1. **Para executar o projeto:**
   - Execute o comando abaixo para subir os containers e a aplicação:
     ```bash
     docker-compose up --build
     ```

2. **Para parar a execução:**
   - Para parar os containers e a execução do projeto, utilize:
     ```bash
     docker-compose down
     ```

3. **Para verificar se a aplicação está rodando corretamente:**
   - Utilize o comando para listar os containers ativos:
     ```bash
     docker ps -a
     ```

---

### 1. Criando o Banco de Dados e a Tabela

1. **Acesse o terminal do container do PostgreSQL:**
   - Entre no container utilizando o comando:
     ```bash
     docker container exec -it postgres_database bash
     ```

2. **Acesse o PostgreSQL:**
   - Utilize o comando `psql` para acessar o PostgreSQL:
     ```bash
     psql -U stacoviaki
     ```

3. **Crie o banco de dados:**
   - Execute o comando para criar o banco de dados:
     ```sql
     CREATE DATABASE crud;
     ```

4. **Conecte-se ao banco de dados `crud`:**
   - Entre no banco `crud`:
     ```sql
     \c crud
     ```

5. **Crie a tabela `livros`:**
   - Execute o comando para criar a tabela de livros:
     ```sql
     CREATE TABLE livros (
         id SERIAL PRIMARY KEY, 
         titulo VARCHAR(255) NOT NULL, 
         categoria VARCHAR(100) NOT NULL, 
         autor VARCHAR(255) NOT NULL, 
         sinopse TEXT NOT NULL
     );
     ```

6. **Verifique se a tabela foi criada:**
   - Para listar as tabelas do banco, execute:
     ```sql
     \dt
     ```

---

## 🛠 Como Testar a API

Você pode testar a API utilizando o **Postman** ou qualquer outra ferramenta para fazer requisições HTTP. Abaixo estão as informações para testar os endpoints.

### **1. Criar Livro (`CREATE`)**

- **Endpoint**: `POST http://localhost:8080/livros/create`
- **Body (JSON)**:
  ```json
  {
    "id": 1,
    "titulo": "O Senhor dos Anéis",
    "categoria": "Fantasia",
    "autor": "J.R.R. Tolkien",
    "sinopse": "Uma história épica sobre a luta contra o mal na Terra Média."
  }

### **2. Editar Livro (`UPDATE`)**

- **Endpoint**: `PUT http://localhost:8080/livros/update`
- **Body (JSON)**:
  ```json
    {
      "id": 1,
      "titulo": "O Senhor dos Fulanos",
      "categoria": "Fantasia",
      "autor": "J.R.R. Tolkien",
      "sinopse": "Uma história épica sobre a luta contra o mal na Terra Média."
    }

### **3. Listar Livro (`READ`)**

- **Endpoint**: `GET http://localhost:8080/livros/read`

### **4. Deletar Livro (`DELETE`)**

- **Endpoint**: `DELETE http://localhost:8080/livros/update?id=1`