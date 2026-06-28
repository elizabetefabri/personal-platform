import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { environment } from '../../../environments/environment';

export interface CulinaryCategory {
  id: string;
  name: string;
  slug: string;
  description: string;
  tag: string;
  color: string;
  icon: string;
  imageUrl?: string;
  order: number;
  active: boolean;
  createdAt?: string;
  updatedAt?: string;
}

export interface CreateCulinaryCategoryDto {
  name: string;
  slug: string;
  description: string;
  tag: string;
  color: string;
  icon: string;
  imageUrl?: string;
  order: number;
  active: boolean;
}

export interface UpdateCulinaryCategoryDto {
  name: string;
  description: string;
  tag: string;
  color: string;
  icon: string;
  imageUrl?: string;
  order: number;
  active: boolean;
}

export interface CulinaryRecipe {
  id: string;
  categoryId: string;
  categorySlug: string;
  name: string;
  slug: string;
  description: string;
  prepTimeMinutes: number;
  cookTimeMinutes: number;
  servingsStr: string;
  difficulty: string;
  status: string;
  tags: string[];
  imageUrl?: string;
  youtubeUrl?: string;
  sourceUrl?: string;
  ingredients: string[];
  preparationSteps: string[];
  utensils?: string;
  tips?: string;
  substitutions?: string;
  storageInstructions?: string;
  estimatedCost: number;
  personalRating: number;
  tested: boolean;
  testedAt?: string;
  notes?: string;
  active: boolean;
  createdAt?: string;
  updatedAt?: string;
}

export interface CreateCulinaryRecipeDto {
  categoryId: string;
  categorySlug: string;
  name: string;
  slug: string;
  description: string;
  prepTimeMinutes: number;
  cookTimeMinutes: number;
  servingsStr: string;
  difficulty: string;
  status: string;
  tags: string[];
  imageUrl?: string;
  youtubeUrl?: string;
  sourceUrl?: string;
  ingredients: string[];
  preparationSteps: string[];
  utensils?: string;
  tips?: string;
  substitutions?: string;
  storageInstructions?: string;
  estimatedCost: number;
  personalRating: number;
  tested: boolean;
  notes?: string;
  active: boolean;
}

export interface UpdateCulinaryRecipeDto extends CreateCulinaryRecipeDto {}

interface ApiResponse<T> {
  success: boolean;
  data: T;
  error?: string;
}

@Injectable({ providedIn: 'root' })
export class CulinaryService {
  private readonly baseCategories = `${environment.apiUrl}/api/v1/culinary/categories`;
  private readonly baseRecipes = `${environment.apiUrl}/api/v1/culinary/recipes`;

  constructor(private http: HttpClient) {}

  // ── Categories ──────────────────────────────────────────────────────────────

  listCategories(): Observable<CulinaryCategory[]> {
    return this.http
      .get<ApiResponse<CulinaryCategory[]>>(this.baseCategories)
      .pipe(map((res) => res.data ?? []));
  }

  getCategoryById(id: string): Observable<CulinaryCategory> {
    return this.http
      .get<ApiResponse<CulinaryCategory>>(`${this.baseCategories}/${id}`)
      .pipe(map((res) => res.data));
  }

  createCategory(dto: CreateCulinaryCategoryDto): Observable<CulinaryCategory> {
    return this.http
      .post<ApiResponse<CulinaryCategory>>(this.baseCategories, dto)
      .pipe(map((res) => res.data));
  }

  updateCategory(id: string, dto: UpdateCulinaryCategoryDto): Observable<CulinaryCategory> {
    return this.http
      .put<ApiResponse<CulinaryCategory>>(`${this.baseCategories}/${id}`, dto)
      .pipe(map((res) => res.data));
  }

  deleteCategory(id: string): Observable<void> {
    return this.http.delete<void>(`${this.baseCategories}/${id}`);
  }

  // ── Recipes ─────────────────────────────────────────────────────────────────

  listRecipes(categorySlug?: string): Observable<CulinaryRecipe[]> {
    const url = categorySlug
      ? `${this.baseRecipes}?category_slug=${encodeURIComponent(categorySlug)}`
      : this.baseRecipes;
    return this.http
      .get<ApiResponse<CulinaryRecipe[]>>(url)
      .pipe(map((res) => res.data ?? []));
  }

  getRecipeById(id: string): Observable<CulinaryRecipe> {
    return this.http
      .get<ApiResponse<CulinaryRecipe>>(`${this.baseRecipes}/${id}`)
      .pipe(map((res) => res.data));
  }

  createRecipe(dto: CreateCulinaryRecipeDto): Observable<CulinaryRecipe> {
    return this.http
      .post<ApiResponse<CulinaryRecipe>>(this.baseRecipes, dto)
      .pipe(map((res) => res.data));
  }

  updateRecipe(id: string, dto: UpdateCulinaryRecipeDto): Observable<CulinaryRecipe> {
    return this.http
      .put<ApiResponse<CulinaryRecipe>>(`${this.baseRecipes}/${id}`, dto)
      .pipe(map((res) => res.data));
  }

  deleteRecipe(id: string): Observable<void> {
    return this.http.delete<void>(`${this.baseRecipes}/${id}`);
  }
}
