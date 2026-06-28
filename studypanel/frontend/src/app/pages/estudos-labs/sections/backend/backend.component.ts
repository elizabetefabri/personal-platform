import { Component } from '@angular/core';
import { StudyCardGrid } from '../../../../shared/components/study-card-grid/study-card-grid';
import { getSectionBySlug } from '../../../../core/constants/study-topics';
import { StudyCardItem } from '../../../../shared/interfaces/study-template.interface';

@Component({
  selector: 'app-backend',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './backend.component.html',
  styleUrl: './backend.component.scss',
})
export class BackendComponent {
  private readonly section = getSectionBySlug('backend')!;

  readonly pageTitle = this.section.label;
  readonly pageDescription = this.section.description;

  readonly items: StudyCardItem[] = this.section.topics.map((topic, i) => ({
    id: i + 1,
    title: topic.label,
    description: topic.description,
    bannerColor: topic.bannerColor,
    iconClass: topic.iconClass,
    skill: topic.skill,
    detailRoute: `/backend/${topic.slug}`,
  }));
}
