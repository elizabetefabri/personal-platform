import { NgClass } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { StudyCardItem } from '../../interfaces/study-template.interface';

@Component({
  selector: 'app-study-card-grid',
  imports: [ButtonModule, NgClass, RouterLink],
  templateUrl: './study-card-grid.html',
  styleUrl: './study-card-grid.scss',
})
export class StudyCardGrid {
  @Input({ required: true }) pageTitle = '';
  @Input({ required: true }) pageDescription = '';
  @Input({ required: true }) items: StudyCardItem[] = [];

  getAccentColor(item: StudyCardItem): string {
    const match = (item.bannerColor ?? '').match(/#[0-9a-fA-F]{6}/);
    return match?.[0] ?? '#4f46e5';
  }
}
