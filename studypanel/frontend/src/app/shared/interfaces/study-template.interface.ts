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
  id: number;
  courseName: string;
  status: StudyStatus;
  date: string;
  url: string;
}
