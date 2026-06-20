import { NgClass } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { TooltipModule } from 'primeng/tooltip';
import { BackButtonComponent } from '../back-button/back-button';
import { StudyCardItem } from '../../interfaces/study-template.interface';

@Component({
  selector: 'app-study-card-grid',
  imports: [ButtonModule, NgClass, RouterLink, TooltipModule, BackButtonComponent],
  templateUrl: './study-card-grid.html',
  styleUrl: './study-card-grid.scss',
})
export class StudyCardGrid {
  @Input({ required: true }) pageTitle = '';
  @Input({ required: true }) pageDescription = '';
  @Input({ required: true }) items: StudyCardItem[] = [];
  @Output() deleteRequest = new EventEmitter<StudyCardItem>();

  getAccentColor(item: StudyCardItem): string {
    const match = (item.bannerColor ?? '').match(/#[0-9a-fA-F]{6}/);
    return match?.[0] ?? '#4f46e5';
  }

  onImageError(event: Event): void {
    (event.target as HTMLImageElement).style.display = 'none';
  }

  requestDelete(item: StudyCardItem, event: Event): void {
    event.stopPropagation();
    this.deleteRequest.emit(item);
  }
}
