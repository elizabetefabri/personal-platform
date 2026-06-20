import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { of } from 'rxjs';
import { InteligenciaArtificialComponent } from './inteligencia-artificial.component';
import { StudyItemService } from '../../core/services/study-item.service';

const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

describe('InteligenciaArtificialComponent', () => {
  let component: InteligenciaArtificialComponent;
  let fixture: ComponentFixture<InteligenciaArtificialComponent>;

  beforeEach(async () => {
    studyItemServiceMock.list.mockReturnValue(of([]));

    await TestBed.configureTestingModule({
      imports: [InteligenciaArtificialComponent, RouterModule.forRoot([])],
      providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(InteligenciaArtificialComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
