import { Component, OnDestroy, OnInit, signal } from '@angular/core';
import { SlicePipe } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { AccordionContent, AccordionHeader, AccordionModule, AccordionPanel } from 'primeng/accordion';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { ProgressBarModule } from 'primeng/progressbar';
import { SliderModule } from 'primeng/slider';
import { TableModule } from 'primeng/table';
import { TagModule } from 'primeng/tag';
import { TextareaModule } from 'primeng/textarea';
import { BreadcrumbService } from '../../core/services/breadcrumb.service';
import { QuizQuestionService } from '../../core/services/quiz-question.service';
import { StudyItemService } from '../../core/services/study-item.service';
import { StudyNoteService, CreateNoteDto } from '../../core/services/study-note.service';
import { StudyResourceService, CreateResourceDto } from '../../core/services/study-resource.service';
import { StudySessionService, CreateSessionDto } from '../../core/services/study-session.service';
import {
  NoteStatus,
  QuizDifficulty,
  QuizQuestion,
  ResourceType,
  StudyMilestone,
  StudyNote,
  StudyResource,
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
    DialogModule,
    FormsModule,
    InputTextModule,
    ProgressBarModule,
    RouterLink,
    SlicePipe,
    SliderModule,
    TableModule,
    TagModule,
    TextareaModule,
  ],
  templateUrl: './course-detail.component.html',
  styleUrl: './course-detail.component.scss',
})
export class CourseDetailComponent implements OnInit, OnDestroy {
  // ── Item principal ──────────────────────────────────────────────────────────
  item = signal<StudyTableItem | null>(null);
  loading = signal(true);
  error = signal(false);

  sectionSlug = '';
  topicSlug = '';
  itemId = '';

  // ── Anotações ──────────────────────────────────────────────────────────────
  notes = signal<StudyNote[]>([]);
  noteModalVisible = false;
  editingNote: StudyNote | null = null;
  savingNote = signal(false);
  noteForm: CreateNoteDto = { date: '', title: '', description: '', progress: 0, status: 'Rascunho', link: '' };
  readonly noteStatusOptions: NoteStatus[] = ['Rascunho', 'Revisado', 'Finalizado'];

  // ── Recursos ───────────────────────────────────────────────────────────────
  resources = signal<StudyResource[]>([]);
  resourceModalVisible = false;
  editingResource: StudyResource | null = null;
  savingResource = signal(false);
  resourceForm: CreateResourceDto = { title: '', url: '', type: 'Documentação', description: '' };
  readonly resourceTypeOptions: ResourceType[] = ['Documentação', 'Vídeo', 'Artigo', 'Ferramenta', 'Livro', 'Outro'];
  readonly resourceTypeIcons: Record<ResourceType, string> = {
    'Documentação': 'pi-file-pdf',
    'Vídeo': 'pi-video',
    'Artigo': 'pi-file',
    'Ferramenta': 'pi-wrench',
    'Livro': 'pi-book',
    'Outro': 'pi-link',
  };

  // ── Quiz / Flashcards ──────────────────────────────────────────────────────
  questions = signal<QuizQuestion[]>([]);
  quizStarted = false;
  quizFinished = false;
  currentQuestionIndex = 0;
  revealed = false;
  correctCount = 0;
  incorrectCount = 0;
  readonly difficultyColors: Record<QuizDifficulty, string> = {
    'Fácil': 'success',
    'Médio': 'warn',
    'Difícil': 'danger',
  };

  // ── Progresso & Metas ──────────────────────────────────────────────────────
  milestones = signal<StudyMilestone[]>([
    { id: '1', title: 'Assistir aulas introdutórias', completed: false },
    { id: '2', title: 'Completar labs práticos', completed: false },
    { id: '3', title: 'Resolver simulado (20 questões)', completed: false },
    { id: '4', title: 'Revisar pontos fracos', completed: false },
    { id: '5', title: 'Aprovado no quiz de revisão', completed: false },
  ]);

  get overallProgress(): number {
    const list = this.milestones();
    if (!list.length) return 0;
    return Math.round((list.filter((m) => m.completed).length / list.length) * 100);
  }

  // ── Histórico de Sessões ──────────────────────────────────────────────────
  sessions = signal<StudySession[]>([]);
  sessionStartTime: Date | null = null;
  sessionModalVisible = false;
  savingSession = signal(false);
  sessionRating = 3;
  sessionNotes = '';
  sessionElapsedSeconds = 0;
  private sessionTimer: ReturnType<typeof setInterval> | null = null;

  constructor(
    private route: ActivatedRoute,
    private studyItemService: StudyItemService,
    private noteService: StudyNoteService,
    private resourceService: StudyResourceService,
    private quizService: QuizQuestionService,
    private sessionService: StudySessionService,
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
    this.startSession();
  }

  ngOnDestroy(): void {
    this.breadcrumbService.clear();
    this.stopSessionTimer();
    // Sessão só salva via modal — não auto-salva no destroy
  }

  // ── Load ────────────────────────────────────────────────────────────────────

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
        this.loadAll();
      },
      error: () => {
        this.error.set(true);
        this.loading.set(false);
      },
    });
  }

  private loadAll(): void {
    this.noteService.list(this.itemId).subscribe({ next: (d) => this.notes.set(d) });
    this.resourceService.list(this.itemId).subscribe({ next: (d) => this.resources.set(d) });
    this.quizService.list(this.sectionSlug, this.topicSlug).subscribe({ next: (d) => this.questions.set(d) });
    this.sessionService.list(this.itemId).subscribe({ next: (d) => this.sessions.set(d) });
  }

  get backRoute(): string { return `/${this.sectionSlug}/${this.topicSlug}`; }

  // ── Anotações ──────────────────────────────────────────────────────────────

  openNoteModal(note?: StudyNote): void {
    this.editingNote = note ?? null;
    this.noteForm = note
      ? { date: note.date, title: note.title, description: note.description, progress: note.progress, status: note.status, link: note.link ?? '' }
      : { date: new Date().toISOString().slice(0, 10), title: '', description: '', progress: 0, status: 'Rascunho', link: '' };
    this.noteModalVisible = true;
  }

  saveNote(): void {
    this.savingNote.set(true);
    const dto: CreateNoteDto = { ...this.noteForm };

    const obs = this.editingNote
      ? this.noteService.update(this.itemId, this.editingNote.id, dto)
      : this.noteService.create(this.itemId, dto);

    obs.subscribe({
      next: (saved) => {
        this.notes.update((list) =>
          this.editingNote
            ? list.map((n) => (n.id === saved.id ? saved : n))
            : [...list, saved],
        );
        this.savingNote.set(false);
        this.noteModalVisible = false;
      },
      error: () => this.savingNote.set(false),
    });
  }

  deleteNote(note: StudyNote): void {
    this.noteService.delete(this.itemId, note.id).subscribe({
      next: () => this.notes.update((list) => list.filter((n) => n.id !== note.id)),
    });
  }

  // ── Recursos ───────────────────────────────────────────────────────────────

  openResourceModal(res?: StudyResource): void {
    this.editingResource = res ?? null;
    this.resourceForm = res
      ? { title: res.title, url: res.url, type: res.type, description: res.description ?? '' }
      : { title: '', url: '', type: 'Documentação', description: '' };
    this.resourceModalVisible = true;
  }

  saveResource(): void {
    this.savingResource.set(true);
    const dto: CreateResourceDto = { ...this.resourceForm };

    const obs = this.editingResource
      ? this.resourceService.update(this.itemId, this.editingResource.id, dto)
      : this.resourceService.create(this.itemId, dto);

    obs.subscribe({
      next: (saved) => {
        this.resources.update((list) =>
          this.editingResource
            ? list.map((r) => (r.id === saved.id ? saved : r))
            : [...list, saved],
        );
        this.savingResource.set(false);
        this.resourceModalVisible = false;
      },
      error: () => this.savingResource.set(false),
    });
  }

  deleteResource(res: StudyResource): void {
    this.resourceService.delete(this.itemId, res.id).subscribe({
      next: () => this.resources.update((list) => list.filter((r) => r.id !== res.id)),
    });
  }

  // ── Quiz / Flashcards ──────────────────────────────────────────────────────

  startQuiz(): void {
    this.quizStarted = true;
    this.quizFinished = false;
    this.currentQuestionIndex = 0;
    this.correctCount = 0;
    this.incorrectCount = 0;
    this.revealed = false;
  }

  revealAnswer(): void { this.revealed = true; }

  answerQuiz(correct: boolean): void {
    if (correct) this.correctCount++;
    else this.incorrectCount++;

    const next = this.currentQuestionIndex + 1;
    if (next >= this.questions().length) {
      this.quizFinished = true;
      if (correct && this.overallProgress === 100) {
        this.milestones.update((list) =>
          list.map((m) => (m.id === '5' ? { ...m, completed: true } : m)),
        );
      }
    } else {
      this.currentQuestionIndex = next;
      this.revealed = false;
    }
  }

  restartQuiz(): void {
    this.quizStarted = false;
    this.quizFinished = false;
    this.currentQuestionIndex = 0;
    this.correctCount = 0;
    this.incorrectCount = 0;
    this.revealed = false;
  }

  get quizScore(): number {
    const total = this.correctCount + this.incorrectCount;
    return total > 0 ? Math.round((this.correctCount / total) * 100) : 0;
  }

  get currentQuestion(): QuizQuestion | null {
    return this.questions()[this.currentQuestionIndex] ?? null;
  }

  // ── Progresso & Metas ──────────────────────────────────────────────────────

  toggleMilestone(id: string): void {
    this.milestones.update((list) =>
      list.map((m) => (m.id === id ? { ...m, completed: !m.completed } : m)),
    );
  }

  // ── Histórico de Sessões ──────────────────────────────────────────────────

  startSession(): void {
    this.sessionStartTime = new Date();
    this.sessionElapsedSeconds = 0;
    this.sessionTimer = setInterval(() => this.sessionElapsedSeconds++, 1000);
  }

  private stopSessionTimer(): void {
    if (this.sessionTimer) {
      clearInterval(this.sessionTimer);
      this.sessionTimer = null;
    }
  }

  openSessionModal(): void {
    this.stopSessionTimer();
    this.sessionModalVisible = true;
  }

  saveSession(): void {
    if (!this.sessionStartTime) return;
    this.savingSession.set(true);

    const now = new Date();
    const dto: CreateSessionDto = {
      startedAt: this.sessionStartTime.toISOString(),
      endedAt: now.toISOString(),
      durationSeconds: this.sessionElapsedSeconds || Math.floor((now.getTime() - this.sessionStartTime.getTime()) / 1000),
      topic: this.item()?.courseName ?? this.topicSlug,
      notes: this.sessionNotes,
      rating: this.sessionRating,
      milestonesCompleted: this.milestones().filter((m) => m.completed).length,
    };

    this.sessionService.create(this.itemId, dto).subscribe({
      next: (saved) => {
        this.sessions.update((list) => [saved, ...list]);
        this.savingSession.set(false);
        this.sessionModalVisible = false;
        this.sessionNotes = '';
        this.sessionRating = 3;
        // reinicia timer para próxima sessão
        this.startSession();
      },
      error: () => this.savingSession.set(false),
    });
  }

  formatDuration(seconds: number): string {
    const h = Math.floor(seconds / 3600);
    const m = Math.floor((seconds % 3600) / 60);
    const s = seconds % 60;
    if (h > 0) return `${h}h ${m}min`;
    if (m > 0) return `${m}min ${s}s`;
    return `${s}s`;
  }

  get elapsedLabel(): string { return this.formatDuration(this.sessionElapsedSeconds); }

  sessionTrend(idx: number): { icon: string; label: string; cls: string } | null {
    const list = this.sessions();
    if (idx >= list.length - 1) return null;
    const curr = list[idx].durationSeconds;
    const prev = list[idx + 1].durationSeconds;
    if (curr < prev * 0.85) return { icon: 'pi-arrow-down', label: 'Mais rápido', cls: 'trend--faster' };
    if (curr > prev * 1.15) return { icon: 'pi-arrow-up', label: 'Mais lento', cls: 'trend--slower' };
    return { icon: 'pi-minus', label: 'Estável', cls: 'trend--stable' };
  }

  // ── Status helpers ─────────────────────────────────────────────────────────

  getStatusSeverity(status: StudyStatus): 'success' | 'info' | 'warn' | 'danger' {
    const map: Record<StudyStatus, 'success' | 'info' | 'warn' | 'danger'> = {
      Concluído: 'success', 'Em andamento': 'info', 'Não iniciado': 'warn', Pausado: 'danger',
    };
    return map[status];
  }

  getNoteStatusSeverity(status: NoteStatus): 'success' | 'info' | 'warn' {
    const map: Record<NoteStatus, 'success' | 'info' | 'warn'> = {
      Finalizado: 'success', Revisado: 'info', Rascunho: 'warn',
    };
    return map[status];
  }

  getDifficultyTag(d: QuizDifficulty): 'success' | 'warn' | 'danger' {
    return this.difficultyColors[d] as 'success' | 'warn' | 'danger';
  }

  ratingStars(rating: number): string {
    return '★'.repeat(rating) + '☆'.repeat(5 - rating);
  }
}
