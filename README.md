# Desarrollo de una CLI con Cobra en Go

La empresa necesita una herramienta de línea de comandos (CLI) para gestionar operaciones en un sistema de banca digital. La CLI debe permitir a los usuarios realizar operaciones como consultar saldo, transferir fondos y obtener historial de transacciones. El sistema de banca digital tiene restricciones específicas: las transferencias deben ser idempotentes con una clave de operación única, y el sistema debe manejar errores de red con reintentos automáticos. El CLI debe ser robusto y manejar correctamente los casos límite del dominio.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go |
| **Nivel** | junior-l1 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Implementación de la estructura básica de la CLI

**Objetivo:** Crear la estructura básica de la CLI con Cobra, incluyendo comandos para consultar saldo y transferir fondos.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Diseña la estructura de la CLI con Cobra.
- Implementa el comando para consultar el saldo de una cuenta.
- Implementa el comando para transferir fondos entre cuentas.

**Entregable:** CLI con comandos básicos operativos para consultar saldo y transferir fondos.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo estructurar los comandos para que sean intuitivos para el usuario.
- Piensa en cómo manejar los errores de entrada de datos.

</details>

### Fase 2: Implementación de la idempotencia en las transferencias

**Objetivo:** Añadir idempotencia a las transferencias utilizando una clave de operación única.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Implementa la idempotencia en las transferencias utilizando una clave de operación única.
- Asegura que dos invocaciones con la misma clave produzcan un solo registro y devuelvan la misma respuesta dentro de una ventana de tiempo definida.

**Entregable:** CLI con transferencias idempotentes utilizando una clave de operación única.

<details>
<summary>Pistas de conocimiento</summary>

- Piensa en cómo generar y almacenar la clave de operación.
- Considera los posibles modos de falla y cómo manejarlos.

</details>

### Fase 3: Manejo de errores de red con reintentos automáticos

**Objetivo:** Implementar reintentos automáticos para manejar errores de red durante las transferencias.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Implementa reintentos automáticos para manejar errores de red durante las transferencias.
- Asegura que la CLI maneje correctamente los errores de red y reintenta las operaciones fallidas.

**Entregable:** CLI con reintentos automáticos para manejar errores de red durante las transferencias.

<details>
<summary>Pistas de conocimiento</summary>

- Piensa en cómo determinar el número de reintentos y el intervalo entre ellos.
- Considera los posibles modos de falla y cómo manejarlos.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es la idempotencia y por qué es importante en las transferencias?
- **paraQueSirve**: ¿Para qué sirve la CLI en el contexto de la banca digital?
- **comoSeUsa**: ¿Cómo se usa la CLI para realizar transferencias idempotentes?
- **erroresComunes**: ¿Cuáles son los errores comunes que puede encontrar la CLI y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones implica la implementación de reintentos automáticos para manejar errores de red?

## Criterios de Evaluacion

- Implementación correcta de la estructura básica de la CLI con Cobra.
- Implementación correcta de la idempotencia en las transferencias utilizando una clave de operación única.
- Implementación correcta de reintentos automáticos para manejar errores de red durante las transferencias.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
