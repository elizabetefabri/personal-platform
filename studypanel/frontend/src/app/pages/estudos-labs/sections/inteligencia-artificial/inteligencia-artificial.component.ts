import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { getSectionBySlug } from '../../core/constants/study-topics';

@Component({
  selector: 'app-inteligencia-artificial',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './inteligencia-artificial.component.html',
  styleUrl: './inteligencia-artificial.component.scss',
})
export class InteligenciaArtificialComponent {
  private readonly section = getSectionBySlug('inteligencia-artificial')!;

  readonly pageTitle = this.section.label;
  readonly pageDescription = this.section.description;

  readonly items: StudyCardItem[] = this.section.topics.map((topic, i) => ({
    id: i + 1,
    title: topic.label,
    description: topic.description,
    bannerColor: topic.bannerColor,
    iconClass: topic.iconClass,
    skill: topic.skill,
    detailRoute: `/inteligencia-artificial/${topic.slug}`,
  }));
}
