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

const FALLBACK: ProjectItem[] = [
  {
    id: 1, title: 'Rollout Service',
    description: 'Plataforma para estudo e implementação de rollout gradual, score, agendamentos e instrumentação.',
    tags: ['Angular', 'Backend', 'Datadog', 'Feature Rollout'],
    iconClass: 'pi-sliders-h',
    bannerColor: 'linear-gradient(135deg, #CD3F7B, #8b043d)',
    imageUrl: '/assets/images/projetos/projetos-profissionais/rollout-service/cover.png',
    detailRoute: '/rollout-service',
  },
  {
    id: 2, title: 'IUDev',
    description: 'Ferramenta interna para apoiar fluxos de desenvolvimento, automações e produtividade.',
    tags: ['Go', 'Angular', 'CLI', 'DevTools'],
    iconClass: 'pi-wrench',
    bannerColor: 'linear-gradient(135deg, #F7670F, #a0440c)',
    imageUrl: '/assets/images/projetos/projetos-profissionais/iudev/cover.png',
    detailRoute: '/iudev',
  },
];

@Component({
  selector: 'app-projetos-profissionais',
  standalone: true,
  imports: [ProjectCardGrid, ButtonModule, DialogModule, FormsModule, InputTextModule, TooltipModule],
  templateUrl: './projetos-profissionais.component.html',
  styleUrl: './projetos-profissionais.component.scss',
})
export class ProjetosProfissionaisComponent implements OnInit {
  readonly pageTitle = 'Projetos Profissionais';
  readonly pageDescription = 'Projetos relacionados à atuação profissional, plataforma, engenharia e observabilidade.';

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
    name: '', type: 'profissional', description: '', tags: [],
    repoUrl: '', deployUrl: '', slug: '',
    bannerColor: 'linear-gradient(135deg, #4f46e5, #7c3aed)',
    imageUrl: '', imageAlt: '', detailRoute: '', active: true, order: 0,
  };
  tagsInput = '';

  constructor(private projectService: ProjectService) {}

  ngOnInit(): void { this.loadProjects(); }

  private loadProjects(): void {
    this.projectService.list('profissional').subscribe({
      next: (projects) => {
        this.rawProjects = projects;
        this.items.set(projects.length > 0 ? this.toCards(projects) : FALLBACK);
      },
      error: () => this.items.set(FALLBACK),
    });
  }

  private toCards(projects: Project[]): ProjectItem[] {
    return projects.map((p, i) => ({
      id: i + 1,
      apiId: p.id,
      title: p.name,
      description: p.description,
      tags: p.tags ?? [],
      iconClass: 'pi-wrench',
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
      name: '', type: 'profissional', description: '', tags: [],
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
