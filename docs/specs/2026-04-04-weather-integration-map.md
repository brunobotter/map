# Integração de Clima no Map Data + Resiliência HTTP

## 1) Contexto
- Problema atual: o endpoint de mapa devolve clima mockado e sem integração externa.
- Impacto na operação da clínica: dashboards operacionais e visão territorial não refletem condição climática real.
- Stakeholders: time de produto, backend e operação.

## 2) Objetivo
- Resultado esperado de negócio: retornar clima real (temperatura + condição) no payload de mapa.
- Métrica de sucesso: campo `weather` preenchido com dados do OpenWeather em chamadas válidas.

## 3) Escopo
### Inclui
- Cliente HTTP com timeout de 5s e retry para falhas transitórias.
- Integração OpenWeather via camada de integração.
- Encadeamento Handler -> UseCase -> Service -> Integration/Repository.

### Não inclui
- Persistência de histórico climático em banco.
- Alteração de contrato de tráfego/eventos do mapa.

## 4) Regras de negócio
- O `MapService` deve buscar clima via `WeatherService` e dados complementares via `MapRepository`.
- O `WeatherService` deve delegar obtenção externa à camada de integração.
- A integração OpenWeather deve retornar `temperature` e `status` (condition).
- Em erro na integração, o mapa deve responder fallback de clima (`unknown`, `0`, `C`).
- Em erro de repositório, tráfego/eventos devem retornar vazios sem quebrar a resposta.

## 5) Fluxos principais
- Fluxo A: Handler chama UseCase; UseCase chama MapService; MapService chama WeatherService e MapRepository; WeatherService chama OpenWeatherIntegration; resposta de mapa contém clima real e dados locais de tráfego/eventos.
- Fluxo B (erro externo): falha em OpenWeather; MapService aplica fallback de clima e segue resposta com tráfego/eventos.

## 6) Contratos e dados
- Endpoints impactados:
  - `GET /map` (mesmo endpoint, preenchimento real de `weather`).
- Campos obrigatórios/validações:
  - `weather.temperature` (float)
  - `weather.status` (string)
  - `weather.unit` (`C`)
- Configuração:
  - `OPENWEATHER_API_KEY`
  - `OPENWEATHER_BASE_URL` (default `https://api.openweathermap.org`)

## 7) UX/UI (quando aplicável)
- Telas impactadas: visualização de mapa que consome payload `weather`.
- Estados de erro/sucesso: fallback controlado sem quebrar contrato de resposta.

## 8) Segurança e permissões
- Sem mudança de permissões.
- Chave de API deve vir por variável de ambiente.

## 9) Critérios de aceite
- [x] CARD 6: HTTP client com timeout de 5s e retry funcional.
- [x] CARD 7: `WeatherService.GetWeather(lat,lng)` integrado ao OpenWeather.
- [x] CARD 8: `MapService` preenche campo `weather` via `WeatherService`.
- [x] Timeout cancela request sem bloquear fluxo.

## 10) Plano de testes
- Unitários:
  - retry em 5xx no HTTP client.
  - timeout no HTTP client.
  - `WeatherService` delegando para integração.
  - `MapService` preenchendo weather e fallback.
- Integração/API:
  - integração OpenWeather parseando payload esperado.
- E2E/manual:
  - opcional neste card.

## 11) Rollout e observabilidade
- Feature flag: não aplicável.
- Logs/métricas: sem alteração estrutural neste card.
- Plano de rollback: reverter commit e retornar weather mockado.

## 12) Riscos e dependências
- Riscos: indisponibilidade da API externa e latência de rede.
- Dependências: OpenWeather API key válida.
