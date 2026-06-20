import { NgClass } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { BackButtonComponent } from '../back-button/back-button';

export interface ProjectItem {
  id: number;
  title: string;
  description: string;
  tags: string[];
  iconClass: string;
  bannerColor: string;
  imageUrl?: string;
  imageAlt?: string;
  repoUrl?: string;
  deployUrl?: string;
  detailRoute?: string;
}

@Component({
  selector: 'app-project-card-grid',
  standalone: true,
  imports: [ButtonModule, NgClass, RouterLink, BackButtonComponent],
  templateUrl: './project-card-grid.html',
  styleUrl: './project-card-grid.scss',
})
export class ProjectCardGrid {
  @Input({ required: true }) pageTitle = '';
  @Input({ required: true }) pageDescription = '';
  @Input({ required: true }) items: ProjectItem[] = [];

  getAccentColor(item: ProjectItem): string {
    const match = (item.bannerColor ?? '').match(/#[0-9a-fA-F]{6}/);
    return match?.[0] ?? '#4f46e5';
  }

  openUrl(url: string): void {
    window.open(url, '_blank', 'noopener,noreferrer');
  }

  onImageError(event: Event): void {
    (event.target as HTMLImageElement).style.display = 'none';
  }
}
