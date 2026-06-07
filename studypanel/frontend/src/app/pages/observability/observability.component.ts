import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { getSectionBySlug } from '../../core/constants/study-topics';

@Component({
  selector: 'app-observability',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './observability.component.html',
  styleUrl: './observability.component.scss',
})
export class ObservabilityComponent {
  private readonly section = getSectionBySlug('observability')!;

  readonly pageTitle = this.section.label;
  readonly pageDescription = this.section.description;

  readonly items: StudyCardItem[] = this.section.topics.map((topic, i) => ({
    id: i + 1,
    title: topic.label,
    description: topic.description,
    bannerColor: topic.bannerColor,
    iconClass: topic.iconClass,
    skill: topic.skill,
    detailRoute: `/observability/${topic.slug}`,
  }));
}
