import { Component, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { BackButtonComponent } from '../../shared/components/back-button/back-button';
import {
  CulinaryService,
  CulinaryCategory,
  CreateCulinaryCategoryDto,
  UpdateCulinaryCategoryDto,
} from '../../core/services/culinary.service';

@Component({
  selector: 'app-culinaria',
  standalone: true,
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ButtonModule,
    DialogModule,
    InputTextModule,
    BackButtonComponent,
  ],
  templateUrl: './culinaria.html',
  styleUrl: './culinaria.scss',
})
export class Culinaria implements OnInit {
  categories = signal<CulinaryCategory[]>([]);
  loading = signal(false);
  saving = signal(false);
  errorMessage = '';

  modalVisible = false;
  editingCategory: CulinaryCategory | null = null;

  deleteModalVisible = false;
  deleteTargetId = '';
  deleteTargetName = '';

  form: CreateCulinaryCategoryDto = {
    name: '',
    slug: '',
    description: '',
    tag: '',
    color: '#201F25',
    icon: 'pi-utensils',
    imageUrl: '',
    order: 0,
    active: true,
  };

  constructor(private culinaryService: CulinaryService) {}

  ngOnInit(): void {
    this.loadCategories();
  }

  loadCategories(): void {
    this.loading.set(true);
    this.culinaryService.listCategories().subscribe({
      next: (cats) => {
        this.categories.set(cats);
        this.loading.set(false);
      },
      error: () => {
        this.categories.set([]);
        this.loading.set(false);
      },
    });
  }

  openCreateModal(): void {
    this.editingCategory = null;
    this.form = {
      name: '',
      slug: '',
      description: '',
      tag: '',
      color: '#201F25',
      icon: 'pi-utensils',
      imageUrl: '',
      order: 0,
      active: true,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  openEditModal(cat: CulinaryCategory): void {
    this.editingCategory = cat;
    this.form = {
      name: cat.name,
      slug: cat.slug,
      description: cat.description,
      tag: cat.tag,
      color: cat.color,
      icon: cat.icon,
      imageUrl: cat.imageUrl ?? '',
      order: cat.order,
      active: cat.active,
    };
    this.errorMessage = '';
    this.modalVisible = true;
  }

  saveCategory(): void {
    if (!this.form.name || !this.form.slug) {
      this.errorMessage = 'Nome e slug são obrigatórios.';
      return;
    }
    this.saving.set(true);
    this.errorMessage = '';

    const obs = this.editingCategory
      ? this.culinaryService.updateCategory(this.editingCategory.id, this.form as UpdateCulinaryCategoryDto)
      : this.culinaryService.createCategory(this.form);

    obs.subscribe({
      next: () => {
        this.saving.set(false);
        this.modalVisible = false;
        this.loadCategories();
      },
      error: (err) => {
        this.saving.set(false);
        this.errorMessage = err?.error?.error ?? 'Erro ao salvar categoria.';
      },
    });
  }

  cancelModal(): void {
    this.modalVisible = false;
    this.errorMessage = '';
  }

  openDeleteModal(cat: CulinaryCategory): void {
    this.deleteTargetId = cat.id;
    this.deleteTargetName = cat.name;
    this.deleteModalVisible = true;
  }

  confirmDelete(): void {
    if (!this.deleteTargetId) { this.deleteModalVisible = false; return; }
    this.culinaryService.deleteCategory(this.deleteTargetId).subscribe({
      next: () => {
        this.deleteModalVisible = false;
        this.loadCategories();
      },
      error: () => { this.deleteModalVisible = false; },
    });
  }

  cancelDelete(): void {
    this.deleteModalVisible = false;
    this.deleteTargetId = '';
    this.deleteTargetName = '';
  }
}
