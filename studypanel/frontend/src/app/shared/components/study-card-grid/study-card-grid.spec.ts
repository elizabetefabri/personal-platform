import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StudyCardGrid } from './study-card-grid';

describe('StudyCardGrid', () => {
  let component: StudyCardGrid;
  let fixture: ComponentFixture<StudyCardGrid>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StudyCardGrid],
    }).compileComponents();

    fixture = TestBed.createComponent(StudyCardGrid);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
