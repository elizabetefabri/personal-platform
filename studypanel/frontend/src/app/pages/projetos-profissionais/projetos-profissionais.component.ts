import { Component } from '@angular/core';
import {
  ProjectCardGrid,
  ProjectItem,
} from '../../shared/components/project-card-grid/project-card-grid';

const PROFESSIONAL_PROJECTS: ProjectItem[] = [
  {
    id: 1,
    title: 'Rollout Service',
    description:
      'Plataforma para estudo e implementação de rollout gradual, score, agendamentos, backend e instrumentação Datadog.',
    tags: ['Angular', 'Backend', 'Datadog', 'Observability', 'Feature Rollout'],
    iconClass: 'pi-sliders-h',
    bannerColor: 'linear-gradient(135deg, #047857, #059669)',
    imageUrl: '/assets/images/projetos/projetos-profissionais/rollout-service/cover.png',
    imageAlt: 'Imagem do projeto Rollout Service',
    detailRoute: '/rollout-service',
  },
  {
    id: 2,
    title: 'IUDev',
    description:
      'Ferramenta interna/de estudo para apoiar fluxos de desenvolvimento, ambiente local, automações e produtividade em plataforma.',
    tags: ['Go', 'Angular', 'CLI', 'Platform Engineering', 'DevTools'],
    iconClass: 'pi-wrench',
    bannerColor: 'linear-gradient(135deg, #1d4ed8, #2563eb)',
    imageUrl: '/assets/images/projetos/projetos-profissionais/iudev/cover.png',
    imageAlt: 'Imagem do projeto IUDev',
  },
];

@Component({
  selector: 'app-projetos-profissionais',
  standalone: true,
  imports: [ProjectCardGrid],
  templateUrl: './projetos-profissionais.component.html',
  styleUrl: './projetos-profissionais.component.scss',
})
export class ProjetosProfissionaisComponent {
  readonly pageTitle = 'Projetos Profissionais';
  readonly pageDescription =
    'Projetos relacionados à atuação profissional, plataforma, engenharia, automações e observabilidade.';
  readonly items = PROFESSIONAL_PROJECTS;
}
