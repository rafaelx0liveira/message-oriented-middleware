# Message-Oriented-Middleware em Go

## Visão Geral

Projeto acadêmico para a matéria Sistemas Distribuídos e Computação Paralela administrada pelo professor Carlos Paes da universidade PUC-SPd de um middleware orientado a mensagens desenvolvido em Go (Golang), que visa fornecer uma infraestrutura para a troca eficiente de mensagens entre componentes de software em um sistema distribuído. O middleware é projetado para permitir a comunicação entre produtores e consumidores de mensagens por meio de um Broker de mensagens central.

## Estrutura de Pastas

O projeto está organizado em uma estrutura de pastas que inclui:

- `cmd`: Aplicativos principais para os componentes do sistema, como `consumer_app`, `message_server_app` e `producer_app`.

- `internal`: Pacotes internos que representam os principais componentes do sistema, incluindo `consumer`, `message_server` e `producer`.

- `api`: Uma pasta adicional contendo o projeto da API que permite a interação com o Broker de mensagens.
  
## Funcionalidades

- O `consumer` é responsável por receber mensagens do Broker e processá-las.

- O `producer` envia mensagens para o Broker de mensagens.

- O `message_server` atua como um Broker de mensagens, gerenciando filas e garantindo a entrega de mensagens aos consumidores apropriados.

- A `api` fornece endpoints para enviar e receber mensagens no Broker de mensagens.

- `Broker`: Componente central do sistema que gerencia o roteamento de mensagens entre produtores e consumidores, garantindo a confiabilidade e a escalabilidade na troca de mensagens.

## Arquitetura

<img src="docs/message-oriented-middleware architecture.png" alt="architecture image">


