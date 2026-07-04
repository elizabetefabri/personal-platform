# StudyPanel Frontend

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 21.2.9.

## Documentation

All StudyPanel frontend documentation is centralized under:

`@personal-platform/studypanel/documentation/`

Key files:

- `README.md` — this file
- `@personal-platform/studypanel/documentation/commands.md` — all commands used to install, build, test and deploy
- `@personal-platform/studypanel/documentation/deploy.md` — FTP deploy instructions and `.env.deployment` usage
- `@personal-platform/studypanel/documentation/components.md` — list of components and routes
- `@personal-platform/studypanel/documentation/agents_update.md` — agent responsibilities and notes

## Development server

To start a local development server, run:

```bash
npm install
ng serve
# or
npm start
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
npm run build
# production build
npm run build:prod
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Deploying

Follow the instructions in `@personal-platform/studypanel/documentation/deploy.md`. In short:

1. Copy `.env.deployment.example` -> `.env.deployment` and fill FTP credentials
2. Run `npm run deploy` (build + deploy)

## Running unit tests

To execute unit tests with the [Vitest](https://vitest.dev/) test runner, use the following command:

```bash
npm run test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
npm run e2e
```

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.
