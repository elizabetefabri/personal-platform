import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { STUDY_SECTIONS } from '../../core/constants/study-topics';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss',
})
export class DashboardComponent {
  readonly pageTitle = 'Dashboard';
  readonly pageDescription = 'Acompanhe seu progresso nos diferentes tópicos de estudo.';

  readonly items: StudyCardItem[] = STUDY_SECTIONS.map((section, i) => ({
    id: i + 1,
    title: section.label,
    description: section.description,
    bannerColor: section.bannerColor,
    iconClass: section.iconClass,
    skill: 'Estudo',
    detailRoute: `/${section.slug}`,
  }));
}
