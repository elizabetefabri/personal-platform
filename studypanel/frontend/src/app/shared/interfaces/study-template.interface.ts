export type StudyStatus = 'Não iniciado' | 'Em andamento' | 'Concluído' | 'Pausado';
export type NoteStatus = 'Rascunho' | 'Revisado' | 'Finalizado';
export type ResourceType = 'Documentação' | 'Vídeo' | 'Artigo' | 'Ferramenta' | 'Livro' | 'Outro';
export type QuizDifficulty = 'Fácil' | 'Médio' | 'Difícil';

// ─── Study Item ───────────────────────────────────────────────────────────────

export interface StudyCardItem {
  id: number;
  title: string;
  description: string;
  imageUrl?: string;
  bannerColor?: string;
  iconClass: string;
  skill: string;
  detailRoute: string;
}

export interface StudyTableItem {
  id: string;
  section: string;
  topic: string;
  courseName: string;
  status: StudyStatus;
  date: string;
  url: string;
  imageUrl?: string;
  detailRoute: string;
  createdAt?: string;
  updatedAt?: string;
}

export interface CreateStudyItemDto {
  section: string;
  topic: string;
  courseName: string;
  status: StudyStatus;
  date: string;
  url: string;
  imageUrl?: string;
}

export interface UpdateStudyItemDto {
  courseName: string;
  status: StudyStatus;
  date: string;
  url: string;
  imageUrl?: string;
}

// ─── Study Note ───────────────────────────────────────────────────────────────

export interface StudyNote {
  id: string;
  studyItemId: string;
  date: string;
  title: string;
  description: string;
  progress: number;
  status: NoteStatus;
  link?: string;
  createdAt?: string;
  updatedAt?: string;
}

// ─── Study Resource ───────────────────────────────────────────────────────────

export interface StudyResource {
  id: string;
  studyItemId: string;
  title: string;
  url: string;
  type: ResourceType;
  description?: string;
  createdAt?: string;
  updatedAt?: string;
}

// ─── Quiz Question ────────────────────────────────────────────────────────────

export interface QuizQuestion {
  id: string;
  section: string;
  topic: string;
  question: string;
  answer: string;
  explanation: string;
  difficulty: QuizDifficulty;
  tags: string[];
  createdAt?: string;
}

// ─── Study Session ────────────────────────────────────────────────────────────

export interface StudySession {
  id: string;
  studyItemId: string;
  startedAt: string;
  endedAt: string;
  durationSeconds: number;
  topic: string;
  notes?: string;
  rating: number;
  milestonesCompleted: number;
  createdAt?: string;
}

// ─── UI only ──────────────────────────────────────────────────────────────────

export interface StudyMilestone {
  id: string;
  title: string;
  completed: boolean;
  date?: string;
}
