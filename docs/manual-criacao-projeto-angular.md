# Manual de Criação de Projeto Angular v21

## Objetivo

Este manual define o padrão para criação de novos projetos Angular dentro da `personal-platform`.

## Estrutura padrão

Todo projeto deve seguir:

```txt
nome-projeto/
├── frontend/
└── backend/
```

## Criar frontend

```
cd personal-platform/nome-projeto
ng new frontend --routing --style=scss --standalone
cd frontend
```

### Limpeza inicial

**Remover:**

- conteúdo padrão do Angular;
- imagens de exemplo;
- favicon padrão;
- estilos desnecessários.

**Adicionar temporariamente:**

```
<h1>Hello World</h1>
```

## Estrutura obrigatória

```
src/
├── app/
├── core/
├── shared/
├── pages/
├── assets/
├── styles/
└── environments/
```

## SCSS

**Criar:**

```
src/styles/base/_reset.scss
src/styles/styles.scss
```

**Reset obrigatório**

```
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html {
  font-size: 100%;
}

body {
  min-height: 100vh;
  text-rendering: optimizeLegibility;
}
```

### Regras visuais

> Usar rem.
> Evitar px.

> Criar layout mobile first.
> Usar media queries progressivas.
> Evitar largura fixa sem necessidade.
> Fontes

Configurar:

Open Sans;
Lato.
