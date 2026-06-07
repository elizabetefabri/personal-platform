import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { getSectionBySlug } from '../../core/constants/study-topics';

@Component({
  selector: 'app-cloud',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './cloud.component.html',
  styleUrl: './cloud.component.scss',
})
export class CloudComponent {
  private readonly section = getSectionBySlug('cloud')!;

  readonly pageTitle = this.section.label;
  readonly pageDescription = this.section.description;

  readonly items: StudyCardItem[] = this.section.topics.map((topic, i) => ({
    id: i + 1,
    title: topic.label,
    description: topic.description,
    bannerColor: topic.bannerColor,
    iconClass: topic.iconClass,
    skill: topic.skill,
    detailRoute: `/cloud/${topic.slug}`,
  }));
}
