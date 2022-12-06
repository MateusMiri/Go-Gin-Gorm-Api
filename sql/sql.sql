CREATE TABLE IF NOT EXISTS pessoas (
    id serial PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    cpf BIGINT UNIQUE NOT NULL,
    nascimento DATE NOT NULL,
    telefone BIGINT,
    email VARCHAR(80) UNIQUE,
    rua VARCHAR(255) NOT NULL,
    numero INT NOT NULL,
    bairro VARCHAR(255) NOT NULL,
    complemento VARCHAR(255),
    cidade VARCHAR(255) NOT NULL,
    uf VARCHAR(2) NOT NULL,
    cep BIGINT NOT NULL
);