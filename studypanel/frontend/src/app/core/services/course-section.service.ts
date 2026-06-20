import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { environment } from '../../../environments/environment';

export interface CourseSection {
  id: string;
  slug: string;
  name: string;
  description: string;
  bannerColor: string;
  iconClass: string;
  imageUrl?: string;
  active: boolean;
  order: number;
}

export interface CreateCourseSectionDto {
  slug: string;
  name: string;
  description: string;
  bannerColor: string;
  iconClass: string;
  imageUrl?: string;
  active: boolean;
  order: number;
}

export interface UpdateCourseSectionDto {
  name: string;
  description: string;
  bannerColor: string;
  iconClass: string;
  imageUrl?: string;
  active: boolean;
  order: number;
}

interface ApiResponse<T> {
  success: boolean;
  data: T;
  error?: string;
}

@Injectable({ providedIn: 'root' })
export class CourseSectionService {
  private readonly base = `${environment.apiUrl}/api/v1/course-sections`;

  constructor(private http: HttpClient) {}

  list(): Observable<CourseSection[]> {
    return this.http
      .get<ApiResponse<CourseSection[]>>(this.base)
      .pipe(map((res) => res.data ?? []));
  }

  getById(id: string): Observable<CourseSection> {
    return this.http
      .get<ApiResponse<CourseSection>>(`${this.base}/${id}`)
      .pipe(map((res) => res.data));
  }

  create(dto: CreateCourseSectionDto): Observable<CourseSection> {
    return this.http
      .post<ApiResponse<CourseSection>>(this.base, dto)
      .pipe(map((res) => res.data));
  }

  update(id: string, dto: UpdateCourseSectionDto): Observable<CourseSection> {
    return this.http
      .put<ApiResponse<CourseSection>>(`${this.base}/${id}`, dto)
      .pipe(map((res) => res.data));
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(`${this.base}/${id}`);
  }
}
