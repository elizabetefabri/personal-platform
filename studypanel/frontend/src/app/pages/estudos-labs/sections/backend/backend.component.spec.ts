import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { of } from 'rxjs';
import { BackendComponent } from './backend.component';
import { StudyItemService } from '../../core/services/study-item.service';

const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

describe('BackendComponent', () => {
  let component: BackendComponent;
  let fixture: ComponentFixture<BackendComponent>;

  beforeEach(async () => {
    studyItemServiceMock.list.mockReturnValue(of([]));

    await TestBed.configureTestingModule({
      imports: [BackendComponent, RouterModule.forRoot([])],
      providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(BackendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
