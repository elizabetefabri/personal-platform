# Agents - StudyPanel Notes

This document describes how the agents defined in `.agentes/` should interact with the StudyPanel frontend and what responsibilities they hold.

## Purpose

Keep agent responsibilities aligned with the current implementation and update their tasks when StudyPanel frontend changes.

## Agents and responsibilities (StudyPanel-specific)

- **Frontend Agent**
  - Maintain component library
  - Implement pages and routes
  - Validate responsive behavior on target breakpoints
  - Implement SCSS using the design system
  - Keep `@personal-platform/studypanel/documentation/` up to date

- **DevOps Agent**
  - Manage deployment via the `deploy.ts` script
  - Keep `.env.deployment.example` updated
  - Validate FTP uploads and permissions
  - Automate deployment steps when credentials provided

- **QA Agent**
  - Create test plans for pages/components
  - Run `npm run test` and report coverage issues
  - Validate routes and perform smoke tests after deploy

- **Arquiteto / Product Owner**
  - Approve structural changes (component additions, routing changes)
  - Sign off on deployments to production

## How to update the agents README

Whenever a change is made to the frontend (new component, route, or deployment changes), update `.agentes/README.md` and `@personal-platform/studypanel/documentation/agents_update.md` to reflect the new responsibilities and instructions.
