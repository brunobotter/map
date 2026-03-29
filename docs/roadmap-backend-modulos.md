# Roadmap do Back-end por Módulos (início → fim)

## Objetivo
Padronizar o back-end no fluxo **Controller → UseCase → Service → Repository/Integration**, com contratos por interface, segurança de payload e testes por camada, sem big bang.

---

## Ordem de execução recomendada

1. **Fase 0 — Preparação da base (Sprint 1)**
2. **Fase 1 — Piloto vertical (Sprint 2)**
3. **Fase 2 — Módulos core (Sprints 3 a 7)**
4. **Fase 3 — Módulos satélite (Sprints 8 a 10)**
5. **Fase 4 — Consolidação e descomissionamento legado (Sprint 11+)**

---

## Fase 0 — Preparação (por onde começar)

### Entregas
- Definir template de pastas por módulo:
  - `application/usecase/<modulo>`
  - `application/service/<modulo>`
  - `infra/repository/<modulo>`
  - `infra/integration/<modulo>`
- Criar contratos transversais:
  - `TxManager/UnitOfWork`, `Clock`, `Authorizer`, `IDGenerator`
  - `DomainError` com códigos estáveis (`not_found`, `conflict`, `validation_error`, `forbidden`)
  - `ResponsePolicy` + `DataClassifier`
- Padronizar tratamento de erro HTTP.
- Ajustar DI/container para resolver por interface.
- Criar matriz de exposição de dados por endpoint (`docs/matriz-inicial-exposicao-endpoints.md`).

### Critério de pronto
- 1 endpoint já migrado ponta-a-ponta no padrão novo, com testes de referência.

---

## Fase 1 — Piloto vertical

### Módulo piloto (sugestão)
- **Specialty** ou **PatientOrigin** (baixo risco + alto uso).

### Passos
1. Controller enxuto (somente input/output e erro HTTP).
2. UseCase explícito por cenário.
3. Service com regra de domínio.
4. Repository concreto por interface.
5. Testes por camada:
   - controller
   - usecase
   - service
   - repository (integração/contrato)

### Critério de pronto
- Endpoint legado removido nesse fluxo.
- Contrato e latência preservados.

---

## Fase 2 — Módulos core (ordem recomendada)

### 2.1 Auth + Users
**Por que primeiro:** segurança, autenticação e autorização habilitam os demais módulos.

**Escopo mínimo:**
- login/refresh/logout
- perfis/roles/escopos
- política de visibilidade de campos por perfil

### 2.2 Patients + Prontuário
**Por que agora:** alto impacto de negócio e alto risco de exposição de PII.

**Escopo mínimo:**
- cadastro/listagem/detalhe
- prontuário com DTOs segregados (lista vs detalhe)
- mascaramento e trilha de auditoria

### 2.3 Appointments (Agenda)
**Escopo mínimo:**
- criar/reagendar/cancelar/alterar status
- regras de conflito e antecedência
- controle transacional para concorrência

### 2.4 Budgets + Treatments
**Escopo mínimo:**
- criação de orçamento e itens
- aprovação/reprovação (idempotente)
- vínculo com tratamento/plano

### 2.5 Financeiro (cash/AR/AP/transactions)
**Por que por último nos core:** maior complexidade transacional e risco de regressão.

**Escopo mínimo:**
- caixa (abertura/fechamento)
- contas a receber/pagar
- conciliação e trilhas de auditoria

### Regra de execução para cada módulo core
1. Inventariar endpoints e payload atual.
2. Classificar campos (`public/internal/sensitive/restricted`).
3. Extrair casos de uso e interfaces.
4. Migrar endpoint a endpoint com feature flag.
5. Validar contrato HTTP + segurança de resposta.
6. Remover legado após estabilização curta.

---

## Fase 3 — Módulos satélite

## Ordem sugerida
1. CRM
2. Relatórios
3. Monitoramento
4. Inventory / Laboratório / Convênios
5. Configurações administrativas

**Objetivo:** reaproveitar padrões estabilizados dos módulos core.

---

## Fase 4 — Consolidação (até onde terminar)

### Entregas finais
- Remover paths/handlers/serviços legados duplicados.
- Fechar gaps de observabilidade:
  - logs estruturados com `request_id`, `user_id`, `use_case`, `duration_ms`
  - métricas por caso de uso (p95, erro, throughput)
- Revisar SLO e capacidade.
- Hardening AWS concluído:
  - CORS restrito
  - TLS obrigatório
  - WAF/rate limiting
  - segredos em Secrets Manager/Parameter Store
  - logs sem PII sensível
  - auditoria de acesso em dados clínicos/financeiros

### Critério de encerramento
- 100% dos endpoints ativos no padrão novo.
- 0 controller acessando repo/SQL diretamente.
- Matriz de exposição aprovada para endpoints críticos.
- OpenAPI atualizado e testes de contrato verdes.

---

## Cronograma enxuto (11+ sprints)

- **Sprint 1:** Fase 0 (base + contratos + endpoint piloto base)
- **Sprint 2:** Fase 1 (piloto vertical completo)
- **Sprints 3–7:** Fase 2 (módulos core)
- **Sprints 8–10:** Fase 3 (satélites)
- **Sprint 11+:** Fase 4 (consolidação final)

---

## Definition of Done por endpoint

1. Controller sem regra e sem SQL/repo direto.
2. UseCase testado (fluxo feliz + erros principais).
3. Service testado (regras de domínio).
4. Repository com teste de integração/contrato.
5. `DomainError` padronizado e mapeado para HTTP.
6. DTO de saída sem exposição indevida.
7. Teste de contrato garantindo ausência de campos proibidos.
8. OpenAPI atualizado.
9. Logs e métricas mínimos instrumentados.
