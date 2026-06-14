import { Component, OnInit, signal } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { ChartModule } from 'primeng/chart';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { ProgressBarModule } from 'primeng/progressbar';
import { TagModule } from 'primeng/tag';
import { StudyItemService } from '../../core/services/study-item.service';
import { STUDY_SECTIONS, SectionConfig, TopicConfig } from '../../core/constants/study-topics';
import {
  StudyStatus,
  StudyTableItem,
  CreateStudyItemDto,
} from '../../shared/interfaces/study-template.interface';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    ButtonModule,
    ChartModule,
    DialogModule,
    FormsModule,
    InputTextModule,
    ProgressBarModule,
    TagModule,
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss',
})
export class DashboardComponent implements OnInit {
  readonly sections = STUDY_SECTIONS;

  // ── Data ────────────────────────────────────────────────────────────────────
  items = signal<StudyTableItem[]>([]);
  loading = signal(true);

  // ── Stats ───────────────────────────────────────────────────────────────────
  get totalItems(): number { return this.items().length; }
  get inProgressItems(): number { return this.items().filter(i => i.status === 'Em andamento').length; }
  get completedItems(): number { return this.items().filter(i => i.status === 'Concluído').length; }
  get activeSections(): number {
    return new Set(this.items().map(i => i.section)).size;
  }
  get completionRate(): number {
    if (!this.totalItems) return 0;
    return Math.round((this.completedItems / this.totalItems) * 100);
  }

  // ── Charts ──────────────────────────────────────────────────────────────────
  statusChartData: any = null;
  evolutionChartData: any = null;

  readonly statusChartOptions = {
    plugins: {
      legend: { position: 'bottom', labels: { padding: 16, boxWidth: 12, font: { size: 12 } } },
    },
    responsive: true,
    maintainAspectRatio: false,
    cutout: '65%',
  };

  readonly evolutionChartOptions = {
    plugins: { legend: { display: false } },
    responsive: true,
    maintainAspectRatio: false,
    scales: {
      x: { grid: { display: false }, ticks: { font: { size: 11 } } },
      y: { beginAtZero: true, ticks: { precision: 0, font: { size: 11 } }, grid: { color: '#f3f4f6' } },
    },
  };

  // ── Section stats helper ───────────────────────────────────────────────────
  getSectionItems(sectionSlug: string): StudyTableItem[] {
    return this.items().filter(i => i.section === sectionSlug);
  }

  getSectionProgress(sectionSlug: string): number {
    const list = this.getSectionItems(sectionSlug);
    if (!list.length) return 0;
    return Math.round((list.filter(i => i.status === 'Concluído').length / list.length) * 100);
  }

  getSectionAccent(section: SectionConfig): string {
    const match = section.bannerColor.match(/#[0-9a-fA-F]{6}/);
    return match?.[0] ?? '#4f46e5';
  }

  getTopics(sectionSlug: string): TopicConfig[] {
    return this.sections.find(s => s.slug === sectionSlug)?.topics ?? [];
  }

  getStatusSeverity(status: StudyStatus): 'success' | 'info' | 'warn' | 'danger' | 'secondary' {
    const map: Record<StudyStatus, 'success' | 'info' | 'warn' | 'danger'> = {
      Concluído: 'success', 'Em andamento': 'info', 'Não iniciado': 'warn', Pausado: 'danger',
    };
    return map[status];
  }

  // ── Quick-add modal ────────────────────────────────────────────────────────
  addModalVisible = false;
  saving = signal(false);
  selectedSection: SectionConfig | null = null;
  availableTopics: TopicConfig[] = [];

  addForm: CreateStudyItemDto = {
    section: '',
    topic: '',
    courseName: '',
    status: 'Não iniciado',
    date: new Date().toISOString().slice(0, 10),
    url: '',
  };

  readonly statusOptions: StudyStatus[] = ['Não iniciado', 'Em andamento', 'Concluído', 'Pausado'];

  openAddModal(section?: SectionConfig): void {
    this.selectedSection = section ?? null;
    this.addForm = {
      section: section?.slug ?? '',
      topic: '',
      courseName: '',
      status: 'Não iniciado',
      date: new Date().toISOString().slice(0, 10),
      url: '',
    };
    this.availableTopics = section?.topics ?? [];
    this.addModalVisible = true;
  }

  onSectionChange(): void {
    const sec = this.sections.find(s => s.slug === this.addForm.section);
    this.selectedSection = sec ?? null;
    this.availableTopics = sec?.topics ?? [];
    this.addForm.topic = '';
  }

  saveItem(): void {
    if (!this.addForm.section || !this.addForm.topic || !this.addForm.courseName) return;
    this.saving.set(true);

    this.studyItemService.create(this.addForm).subscribe({
      next: (item) => {
        this.items.update(list => [item, ...list]);
        this.buildCharts();
        this.saving.set(false);
        this.addModalVisible = false;
      },
      error: () => this.saving.set(false),
    });
  }

  // ── Lifecycle ──────────────────────────────────────────────────────────────
  constructor(
    private studyItemService: StudyItemService,
    private router: Router,
  ) {}

  ngOnInit(): void {
    this.studyItemService.listAll().subscribe({
      next: (data) => {
        this.items.set(data);
        this.buildCharts();
        this.loading.set(false);
      },
      error: () => this.loading.set(false),
    });
  }

  private buildCharts(): void {
    const items = this.items();

    // Status doughnut
    const statusCounts: Record<StudyStatus, number> = {
      'Não iniciado': 0, 'Em andamento': 0, 'Concluído': 0, 'Pausado': 0,
    };
    items.forEach(i => { statusCounts[i.status] = (statusCounts[i.status] ?? 0) + 1; });

    this.statusChartData = {
      labels: ['Não iniciado', 'Em andamento', 'Concluído', 'Pausado'],
      datasets: [{
        data: [
          statusCounts['Não iniciado'],
          statusCounts['Em andamento'],
          statusCounts['Concluído'],
          statusCounts['Pausado'],
        ],
        backgroundColor: ['#e5e7eb', '#3b82f6', '#22c55e', '#f59e0b'],
        borderWidth: 0,
      }],
    };

    // Evolution bar — items by month (last 6 months)
    const monthLabels: string[] = [];
    const monthCounts: number[] = [];
    const now = new Date();

    for (let i = 5; i >= 0; i--) {
      const d = new Date(now.getFullYear(), now.getMonth() - i, 1);
      const label = d.toLocaleString('pt-BR', { month: 'short', year: '2-digit' });
      monthLabels.push(label);

      const count = items.filter(item => {
        if (!item.createdAt) return false;
        const created = new Date(item.createdAt);
        return created.getFullYear() === d.getFullYear() && created.getMonth() === d.getMonth();
      }).length;
      monthCounts.push(count);
    }

    this.evolutionChartData = {
      labels: monthLabels,
      datasets: [{
        label: 'Cursos cadastrados',
        data: monthCounts,
        backgroundColor: 'rgba(79,70,229,0.15)',
        borderColor: '#4f46e5',
        borderWidth: 2,
        borderRadius: 6,
        tension: 0.4,
        fill: true,
        pointBackgroundColor: '#4f46e5',
        pointRadius: 4,
      }],
    };
  }

  navigateToSection(sectionSlug: string): void {
    this.router.navigate([sectionSlug]);
  }
}
