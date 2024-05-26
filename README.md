# Publish Exception
Publish Exception é um sistema que centraliza a publicação de exceções de diferentes aplicações para diferentes canais de notificação. O objetivo do projeto é simplificar o processo de notificação de exceções e fornecer uma visão centralizada de todas as exceções que ocorrem no sistema.

## Componentes do sistema

Publisher: São responsáveis publicar as mensagens de exceção das aplicações para o tópico de exceções dentro de um container Docker
Tópico de exceções: O tópico de exceções é um tópico do Kafka que recebe as mensagens de exceção das aplicações
Publish Exception Manager: O Publish Exception Manager é um serviço responsável por consumir as mensagens de exceção do tópico de exceções e publicar as exceções para os canais de notificação

## Funcionalidades do sistema

Publicação centralizada de exceções para diferentes canais de notificação.
Visão centralizada de todas as exceções que ocorrem no sistema.
Filtragem e pesquisa de exceções.
Notificação de exceções por e-mail e Slack.

## Pré-requisitos
Para executar o projeto, você precisará ter os seguintes softwares instalados:

Kafka: A versão 2.8 ou superior é recomendada.
Docker: A versão 24 ou superior é recomendada.

[EM DESENVOLVIMENTO]

## Como executar o projeto
[EM DESENVOLVIMENTO]

## Publicar as mensagens de exceção com o Publisher
[EM DESENVOLVIMENTO]

## Configuração no gerenciador
Como configurar no gerenciador como receber e notificar as exceções
### Tópico
[EM DESENVOLVIMENTO]
### Alertas
[EM DESENVOLVIMENTO]
### Conectores
[EM DESENVOLVIMENTO]
