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

  private rawSections: CourseSection[] = [];

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
        this.rawSections = sections;
        this.items.set(this.toCards(sections));
        this.loading.set(false);
      },
      error: () => {
        this.items.set([]);
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
      iconUrl: s.imageUrl ? undefined : undefined,
      skill: s.name,
      detailRoute: `/${s.slug}`,
      imageUrl: s.imageUrl || undefined,
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

  onEditRequest(item: StudyCardItem): void {
    const section = this.rawSections.find((s) => s.id === item.apiId);
    if (section) this.openEditModal(section);
  }

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
