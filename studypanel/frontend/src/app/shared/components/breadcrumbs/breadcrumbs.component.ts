import { Component, computed, OnDestroy, OnInit } from '@angular/core';
import { Router, NavigationEnd, RouterLink } from '@angular/router';
import { filter, Subscription } from 'rxjs';
import { getSectionBySlug, getTopicBySlug } from '../../../core/constants/study-topics';
import { BreadcrumbItem, BreadcrumbService } from '../../../core/services/breadcrumb.service';

export type { BreadcrumbItem };

const PAGE_LABELS: Record<string, string> = {
  settings: 'Configurações',
};

@Component({
  selector: 'app-breadcrumbs',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './breadcrumbs.component.html',
  styleUrl: './breadcrumbs.component.scss',
})
export class BreadcrumbsComponent implements OnInit, OnDestroy {
  baseItems: BreadcrumbItem[] = [];

  readonly items = computed<BreadcrumbItem[]>(() => {
    const extra = this.breadcrumbService.extra();
    return extra ? [...this.baseItems, extra] : this.baseItems;
  });

  private sub = new Subscription();

  constructor(
    private router: Router,
    private breadcrumbService: BreadcrumbService,
  ) {}

  ngOnInit(): void {
    this.build(this.router.url);

    this.sub.add(
      this.router.events
        .pipe(filter((e): e is NavigationEnd => e instanceof NavigationEnd))
        .subscribe((e) => this.build(e.urlAfterRedirects)),
    );
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  private build(rawUrl: string): void {
    const url = rawUrl.split('?')[0];
    const segments = url.split('/').filter(Boolean);

    if (segments.length === 0 || segments[0] === 'dashboard') {
      this.baseItems = [];
      return;
    }

    const sectionSlug = segments[0];
    const section = getSectionBySlug(sectionSlug);
    const sectionLabel =
      section?.label ?? PAGE_LABELS[sectionSlug] ?? this.toLabel(sectionSlug);

    if (segments.length === 1) {
      this.baseItems = [
        { label: 'Dashboard', href: '/dashboard' },
        { label: sectionLabel },
      ];
      return;
    }

    const topicSlug = segments[1];
    const topic = section ? getTopicBySlug(sectionSlug, topicSlug) : undefined;
    const topicLabel = topic?.label ?? this.toLabel(topicSlug);

    if (segments.length === 2) {
      this.baseItems = [
        { label: 'Dashboard', href: '/dashboard' },
        { label: sectionLabel, href: `/${sectionSlug}` },
        { label: topicLabel },
      ];
      return;
    }

    // 3+ segments: section/topic/itemId — extra label comes from BreadcrumbService
    this.baseItems = [
      { label: 'Dashboard', href: '/dashboard' },
      { label: sectionLabel, href: `/${sectionSlug}` },
      { label: topicLabel, href: `/${sectionSlug}/${topicSlug}` },
    ];
  }

  private toLabel(slug: string): string {
    return slug
      .split('-')
      .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
      .join(' ');
  }
}
