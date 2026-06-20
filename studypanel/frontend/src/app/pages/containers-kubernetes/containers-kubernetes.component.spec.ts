import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { of } from 'rxjs';
import { ContainersKubernetesComponent } from './containers-kubernetes.component';
import { StudyItemService } from '../../core/services/study-item.service';

const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

describe('ContainersKubernetesComponent', () => {
  let component: ContainersKubernetesComponent;
  let fixture: ComponentFixture<ContainersKubernetesComponent>;

  beforeEach(async () => {
    studyItemServiceMock.list.mockReturnValue(of([]));

    await TestBed.configureTestingModule({
      imports: [ContainersKubernetesComponent, RouterModule.forRoot([])],
      providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(ContainersKubernetesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
