import { Component, OnDestroy, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { AccordionContent, AccordionHeader, AccordionModule, AccordionPanel } from 'primeng/accordion';
import { ButtonModule } from 'primeng/button';
import { ProgressBarModule } from 'primeng/progressbar';
import { TableModule } from 'primeng/table';
import { TagModule } from 'primeng/tag';
import { BreadcrumbService } from '../../core/services/breadcrumb.service';
import { StudyItemService } from '../../core/services/study-item.service';
import {
  NoteStatus,
  StudyMilestone,
  StudyNote,
  StudySession,
  StudyStatus,
  StudyTableItem,
} from '../../shared/interfaces/study-template.interface';

@Component({
  selector: 'app-course-detail',
  standalone: true,
  imports: [
    AccordionContent,
    AccordionHeader,
    AccordionModule,
    AccordionPanel,
    ButtonModule,
    ProgressBarModule,
    RouterLink,
    TableModule,
    TagModule,
  ],
  templateUrl: './course-detail.component.html',
  styleUrl: './course-detail.component.scss',
})
export class CourseDetailComponent implements OnInit, OnDestroy {
  item = signal<StudyTableItem | null>(null);
  loading = signal(true);
  error = signal(false);

  sectionSlug = '';
  topicSlug = '';
  itemId = '';

  // ── Anotações ──
  notes = signal<StudyNote[]>([]);

  // ── Progresso & Metas ──
  milestones = signal<StudyMilestone[]>([
    { id: '1', title: 'Assistir aulas introdutórias', completed: false },
    { id: '2', title: 'Completar labs práticos', completed: false },
    { id: '3', title: 'Resolver simulado de 20 questões', completed: false },
    { id: '4', title: 'Revisar pontos fracos', completed: false },
    { id: '5', title: 'Aprovado no exame final', completed: false },
  ]);

  // ── Histórico de Sessões ──
  sessions = signal<StudySession[]>([]);

  readonly noteStatuses: NoteStatus[] = ['Rascunho', 'Revisado', 'Finalizado'];

  get overallProgress(): number {
    const list = this.milestones();
    if (!list.length) return 0;
    return Math.round((list.filter((m) => m.completed).length / list.length) * 100);
  }

  constructor(
    private route: ActivatedRoute,
    private studyItemService: StudyItemService,
    private breadcrumbService: BreadcrumbService,
  ) {}

  ngOnInit(): void {
    const allSegments = this.route.snapshot.pathFromRoot
      .flatMap((r) => r.url)
      .map((s) => s.path)
      .filter(Boolean);

    this.sectionSlug = allSegments[0] ?? '';
    this.topicSlug = allSegments[1] ?? '';
    this.itemId = allSegments[2] ?? '';

    this.loadItem();
  }

  ngOnDestroy(): void {
    this.breadcrumbService.clear();
  }

  private loadItem(): void {
    if (!this.itemId) {
      this.error.set(true);
      this.loading.set(false);
      return;
    }

    this.studyItemService.getById(this.itemId).subscribe({
      next: (data) => {
        this.item.set(data);
        this.breadcrumbService.set({ label: data.courseName });
        this.breadcrumbService.setPageTitle(data.courseName);
        this.loading.set(false);
      },
      error: () => {
        this.error.set(true);
        this.loading.set(false);
      },
    });
  }

  get backRoute(): string {
    return `/${this.sectionSlug}/${this.topicSlug}`;
  }

  toggleMilestone(id: string): void {
    this.milestones.update((list) =>
      list.map((m) => (m.id === id ? { ...m, completed: !m.completed } : m)),
    );
  }

  getStatusSeverity(status: StudyStatus): 'success' | 'info' | 'warn' | 'danger' {
    const map: Record<StudyStatus, 'success' | 'info' | 'warn' | 'danger'> = {
      Concluído: 'success',
      'Em andamento': 'info',
      'Não iniciado': 'warn',
      Pausado: 'danger',
    };
    return map[status];
  }

  getNoteStatusSeverity(status: NoteStatus): 'success' | 'info' | 'warn' {
    const map: Record<NoteStatus, 'success' | 'info' | 'warn'> = {
      Finalizado: 'success',
      Revisado: 'info',
      Rascunho: 'warn',
    };
    return map[status];
  }

  ratingStars(rating: number): string {
    return '★'.repeat(rating) + '☆'.repeat(5 - rating);
  }
}
