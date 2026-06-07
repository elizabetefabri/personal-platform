import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StudyDetailTemplate } from '../../shared/components/study-detail-template/study-detail-template';
import { getSectionBySlug, getTopicBySlug, SectionConfig, TopicConfig } from '../../core/constants/study-topics';

@Component({
  selector: 'app-study-detail-page',
  standalone: true,
  imports: [StudyDetailTemplate],
  templateUrl: './study-detail-page.component.html',
  styleUrl: './study-detail-page.component.scss',
})
export class StudyDetailPageComponent implements OnInit {
  section: SectionConfig | undefined;
  topic: TopicConfig | undefined;

  constructor(private router: Router) {}

  ngOnInit(): void {
    const segments = this.router.url.split('?')[0].split('/').filter(Boolean);
    this.section = getSectionBySlug(segments[0]);
    this.topic = segments[1] ? getTopicBySlug(segments[0], segments[1]) : undefined;
  }

  get pageTitle(): string {
    return this.topic?.label ?? this.section?.label ?? 'Detalhe';
  }

  get pageDescription(): string {
    return this.topic?.description ?? this.section?.description ?? '';
  }

  get accentColor(): string {
    const gradient = this.topic?.bannerColor ?? this.section?.bannerColor ?? '';
    const match = gradient.match(/#[0-9a-fA-F]{6}/);
    return match?.[0] ?? '#4f46e5';
  }
}
