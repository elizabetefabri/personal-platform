import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Culinaria } from './culinaria';

describe('Culinaria', () => {
  let component: Culinaria;
  let fixture: ComponentFixture<Culinaria>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Culinaria],
    }).compileComponents();

    fixture = TestBed.createComponent(Culinaria);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
