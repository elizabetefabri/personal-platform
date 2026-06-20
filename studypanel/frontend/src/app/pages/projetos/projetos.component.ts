import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';

@Component({
  selector: 'app-projetos',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './projetos.component.html',
  styleUrl: './projetos.component.scss',
})
export class ProjetosComponent {
  readonly pageTitle = 'Projetos';
  readonly pageDescription =
    'Explore os projetos pessoais e profissionais desenvolvidos ao longo da jornada técnica.';

  readonly items: StudyCardItem[] = [
    {
      id: 1,
      title: 'Projetos Pessoais',
      description:
        'Projetos criados para estudo, portfólio, prática técnica e evolução profissional.',
      bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
      iconClass: 'pi-user',
      skill: 'Portfólio',
      detailRoute: '/projetos/pessoais',
      buttonLabel: 'Abrir Projetos',
    },
    {
      id: 2,
      title: 'Projetos Profissionais',
      description:
        'Projetos relacionados à atuação profissional, plataforma, engenharia, automações e observabilidade.',
      bannerColor: 'linear-gradient(135deg, #047857, #059669)',
      iconClass: 'pi-briefcase',
      skill: 'Profissional',
      detailRoute: '/projetos/profissionais',
      buttonLabel: 'Abrir Projetos',
    },
  ];
}
