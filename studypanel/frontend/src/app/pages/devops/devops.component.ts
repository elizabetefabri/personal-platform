import { Component } from '@angular/core';
import { StudyTableItem } from '../../shared/interfaces/study-template.interface';
import { StudyDetailTemplate } from '../../shared/components/study-detail-template/study-detail-template';

@Component({
  selector: 'app-devops',
  standalone: true,
  imports: [StudyDetailTemplate],
  templateUrl: './devops.component.html',
  styleUrl: './devops.component.scss',
})
export class DevOpsComponent {
   readonly items: StudyTableItem[] = [
    {
      id: 1,
      courseName: 'Angular v21 do Zero ao Avançado',
      status: 'Em andamento',
      date: '2026-06-07',
      url: 'https://angular.dev',
    },
    {
      id: 2,
      courseName: 'Testes Unitários com Jest',
      status: 'Não iniciado',
      date: '2026-06-10',
      url: 'https://jestjs.io',
    },
    {
      id: 3,
      courseName: 'SCSS com Design System',
      status: 'Concluído',
      date: '2026-06-15',
      url: 'https://sass-lang.com',
    },
  ];
}
