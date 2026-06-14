import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { environment } from '../../../environments/environment';
import {
  StudyTableItem,
  CreateStudyItemDto,
  UpdateStudyItemDto,
} from '../../shared/interfaces/study-template.interface';

interface ApiResponse<T> {
  success: boolean;
  data: T;
  error?: string;
}

@Injectable({ providedIn: 'root' })
export class StudyItemService {
  private readonly base = `${environment.apiUrl}/api/v1/study-items`;

  constructor(private http: HttpClient) {}

  list(section: string, topic: string): Observable<StudyTableItem[]> {
    const params = new HttpParams().set('section', section).set('topic', topic);
    return this.http
      .get<ApiResponse<StudyTableItem[]>>(this.base, { params })
      .pipe(map((res) => res.data ?? []));
  }

  listAll(): Observable<StudyTableItem[]> {
    return this.http
      .get<ApiResponse<StudyTableItem[]>>(this.base)
      .pipe(map((res) => res.data ?? []));
  }

  getById(id: string): Observable<StudyTableItem> {
    return this.http
      .get<ApiResponse<StudyTableItem>>(`${this.base}/${id}`)
      .pipe(map((res) => res.data));
  }

  create(dto: CreateStudyItemDto): Observable<StudyTableItem> {
    return this.http
      .post<ApiResponse<StudyTableItem>>(this.base, dto)
      .pipe(map((res) => res.data));
  }

  update(id: string, dto: UpdateStudyItemDto): Observable<StudyTableItem> {
    return this.http
      .put<ApiResponse<StudyTableItem>>(`${this.base}/${id}`, dto)
      .pipe(map((res) => res.data));
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(`${this.base}/${id}`);
  }
}
