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
    id: 1, title: 'Comanda Flow',
    description: 'Sistema de comanda digital para gestão de pedidos, fluxo de atendimento e organização operacional.',
    tags: ['Angular'], iconClass: 'pi-desktop',
    bannerColor: 'linear-gradient(135deg, #c3002f, #dd0031)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/comanda-flow/cover.png',
    repoUrl: 'https://github.com/elizabetefabri/personal-platform/tree/main/comandaflow/frontend',
    deployUrl: 'https://comandaflow.elizabetesousafabri.com.br/login/',
  },
  {
    id: 2, title: 'Dataverse Chat',
    description: 'Aplicação em React com experiência conversacional baseada em dados.',
    tags: ['React'], iconClass: 'pi-comments',
    bannerColor: 'linear-gradient(135deg, #1565c0, #1976d2)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/dataverse-chat/cover.png',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-dataverse-chat',
    deployUrl: 'https://dataversechat.vercel.app/',
  },
  {
    id: 3, title: 'Dataverse',
    description: 'Projeto em JavaScript para exibição, filtragem e organização de dados em interface interativa.',
    tags: ['JavaScript'], iconClass: 'pi-database',
    bannerColor: 'linear-gradient(135deg, #d97706, #f59e0b)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/dataverse/cover.png',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-dataverse',
    deployUrl: 'https://dataverse-coffee.vercel.app/',
  },
  {
    id: 4, title: 'Movie Challenge',
    description: 'Aplicação Angular para listagem, busca e visualização de filmes.',
    tags: ['Angular'], iconClass: 'pi-video',
    bannerColor: 'linear-gradient(135deg, #374151, #1f2937)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/movie-challenge/cover.png',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-movie-challenge-fw',
    deployUrl: 'https://moviechallenge.vercel.app/home',
  },
  {
    id: 5, title: 'Formulário Pré-candidatos',
    description: 'Formulário Angular para cadastro e organização de informações de pré-candidatos.',
    tags: ['Angular'], iconClass: 'pi-file-edit',
    bannerColor: 'linear-gradient(135deg, #0369a1, #0ea5e9)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/form-precandidatos/cover.png',
    repoUrl: 'https://github.com/elizabetefabri/formulario-pre-candidatos',
    deployUrl: 'https://formulario-precandidatos.vercel.app/formulario',
  },
];

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
      next: (projects) => { this.items.set(projects.length > 0 ? this.toCards(projects) : FALLBACK); },
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
