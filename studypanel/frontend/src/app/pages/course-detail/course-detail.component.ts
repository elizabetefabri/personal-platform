import { Component, OnDestroy, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { TagModule } from 'primeng/tag';
import { BreadcrumbService } from '../../core/services/breadcrumb.service';
import { StudyItemService } from '../../core/services/study-item.service';
import { StudyStatus, StudyTableItem } from '../../shared/interfaces/study-template.interface';

@Component({
  selector: 'app-course-detail',
  standalone: true,
  imports: [RouterLink, TagModule],
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

  getStatusSeverity(status: StudyStatus): 'success' | 'info' | 'warn' | 'danger' {
    const map: Record<StudyStatus, 'success' | 'info' | 'warn' | 'danger'> = {
      Concluído: 'success',
      'Em andamento': 'info',
      'Não iniciado': 'warn',
      Pausado: 'danger',
    };
    return map[status];
  }
}
