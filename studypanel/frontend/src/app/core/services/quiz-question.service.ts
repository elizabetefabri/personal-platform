import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { environment } from '../../../environments/environment';
import { QuizQuestion } from '../../shared/interfaces/study-template.interface';

interface ApiResponse<T> {
  success: boolean;
  data: T;
}

@Injectable({ providedIn: 'root' })
export class QuizQuestionService {
  private readonly base = `${environment.apiUrl}/api/v1/quiz-questions`;

  constructor(private http: HttpClient) {}

  list(section: string, topic: string): Observable<QuizQuestion[]> {
    const params = new HttpParams().set('section', section).set('topic', topic);
    return this.http
      .get<ApiResponse<QuizQuestion[]>>(this.base, { params })
      .pipe(map((r) => r.data ?? []));
  }
}
