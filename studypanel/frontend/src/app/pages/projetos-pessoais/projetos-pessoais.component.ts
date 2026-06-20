import { Component } from '@angular/core';
import {
  ProjectCardGrid,
  ProjectItem,
} from '../../shared/components/project-card-grid/project-card-grid';

const PERSONAL_PROJECTS: ProjectItem[] = [
  {
    id: 1,
    title: 'Comanda Flow',
    description:
      'Sistema de comanda digital para gestão de pedidos, fluxo de atendimento e organização operacional.',
    tags: ['Angular'],
    iconClass: 'pi-desktop',
    bannerColor: 'linear-gradient(135deg, #c3002f, #dd0031)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/comanda-flow/cover.png',
    imageAlt: 'Imagem do projeto Comanda Flow',
    repoUrl:
      'https://github.com/elizabetefabri/personal-platform/tree/main/comandaflow/frontend',
    deployUrl: 'https://comandaflow.elizabetesousafabri.com.br/login/',
  },
  {
    id: 2,
    title: 'Dataverse Chat',
    description:
      'Aplicação em React com experiência conversacional baseada em dados e interação com interface web.',
    tags: ['React'],
    iconClass: 'pi-comments',
    bannerColor: 'linear-gradient(135deg, #1565c0, #1976d2)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/dataverse-chat/cover.png',
    imageAlt: 'Imagem do projeto Dataverse Chat',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-dataverse-chat',
    deployUrl: 'https://dataversechat.vercel.app/',
  },
  {
    id: 3,
    title: 'Dataverse',
    description:
      'Projeto em JavaScript para exibição, filtragem e organização de dados em uma interface interativa.',
    tags: ['JavaScript'],
    iconClass: 'pi-database',
    bannerColor: 'linear-gradient(135deg, #d97706, #f59e0b)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/dataverse/cover.png',
    imageAlt: 'Imagem do projeto Dataverse',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-dataverse',
    deployUrl: 'https://dataverse-coffee.vercel.app/',
  },
  {
    id: 4,
    title: 'Movie Challenge',
    description:
      'Aplicação Angular para listagem, busca e visualização de filmes, criada como desafio técnico.',
    tags: ['Angular'],
    iconClass: 'pi-video',
    bannerColor: 'linear-gradient(135deg, #374151, #1f2937)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/movie-challenge/cover.png',
    imageAlt: 'Imagem do projeto Movie Challenge',
    repoUrl: 'https://github.com/elizabetefabri/SAP012-movie-challenge-fw',
    deployUrl: 'https://moviechallenge.vercel.app/home',
  },
  {
    id: 5,
    title: 'Formulário Pré-candidatos',
    description:
      'Formulário Angular para cadastro e organização de informações de pré-candidatos.',
    tags: ['Angular'],
    iconClass: 'pi-file-edit',
    bannerColor: 'linear-gradient(135deg, #0369a1, #0ea5e9)',
    imageUrl: '/assets/images/projetos/projetos-pessoais/form-precandidatos/cover.png',
    imageAlt: 'Imagem do projeto Formulário Pré-candidatos',
    repoUrl: 'https://github.com/elizabetefabri/formulario-pre-candidatos',
    deployUrl: 'https://formulario-precandidatos.vercel.app/formulario',
  },
];

@Component({
  selector: 'app-projetos-pessoais',
  standalone: true,
  imports: [ProjectCardGrid],
  templateUrl: './projetos-pessoais.component.html',
  styleUrl: './projetos-pessoais.component.scss',
})
export class ProjetosPessoaisComponent {
  readonly pageTitle = 'Projetos Pessoais';
  readonly pageDescription =
    'Projetos criados para estudo, portfólio, prática técnica e evolução profissional.';
  readonly items = PERSONAL_PROJECTS;
}
