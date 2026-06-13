export type StudyStatus = 'Não iniciado' | 'Em andamento' | 'Concluído' | 'Pausado';

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
