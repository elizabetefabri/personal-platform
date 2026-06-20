import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { of } from 'rxjs';
import { DashboardComponent } from './dashboard.component';
import { StudyItemService } from '../../core/services/study-item.service';

const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  listAll: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

describe('DashboardComponent', () => {
  let component: DashboardComponent;
  let fixture: ComponentFixture<DashboardComponent>;

  beforeEach(async () => {
    studyItemServiceMock.listAll.mockReturnValue(of([]));

    await TestBed.configureTestingModule({
      imports: [DashboardComponent, RouterModule.forRoot([])],
      providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(DashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render study sections heading', () => {
    const h2 = fixture.nativeElement.querySelector('h2');
    expect(h2?.textContent).toContain('Seções de estudo');
  });
});
