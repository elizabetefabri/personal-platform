import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TooltipModule } from 'primeng/tooltip';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import {
  CourseTopic,
  CourseTopicService,
  CreateCourseTopicDto,
  UpdateCourseTopicDto,
} from '../../core/services/course-topic.service';
import { getSectionBySlug } from '../../core/constants/study-topics';

@Component({
  selector: 'app-study-section-page',
  standalone: true,
  imports: [StudyCardGrid, ButtonModule, DialogModule, FormsModule, InputTextModule, TooltipModule],
  templateUrl: './study-section-page.component.html',
  styleUrl: './study-section-page.component.scss',
})
export class StudySectionPageComponent implements OnInit {
  sectionSlug = '';
  pageTitle = '';
  pageDescription = '';
  accentColor = '';

  items = signal<StudyCardItem[]>([]);
  loading = signal(false);
  saving = signal(false);
  errorMessage = '';

  modalVisible = false;
  editingTopic: CourseTopic | null = null;

  private rawTopics: CourseTopic[] = [];

  deleteModalVisible = false;
  deleteTargetId = '';
  deleteTargetName = '';

  form: CreateCourseTopicDto = {
    sectionSlug: '', slug: '', label: '', description: '',
    bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
    iconClass: 'pi-code', skill: '', imageUrl: '', active: true, order: 0,
  };

  constructor(
    private route: ActivatedRoute,
    private topicService: CourseTopicService,
  ) {}

  ngOnInit(): void {
    this.sectionSlug = this.route.parent?.snapshot.url[0]?.path ?? '';
    const section = getSectionBySlug(this.sectionSlug);
    this.pageTitle = section?.label ?? this.toLabel(this.sectionSlug);
    this.pageDescription = section?.description ?? '';
    this.accentColor = section?.bannerColor ?? '#4f46e5';
    this.loadTopics();
  }

  private loadTopics(): void {
    this.loading.set(true);
    this.topicService.list(this.sectionSlug).subscribe({
      next: (topics) => {
        this.rawTopics = topics;
        this.items.set(topics.length > 0 ? this.toCards(topics) : this.fallbackCards());
        this.loading.set(false);
      },
      error: () => { this.items.set(this.fallbackCards()); this.loading.set(false); },
    });
  }

  private toCards(topics: CourseTopic[]): StudyCardItem[] {
    return topics.map((t, i) => ({
      id: i + 1,
      apiId: t.id,
      title: t.label,
      description: t.description,
      bannerColor: t.bannerColor,
      iconClass: t.iconClass,
      iconUrl: `/assets/images/estudos-labs/${this.sectionSlug}/${t.slug}.svg`,
      skill: t.skill,
      detailRoute: `/${this.sectionSlug}/${t.slug}`,
      imageUrl: t.imageUrl,
    }));
  }

  private fallbackCards(): StudyCardItem[] {
    const section = getSectionBySlug(this.sectionSlug);
    if (!section) return [];
    return section.topics.map((t, i) => ({
      id: i + 1,
      title: t.label,
      description: t.description,
      bannerColor: t.bannerColor,
      iconClass: t.iconClass,
      iconUrl: `/assets/images/estudos-labs/${this.sectionSlug}/${t.slug}.svg`,
      skill: t.skill,
      detailRoute: `/${this.sectionSlug}/${t.slug}`,
    }));
  }

  openCreateModal(): void {
    this.editingTopic = null;
    this.form = {
      sectionSlug: this.sectionSlug, slug: '', label: '', description: '',
      bannerColor: this.accentColor || 'linear-gradient(135deg, #4f46e5, #7c3aed)',
      iconClass: 'pi-code', skill: '', imageUrl: '', active: true, order: 0,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  openEditModal(topic: CourseTopic): void {
    this.editingTopic = topic;
    this.form = {
      sectionSlug: topic.sectionSlug, slug: topic.slug, label: topic.label,
      description: topic.description, bannerColor: topic.bannerColor,
      iconClass: topic.iconClass, skill: topic.skill,
      imageUrl: topic.imageUrl ?? '', active: topic.active, order: topic.order,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  saveTopic(): void {
    if (!this.form.label || !this.form.slug) { this.errorMessage = 'Label e slug são obrigatórios.'; return; }
    this.saving.set(true);
    this.errorMessage = '';
    const obs = this.editingTopic
      ? this.topicService.update(this.editingTopic.id, this.form as UpdateCourseTopicDto)
      : this.topicService.create(this.form);
    obs.subscribe({
      next: () => { this.saving.set(false); this.modalVisible = false; this.loadTopics(); },
      error: (err) => { this.saving.set(false); this.errorMessage = err?.error?.error ?? 'Erro ao salvar.'; },
    });
  }

  cancelModal(): void { this.modalVisible = false; this.errorMessage = ''; }

  onEditRequest(item: StudyCardItem): void {
    const topic = this.rawTopics.find((t) => t.id === item.apiId);
    if (topic) this.openEditModal(topic);
  }

  onDeleteRequest(item: StudyCardItem): void {
    this.deleteTargetId = item.apiId ?? '';
    this.deleteTargetName = item.title;
    this.deleteModalVisible = true;
  }

  confirmDelete(): void {
    if (!this.deleteTargetId) { this.deleteModalVisible = false; return; }
    this.topicService.delete(this.deleteTargetId).subscribe({
      next: () => { this.deleteModalVisible = false; this.loadTopics(); },
      error: () => { this.deleteModalVisible = false; },
    });
  }

  cancelDelete(): void {
    this.deleteModalVisible = false;
    this.deleteTargetId = '';
    this.deleteTargetName = '';
  }

  private toLabel(slug: string): string {
    return slug.split('-').map((w) => w.charAt(0).toUpperCase() + w.slice(1)).join(' ');
  }
}
