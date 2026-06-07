import { Component, Input } from '@angular/core';
import { StudyCardItem } from '../../interfaces/study-template.interface';
import { NgClass } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { TagModule } from 'primeng/tag';

@Component({
  selector: 'app-study-card-grid',
  imports: [ButtonModule, NgClass, RouterLink, TagModule],
  templateUrl: './study-card-grid.html',
  styleUrl: './study-card-grid.scss',
})
export class StudyCardGrid {
  @Input({ required: true }) pageTitle = '';
  @Input({ required: true }) pageDescription = '';
  @Input({ required: true }) items: StudyCardItem[] = [];
}
