import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { of } from 'rxjs';
import { StudyDetailTemplate } from './study-detail-template';
import { StudyItemService } from '../../../core/services/study-item.service';

const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

describe('StudyDetailTemplate', () => {
  let component: StudyDetailTemplate;
  let fixture: ComponentFixture<StudyDetailTemplate>;

  beforeEach(async () => {
    studyItemServiceMock.list.mockReturnValue(of([]));

    await TestBed.configureTestingModule({
      imports: [StudyDetailTemplate, RouterModule.forRoot([])],
      providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(StudyDetailTemplate);
    component = fixture.componentInstance;
    component.pageTitle = 'Backend';
    component.pageDescription = 'Recursos de backend';
    component.section = 'backend';
    component.topic = 'node';
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render page title', () => {
    const h1 = fixture.nativeElement.querySelector('.detailHeader__title');
    expect(h1?.textContent).toContain('Backend');
  });

  it('should render the eyebrow text', () => {
    const eyebrow = fixture.nativeElement.querySelector('.detailHeader__eyebrow');
    expect(eyebrow?.textContent?.trim()).toBe('Detalhes do estudo');
  });

  it('should call loadItems on init', () => {
    expect(studyItemServiceMock.list).toHaveBeenCalledWith('backend', 'node');
  });

  it('should start with empty items', () => {
    expect(component.items()).toEqual([]);
  });

  it('should open create modal with clean form', () => {
    component.openCreateModal();
    expect(component.modalVisible).toBe(true);
    expect(component.editingItem).toBeNull();
    expect(component.form.courseName).toBe('');
  });

  it('should close modal on closeModal()', () => {
    component.modalVisible = true;
    component.closeModal();
    expect(component.modalVisible).toBe(false);
  });

  it('should map status to correct severity', () => {
    expect(component.getStatusSeverity('Concluído')).toBe('success');
    expect(component.getStatusSeverity('Em andamento')).toBe('info');
    expect(component.getStatusSeverity('Não iniciado')).toBe('warn');
    expect(component.getStatusSeverity('Pausado')).toBe('danger');
  });

  it('should build correct course detail route', () => {
    const item = { id: '42' } as any;
    expect(component.courseDetailRoute(item)).toBe('/backend/node/42');
  });
});
