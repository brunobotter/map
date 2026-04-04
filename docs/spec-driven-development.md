# Spec-Driven Development (SDD) — Clínica Odontológica

Este documento define como evoluir o projeto usando **Spec-Driven Development**, onde a especificação vem antes da implementação e guia todo o ciclo de entrega.

## Objetivos

- Reduzir retrabalho e ambiguidades antes de codificar.
- Garantir alinhamento entre negócio , produto e engenharia.
- Aumentar previsibilidade das entregas e facilitar testes/aceite.
- Melhorar rastreabilidade entre requisito, código, testes e documentação.

## Princípios

1. **Spec primeiro, código depois**.
2. **Escopo pequeno e verificável** por incremento.
3. **Critérios de aceite objetivos** (Given/When/Then ou checklist testável).
4. **Rastreabilidade obrigatória** entre Spec → Tasks → PR → Testes.
5. **Evolução contínua da spec**: mudanças devem atualizar a especificação.

## Estrutura padrão de uma Spec

Cada nova funcionalidade deve gerar um arquivo em `docs/specs/` com o nome:

`YYYY-MM-DD-nome-da-feature.md`

Template recomendado:

```md
# [Título da Feature]

## 1) Contexto
- Problema atual:
- Impacto na operação da clínica:
- Stakeholders:

## 2) Objetivo
- Resultado esperado de negócio:
- Métrica de sucesso:

## 3) Escopo
### Inclui
- ...

### Não inclui
- ...

## 4) Regras de negócio
- Regra 1
- Regra 2

## 5) Fluxos principais
- Fluxo A
- Fluxo B

## 6) Contratos e dados
- Endpoints impactados:
- Tabelas/migrações:
- Campos obrigatórios/validações:

## 7) UX/UI (quando aplicável)
- Telas impactadas:
- Estados de erro/sucesso:

## 8) Segurança e permissões
- Perfis que podem acessar:
- Restrições:

## 9) Critérios de aceite
- [ ] Critério 1
- [ ] Critério 2

## 10) Plano de testes
- Unitários:
- Integração/API:
- E2E/manual:

## 11) Rollout e observabilidade
- Feature flag (se necessário):
- Logs/métricas:
- Plano de rollback:

## 12) Riscos e dependências
- Riscos:
- Dependências:
```

## Fluxo de trabalho SDD (padrão do projeto)

### Etapa 1 — Descoberta

- Entender problema com contexto clínico e operacional.
- Validar impacto em módulos existentes (`app/`, `web/`, `migrations/`).
- Escrever a primeira versão da spec com foco em clareza.

### Etapa 2 — Refinamento técnico

- Mapear contratos de API (OpenAPI + handlers/usecases/services/repositórios).
- Mapear alterações de banco (migrations versionadas).
- Definir estratégia de testes antes da implementação.

### Etapa 3 — Quebra em tarefas

Para cada spec, criar um plano em issues/checklist com:

- Backend (domínio, usecase, handler, repo, migração).
- Frontend (queries, páginas/componentes, estados de UI).
- Testes (unitários, integração, regressão).

### Etapa 4 — Implementação guiada pela spec

- Implementar incrementalmente.
- Validar critérios de aceite a cada commit/PR.
- Não introduzir escopo fora da spec sem atualização explícita.

### Etapa 5 — Validação

- Executar testes previstos na spec.
- Validar fluxo principal de ponta a ponta.
- Registrar evidências no PR (checklist + observações).

### Etapa 6 — Entrega e aprendizado

- Atualizar spec para estado final (`as-built`), quando necessário.
- Registrar decisões arquiteturais relevantes nos docs.
- Capturar melhorias para próximas iterações.

## DoR e DoD para SDD

### Definition of Ready (DoR)

Uma tarefa só entra em desenvolvimento quando:

- [ ] Spec escrita e revisada.
- [ ] Critérios de aceite mensuráveis definidos.
- [ ] Escopo “inclui / não inclui” claro.
- [ ] Dependências identificadas.

### Definition of Done (DoD)

Uma tarefa só é concluída quando:

- [ ] Implementação atende a todos critérios de aceite.
- [ ] Testes previstos foram executados.
- [ ] OpenAPI/documentação atualizados (se aplicável).
- [ ] Migrações e rollback descritos (se aplicável).
- [ ] PR com rastreabilidade para a spec.

## Checklist de PR orientado por Spec

Incluir no PR:

- Link da Spec principal.
- Lista de critérios de aceite e status.
- Evidências de testes executados.
- Riscos conhecidos e plano de mitigação.
- Itens fora de escopo (se surgiram).

## Exemplo curto de critério de aceite (Gherkin)

```gherkin
Cenário: Registrar pagamento em atendimento
  Dado que existe um atendimento em aberto
  E o usuário possui permissão financeira
  Quando ele registra um pagamento válido
  Então o sistema atualiza o saldo do atendimento
  E registra a transação financeira relacionada
```

