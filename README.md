# Publish Exception
Publish Exception é um sistema que centraliza a publicação de exceções de diferentes aplicações para diferentes canais de notificação. O objetivo do projeto é simplificar o processo de notificação de exceções e fornecer uma visão centralizada de todas as exceções que ocorrem no sistema.

![Image](/publish-exception.svg)

## Componentes do sistema

### Publisher
São responsáveis publicar as mensagens de exceção das aplicações para o tópico de exceções dentro de um container Docker
### Subjects de exceções
O subject de exceção recebe as mensagens de exceção da aplicação
### Publish Exception Manager
O Publish Exception Manager é um serviço responsável por consumir as mensagens de exceção dos subjects de exceções e publicar as exceções para os canais de notificação

## Funcionalidades do sistema

Publicação centralizada de exceções para diferentes canais de notificação.
Visão centralizada de todas as exceções que ocorrem no sistema.
Filtragem e pesquisa de exceções.
Notificação de exceções por e-mail e Slack.

## Como executar o projeto
Instruções sobre como executar o projeto.

### Instalar Nats
1. Baixar e Instalar o NATS Server
Baixe a versão mais recente do NATS Server a partir do [site oficial do NATS](https://nats.io/). Use wget ou curl para baixar o arquivo tar.gz e extrair o binário:
~~~sh
wget https://github.com/nats-io/nats-server/releases/download/v2.9.20/nats-server-v2.9.20-linux-amd64.tar.gz
tar -xzf nats-server-v2.9.20-linux-amd64.tar.gz
cd nats-server-v2.9.20-linux-amd64
~~~

2. Iniciar o Servidor NATS com JetStream Habilitado
Execute o servidor NATS com JetStream habilitado. Use a opção -js para ativar o JetStream:
~~~sh
./nats-server -js
~~~
O servidor será iniciado e escutará na URL padrão `nats://localhost:4222`.

Para mais detalhes, consulte a [documentação oficial sobre a configuração do NATS Server](https://docs.nats.io/nats-concepts/jetstream/).

3. Instalar o natscli (CLI para NATS)
Baixe e instale a ferramenta de linha de comando natscli:
~~~sh
wget https://github.com/nats-io/natscli/releases/download/v0.16.1/nats-0.16.1-linux-amd64.tar.gz
tar -xzf nats-0.16.1-linux-amd64.tar.gz
sudo mv nats /usr/local/bin/
~~~

Verifique a instalação:
~~~sh
nats --version
~~~
Para mais informações sobre o natscli, consulte a [documentação oficial](https://docs.nats.io/running-a-nats-service/configuration/resource_management/configuration_mgmt/nats-admin-cli).

### Criar Stream
1. Criar um Stream com natscli
Use o natscli para criar um stream chamado publish-exception. O comando nats stream add permite que você defina os detalhes do stream, como assuntos e opções de armazenamento:
~~~sh
nats stream add publish-exception --subjects "<my-app1>-publish-exception;<my-app2>-publish-exception" --storage file
~~~
Este comando cria um stream chamado publish-exception que inclui os assuntos test-app-publish-exception e test-app1-publish-exception e usa armazenamento em arquivo.

Para mais detalhes, consulte a [documentação oficial sobre streams](https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/streams).

2. Verificar a Configuração do Stream
Após criar o stream, você pode verificar sua configuração com o seguinte comando:
~~~sh
nats stream info publish-exception
~~~
Este comando exibe detalhes sobre o stream, incluindo os assuntos e a configuração atual.

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
