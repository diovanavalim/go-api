<h1 align="center">API Devbook</h1>

<p align="center">Repositório criado para hospedar o código utilizado no módulo de API do curso de Go da Udemy.</p>

<p align="center">
  <img align="" alt="gopher" src="https://webartdevelopers.com/blog/wp-content/uploads/2021/11/gopher-golang-css-only-animation.gif" height="70%" width="70%" />
</p>

<p align="center">
 <a href="#Objetivo">Sobre o Projeto</a> •
 <a href="#Tecnologias">Tecnologias</a> • 
 <a href="#Pré-requisitos">Pré-requisitos</a> • 
 <a href="#Rodando o Servidor">Rodando o Servidor</a> • 
 <a href="#Dependências">Dependências</a> •
 <a href="#Autora">Autora</a> •
</p>

<a name="Objetivo"></a>
## 🖊 Sobre o Projeto

<p> 
O objetivo do projeto é desenvolver uma API Rest, escrita na linguagem Go. Esta API implementa uma rede social. Neste contexto, é possível cadastrar usuários, seguir e dar unfollow em um usuário, bem como realizar publicações, com as quais o usuário pode interagir deixando seu like.
</p>

___
<a name="Tecnologias"></a>

## 🛠 Tecnologias
A API foi construída utilizando a linguagem Go, com a biblioteca nativa http.

___
<a name="Pré-requisitos"></a>

## ✅ Pré-requisitos
Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Git](https://git-scm.com), [Go](https://go.dev/), [MySQL](https://www.mysql.com/) e um editor de código a sua escolha.

___

<a name="Rodando o Servidor"></a>

## 🎲 Rodando o Servidor


- Clone este repositório na sua máquina:
    
      git clone https://github.com/diovanavalim/go-api.git

- Acesse a pasta do projeto no terminal:

      cd go-api

- Instale as dependências do projeto:

      go mod downloads

- Preencha as variáveis de ambiente do arquivo `env.dist`

- Execute o arquivo SQL `init.sql`, da pasta `sql`
  
- Execute o servidor:

      go run main.go

O servidor estará sendo escutado na porta indicada na IDE. Geralmente, trata-se da porta 8686.

___
<a name="Dependências"></a>
## 🏁 Dependências  

Foram utilizadas no projeto as seguintes dependências:

- 📧 Checkmail;

- 🔑 JWT Go;

- 🍃 Mux;

- 🌐 Godotenv;

- 🤐 Crypto

___

<a name="Autora"></a>

## 📝 Autores ##

Desenvolvido com 💛 por Diovana Rodrigues Valim, estudante da oitava fase de Sistemas de Informção na UFSC e desenvolvedora de software no Mercado Livre.