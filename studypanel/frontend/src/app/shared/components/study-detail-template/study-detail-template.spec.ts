import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StudyDetailTemplate } from './study-detail-template';

describe('StudyDetailTemplate', () => {
  let component: StudyDetailTemplate;
  let fixture: ComponentFixture<StudyDetailTemplate>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StudyDetailTemplate],
    }).compileComponents();

    fixture = TestBed.createComponent(StudyDetailTemplate);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
