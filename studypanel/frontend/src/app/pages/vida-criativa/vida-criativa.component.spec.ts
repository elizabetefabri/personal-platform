import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VidaCriativaComponent } from './vida-criativa.component';

describe('VidaCriativaComponent', () => {
  let component: VidaCriativaComponent;
  let fixture: ComponentFixture<VidaCriativaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [VidaCriativaComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(VidaCriativaComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
