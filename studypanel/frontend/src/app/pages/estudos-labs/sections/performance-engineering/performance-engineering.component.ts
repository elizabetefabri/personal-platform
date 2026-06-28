import { Component } from '@angular/core';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import { getSectionBySlug } from '../../core/constants/study-topics';

@Component({
  selector: 'app-performance-engineering',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './performance-engineering.component.html',
  styleUrl: './performance-engineering.component.scss',
})
export class PerformanceEngineeringComponent {
  private readonly section = getSectionBySlug('performance-engineering')!;

  readonly pageTitle = this.section.label;
  readonly pageDescription = this.section.description;

  readonly items: StudyCardItem[] = this.section.topics.map((topic, i) => ({
    id: i + 1,
    title: topic.label,
    description: topic.description,
    bannerColor: topic.bannerColor,
    iconClass: topic.iconClass,
    skill: topic.skill,
    detailRoute: `/performance-engineering/${topic.slug}`,
  }));
}
