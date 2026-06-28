import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TooltipModule } from 'primeng/tooltip';
import {
  ProjectCardGrid,
  ProjectItem,
} from '../../shared/components/project-card-grid/project-card-grid';
import {
  CreateProjectDto,
  Project,
  ProjectService,
  UpdateProjectDto,
} from '../../core/services/project.service';

@Component({
  selector: 'app-projetos-pessoais',
  standalone: true,
  imports: [ProjectCardGrid, ButtonModule, DialogModule, FormsModule, InputTextModule, TooltipModule],
  templateUrl: './projetos-pessoais.component.html',
  styleUrl: './projetos-pessoais.component.scss',
})
export class ProjetosPessoaisComponent implements OnInit {
  readonly pageTitle = 'Projetos Pessoais';
  readonly pageDescription = 'Projetos criados para estudo, portfólio, prática técnica e evolução profissional.';

  items = signal<ProjectItem[]>([]);
  saving = signal(false);
  errorMessage = '';
  modalVisible = false;
  editingProject: Project | null = null;

  private rawProjects: Project[] = [];

  deleteModalVisible = false;
  deleteTargetId = '';
  deleteTargetName = '';

  form: CreateProjectDto = {
    name: '', type: 'pessoal', description: '', tags: [],
    repoUrl: '', deployUrl: '', slug: '',
    bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
    imageUrl: '', imageAlt: '', detailRoute: '', active: true, order: 0,
  };
  tagsInput = '';

  constructor(private projectService: ProjectService) {}

  ngOnInit(): void { this.loadProjects(); }

  private loadProjects(): void {
    this.projectService.list('pessoal').subscribe({
      next: (projects) => {
        this.rawProjects = projects;
        this.items.set(this.toCards(projects));
      },
      error: () => this.items.set([]),
    });
  }

  private toCards(projects: Project[]): ProjectItem[] {
    return projects.map((p, i) => ({
      id: i + 1,
      apiId: p.id,
      title: p.name,
      description: p.description,
      tags: p.tags ?? [],
      iconClass: 'pi-desktop',
      bannerColor: p.bannerColor,
      imageUrl: p.imageUrl,
      imageAlt: p.imageAlt,
      repoUrl: p.repoUrl,
      deployUrl: p.deployUrl,
      detailRoute: p.detailRoute,
    }));
  }

  openCreateModal(): void {
    this.editingProject = null;
    this.form = {
      name: '', type: 'pessoal', description: '', tags: [],
      repoUrl: '', deployUrl: '', slug: '',
      bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
      imageUrl: '', imageAlt: '', detailRoute: '', active: true, order: 0,
    };
    this.tagsInput = '';
    this.errorMessage = '';
    this.modalVisible = true;
  }

  openEditModal(project: Project): void {
    this.editingProject = project;
    this.form = {
      name: project.name, type: project.type, description: project.description,
      tags: project.tags ?? [], repoUrl: project.repoUrl ?? '',
      deployUrl: project.deployUrl ?? '', slug: project.slug,
      bannerColor: project.bannerColor, imageUrl: project.imageUrl ?? '',
      imageAlt: project.imageAlt ?? '', detailRoute: project.detailRoute ?? '',
      active: project.active, order: project.order,
    };
    this.tagsInput = (project.tags ?? []).join(', ');
    this.errorMessage = '';
    this.modalVisible = true;
  }

  onEditRequest(item: ProjectItem): void {
    const project = this.rawProjects.find((p) => p.id === item.apiId);
    if (project) this.openEditModal(project);
  }

  saveProject(): void {
    if (!this.form.name) { this.errorMessage = 'Nome é obrigatório.'; return; }
    this.form.tags = this.tagsInput.split(',').map((t) => t.trim()).filter(Boolean);
    this.saving.set(true);
    this.errorMessage = '';
    const obs = this.editingProject
      ? this.projectService.update(this.editingProject.id, this.form as UpdateProjectDto)
      : this.projectService.create(this.form);
    obs.subscribe({
      next: () => { this.saving.set(false); this.modalVisible = false; this.loadProjects(); },
      error: (err) => { this.saving.set(false); this.errorMessage = err?.error?.error ?? 'Erro ao salvar.'; },
    });
  }

  cancelModal(): void { this.modalVisible = false; this.errorMessage = ''; }

  onDeleteRequest(item: ProjectItem): void {
    this.deleteTargetId = item.apiId ?? '';
    this.deleteTargetName = item.title;
    this.deleteModalVisible = true;
  }

  confirmDelete(): void {
    if (!this.deleteTargetId) { this.deleteModalVisible = false; return; }
    this.projectService.delete(this.deleteTargetId).subscribe({
      next: () => { this.deleteModalVisible = false; this.loadProjects(); },
      error: () => { this.deleteModalVisible = false; },
    });
  }

  cancelDelete(): void {
    this.deleteModalVisible = false;
    this.deleteTargetId = '';
    this.deleteTargetName = '';
  }
}
