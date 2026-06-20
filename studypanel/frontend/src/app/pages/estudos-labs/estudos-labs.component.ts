import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { STUDY_SECTIONS } from '../../core/constants/study-topics';

const ESTUDOS_SLUGS = [
  'backend',
  'banco-de-dados',
  'cloud',
  'containers-kubernetes',
  'devops',
  'frontend',
  'inteligencia-artificial',
  'observability',
  'performance-engineering',
];

const ROUTE_MAP: Record<string, string> = {
  'banco-de-dados': '/banco-de-dados',
  cloud: '/cloud',
  'containers-kubernetes': '/containers-kubernetes',
};

const IMAGE_MAP: Record<string, string> = {
  backend: '/assets/images/estudos-labs/backend/cover.png',
  'banco-de-dados': '/assets/images/estudos-labs/banco-de-dados/cover.png',
  cloud: '/assets/images/estudos-labs/cloud-computing/cover.png',
  'containers-kubernetes': '/assets/images/estudos-labs/containers-kubernetes/cover.png',
  devops: '/assets/images/estudos-labs/devops/cover.png',
  frontend: '/assets/images/estudos-labs/frontend/cover.png',
  'inteligencia-artificial': '/assets/images/estudos-labs/inteligencia-artificial/cover.png',
  observability: '/assets/images/estudos-labs/observability/cover.png',
  'performance-engineering': '/assets/images/estudos-labs/performance-engineering/cover.png',
};

@Component({
  selector: 'app-estudos-labs',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './estudos-labs.component.html',
  styleUrl: './estudos-labs.component.scss',
})
export class EstudosLabsComponent {
  readonly pageTitle = 'Estudos e Labs';
  readonly pageDescription =
    'Trilhas de estudo organizadas por área de conhecimento. Escolha uma categoria e comece a explorar.';

  readonly items: StudyCardItem[] = STUDY_SECTIONS.filter((s) =>
    ESTUDOS_SLUGS.includes(s.slug),
  ).map((section, i) => ({
    id: i + 1,
    title: section.label,
    description: section.description,
    bannerColor: section.bannerColor,
    iconClass: section.iconClass,
    skill: section.label,
    detailRoute: ROUTE_MAP[section.slug] ?? `/${section.slug}`,
    imageUrl: IMAGE_MAP[section.slug],
  }));
}
