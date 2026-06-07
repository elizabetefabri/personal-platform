import { Component, Input } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TableModule } from 'primeng/table';
import { TagModule } from 'primeng/tag';
import { StudyStatus, StudyTableItem } from '../../interfaces/study-template.interface';

@Component({
  selector: 'app-study-detail-template',
  imports: [
    ButtonModule,
    DialogModule,
    FormsModule,
    InputTextModule,
    TableModule,
    TagModule,
  ],
  templateUrl: './study-detail-template.html',
  styleUrl: './study-detail-template.scss',
})
export class StudyDetailTemplate {
  @Input({ required: true }) pageTitle = '';
  @Input({ required: true }) pageDescription = '';
  @Input() createButtonLabel = 'Cadastrar item';
  @Input() items: StudyTableItem[] = [];

  modalVisible = false;

  form = {
    courseName: '',
    status: 'Não iniciado' as StudyStatus,
    date: '',
    url: '',
  };

  openCreateModal(): void {
    this.modalVisible = true;
  }

  closeModal(): void {
    this.modalVisible = false;
  }

  saveItem(): void {
    console.log('Cadastro enviado:', this.form);
    this.closeModal();
  }

  viewItem(item: StudyTableItem): void {
    console.log('Visualizar item:', item);
  }

  editItem(item: StudyTableItem): void {
    console.log('Editar item:', item);
  }

  deleteItem(item: StudyTableItem): void {
    console.log('Excluir item:', item);
  }

  getStatusSeverity(status: StudyStatus): 'success' | 'info' | 'warning' | 'danger' {
    const severityMap = {
      'Concluído': 'success',
      'Em andamento': 'info',
      'Não iniciado': 'warning',
      Pausado: 'danger',
    } as const;

    return severityMap[status];
  }
}
