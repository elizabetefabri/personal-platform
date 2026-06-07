# Components Index

Documentação de todos os componentes compartilhados do StudyPanel.

## Componentes Disponíveis

### Layout & Navigation
- **[Sidebar](./SIDEBAR.md)** - Navegação lateral colapsável com menu principal
- **[Header](./HEADER.md)** - Cabeçalho da aplicação com título
- **[Breadcrumbs](./BREADCRUMBS.md)** - Navegação hierárquica automática
- **[Footer](./FOOTER.md)** - Rodapé da aplicação

### Content Display
- **[Study Card Grid](./STUDY_CARD_GRID.md)** - Grade responsiva de cards de estudos

## Como Usar Este Índice

1. Cada componente tem sua própria documentação individual
2. Consulte a documentação específica para:
   - Inputs e Outputs
   - Exemplos de uso
   - Estilos e temas
   - Integração com BD (futura)
   - Notas técnicas

## Padrões Comuns

Todos os componentes seguem:
- Angular 19+ com standalone components
- TypeScript strongly typed
- SCSS para estilos
- PrimeIcons para ícones
- Responsive design

## Adicionando Novos Componentes

Ao criar um novo componente:
1. Crie a pasta em `src/app/shared/components/nome-componente/`
2. Implemente: `.ts`, `.html`, `.scss`
3. Crie documentação: `docs/frontend/components/NOME.md`
4. Atualize este `INDEX.md`

