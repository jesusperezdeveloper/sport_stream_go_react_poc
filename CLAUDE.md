# Sport Stream Go+React POC

> Generado con SpecBox Engine v5.17.0

## Stack

| Tecnologia | Version |
|------------|---------|
| Go | 1.23+ |
| React | 19.x |
| Arquitectura | Clean Architecture (Go backend) + React SPA frontend |

## Arquitectura

Proof of concept para streaming deportivo con backend Go y frontend React.
Arquitectura backend en Go siguiendo clean architecture (cmd/, internal/, pkg/).
Frontend React como SPA separada.

Ver patrones detallados en: specbox-engine/architecture/go/

## Quality Contract -- Calidad sobre velocidad

> Este proyecto usa SpecBox Engine. La velocidad ya esta resuelta por el sistema.
> Tu trabajo como agente es CALIDAD: lee antes de escribir, piensa antes de actuar,
> verifica antes de marcar como hecho.

### Reglas innegociables

1. **Lee antes de escribir** -- El hook `quality-first-guard.mjs` BLOQUEA modificaciones a archivos existentes si no los leiste primero en esta sesion. No hay excepciones.
2. **Piensa antes de actuar** -- Para tareas complejas (>3 archivos o logica no trivial), articula tu enfoque en texto visible antes de escribir codigo.
3. **Verifica antes de cerrar** -- No marques una tarea como completada sin verificar que funciona. "Deberia funcionar" no es verificacion.
4. **Pregunta antes de adivinar** -- Si no estas seguro del enfoque, pregunta al usuario. Una pregunta cuesta ~50 tokens. Una iteracion fallida cuesta miles.
5. **Una correcta > tres rapidas** -- Una implementacion bien pensada vale mas que tres intentos rapidos que cada uno necesita arreglo.

### Antipatrones que cuestan tokens

| Antipatron | Coste real | Alternativa |
|------------|-----------|-------------|
| Escribir sin leer | Rompe funcionalidad existente, requiere rollback | Leer primero (hook lo fuerza) |
| Adivinar enfoque | 3-5 iteraciones de healing, ~5K tokens cada una | Preguntar al usuario (~50 tokens) |
| "Ya esta" sin verificar | Bug descubierto tarde, requiere reabrir UC | Verificar lint + test + funcionalidad |
| Copiar codigo generico | No encaja con patrones del proyecto, refactor posterior | Leer codigo existente, seguir patrones |

## Reglas del Proyecto

### Importar reglas globales
Las reglas de specbox-engine/rules/GLOBAL_RULES.md aplican a este proyecto.

### Reglas especificas de este proyecto

- Backend Go: usar `cmd/` para entrypoints, `internal/` para logica de negocio, `pkg/` para librerias reutilizables
- Frontend React: componentes funcionales, hooks personalizados, TypeScript estricto
- API: comunicacion via REST/WebSocket entre Go backend y React frontend

## Agentes del Proyecto

Roles activos para este proyecto (Go + React fullstack):

| ID | Agente | Rol |
|----|--------|-----|
| AG-01 | Feature Generator | Genera estructura de features |
| AG-02 | UI/UX Designer | Diseño frontend React |
| AG-04 | QA Validation | Tests y validacion |
| AG-08 | Quality Auditor | Auditoria de calidad GO/NO-GO |
| AG-09a | Acceptance Tester | Tests E2E desde AC-XX |
| AG-09b | Acceptance Validator | Validacion ACCEPTED/REJECTED |

## Comandos Disponibles

| Comando | Proposito |
|---------|-----------|
| /prd | Genera PRD + work item |
| /plan | Plan de implementacion + diseños Stitch + VEG |
| /implement | Autopilot: rama + fases + design-to-code + QA + PR |
| /adapt-ui | Escanea y mapea widgets existentes |
| /optimize-agents | Audita y optimiza sistema agentico |
| /feedback | Reporta bugs de testing manual |

## Flujo de Desarrollo

```
/prd -> PRD + Trello/Plane (con Definition Quality Gate)
  |
/plan -> Plan tecnico + Diseños Stitch + VEG
  |
/implement -> Autopilot: rama + fases + QA + Acceptance Gate + PR
  |        -> AG-08 Quality Audit (GO/NO-GO)
  |        -> AG-09a Acceptance Tests
  |        -> AG-09b Acceptance Validator (ACCEPTED/REJECTED)
  |        -> Merge secuencial
  |
/optimize-agents -> Validar configuracion agentica
```

## Acceptance Testing -- Gherkin BDD

### Estructura Go

```
tests/acceptance/
  features/
    UC-XXX_{nombre}.feature
  steps/
    common_steps_test.go
    UC-XXX_steps_test.go
  reports/
    cucumber-report.json
    acceptance-report.pdf
```

### Framework

| Stack | Paquete |
|-------|---------|
| Go | testing + testify + httptest + testcontainers-go |
| React | playwright-bdd ^8.4.2 |

## E2E Testing

| Aspecto | Detalle |
|---------|---------|
| Backend | Go testing + httptest |
| Frontend | Playwright |
| Config | `playwright.config.ts` (React) |
| Report | `doc/test_cases/reports/` |

## Hooks

Los hooks (.mjs) estan enlazados via symlinks al engine (`specbox-engine/.claude/hooks/`).
Se ejecutan con `node` (cross-platform, v5.17.0+). Incluyen carpeta `lib/` con utilidades compartidas.

---

*Generado: 2026-04-04*
*Engine: SpecBox Engine v5.17.0 "Cross-Platform Hooks"*
