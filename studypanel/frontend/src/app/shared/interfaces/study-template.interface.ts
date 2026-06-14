export type StudyStatus = 'Não iniciado' | 'Em andamento' | 'Concluído' | 'Pausado';

export type NoteStatus = 'Rascunho' | 'Revisado' | 'Finalizado';

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

export interface StudyNote {
  id: string;
  date: string;
  title: string;
  description: string;
  progress: number;
  status: NoteStatus;
  link?: string;
}

export interface StudyMilestone {
  id: string;
  title: string;
  completed: boolean;
  date?: string;
}

export interface StudySession {
  id: string;
  date: string;
  duration: string;
  topic: string;
  notes?: string;
  rating: 1 | 2 | 3 | 4 | 5;
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
