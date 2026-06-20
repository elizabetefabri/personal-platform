import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TooltipModule } from 'primeng/tooltip';
import { StudyCardGrid } from '../../shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '../../shared/interfaces/study-template.interface';
import {
  CourseSection,
  CourseSectionService,
  CreateCourseSectionDto,
  UpdateCourseSectionDto,
} from '../../core/services/course-section.service';
import { getSectionBySlug, STUDY_SECTIONS } from '../../core/constants/study-topics';

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

const FALLBACK_SLUGS = [
  'backend', 'banco-de-dados', 'cloud', 'containers-kubernetes', 'devops',
  'frontend', 'inteligencia-artificial', 'observability', 'performance-engineering',
];

@Component({
  selector: 'app-estudos-labs',
  standalone: true,
  imports: [StudyCardGrid, ButtonModule, DialogModule, FormsModule, InputTextModule, TooltipModule],
  templateUrl: './estudos-labs.component.html',
  styleUrl: './estudos-labs.component.scss',
})
export class EstudosLabsComponent implements OnInit {
  readonly pageTitle = 'Estudos e Labs';
  readonly pageDescription =
    'Trilhas de estudo organizadas por área de conhecimento. Escolha uma categoria e comece a explorar.';

  items = signal<StudyCardItem[]>([]);
  loading = signal(false);
  saving = signal(false);
  errorMessage = '';

  modalVisible = false;
  editingSection: CourseSection | null = null;

  deleteModalVisible = false;
  deleteTargetId = '';
  deleteTargetName = '';

  form: CreateCourseSectionDto = {
    slug: '', name: '', description: '',
    bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
    iconClass: 'pi-book', imageUrl: '', active: true, order: 0,
  };

  constructor(private sectionService: CourseSectionService) {}

  ngOnInit(): void {
    this.loadSections();
  }

  private loadSections(): void {
    this.loading.set(true);
    this.sectionService.list().subscribe({
      next: (sections) => {
        this.items.set(sections.length > 0 ? this.toCards(sections) : this.fallbackCards());
        this.loading.set(false);
      },
      error: () => {
        this.items.set(this.fallbackCards());
        this.loading.set(false);
      },
    });
  }

  private toCards(sections: CourseSection[]): StudyCardItem[] {
    return sections.map((s, i) => ({
      id: i + 1,
      apiId: s.id,
      title: s.name,
      description: s.description,
      bannerColor: s.bannerColor,
      iconClass: s.iconClass,
      skill: s.name,
      detailRoute: `/${s.slug}`,
      imageUrl: s.imageUrl || IMAGE_MAP[s.slug],
    }));
  }

  private fallbackCards(): StudyCardItem[] {
    return STUDY_SECTIONS.filter((s) => FALLBACK_SLUGS.includes(s.slug)).map((s, i) => ({
      id: i + 1,
      title: s.label,
      description: s.description,
      bannerColor: s.bannerColor,
      iconClass: s.iconClass,
      skill: s.label,
      detailRoute: `/${s.slug}`,
      imageUrl: IMAGE_MAP[s.slug],
    }));
  }

  openCreateModal(): void {
    this.editingSection = null;
    this.form = {
      slug: '', name: '', description: '',
      bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
      iconClass: 'pi-book', imageUrl: '', active: true, order: 0,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  openEditModal(section: CourseSection): void {
    this.editingSection = section;
    this.form = {
      slug: section.slug, name: section.name, description: section.description,
      bannerColor: section.bannerColor, iconClass: section.iconClass,
      imageUrl: section.imageUrl ?? '', active: section.active, order: section.order,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  saveSection(): void {
    if (!this.form.name || !this.form.slug) { this.errorMessage = 'Nome e slug são obrigatórios.'; return; }
    this.saving.set(true);
    this.errorMessage = '';
    const obs = this.editingSection
      ? this.sectionService.update(this.editingSection.id, this.form as UpdateCourseSectionDto)
      : this.sectionService.create(this.form);
    obs.subscribe({
      next: () => { this.saving.set(false); this.modalVisible = false; this.loadSections(); },
      error: (err) => { this.saving.set(false); this.errorMessage = err?.error?.error ?? 'Erro ao salvar.'; },
    });
  }

  cancelModal(): void { this.modalVisible = false; this.errorMessage = ''; }

  onDeleteRequest(item: StudyCardItem): void {
    this.deleteTargetId = item.apiId ?? '';
    this.deleteTargetName = item.title;
    this.deleteModalVisible = true;
  }

  confirmDelete(): void {
    if (!this.deleteTargetId) { this.deleteModalVisible = false; return; }
    this.sectionService.delete(this.deleteTargetId).subscribe({
      next: () => { this.deleteModalVisible = false; this.loadSections(); },
      error: () => { this.deleteModalVisible = false; },
    });
  }

  cancelDelete(): void {
    this.deleteModalVisible = false;
    this.deleteTargetId = '';
    this.deleteTargetName = '';
  }
}
