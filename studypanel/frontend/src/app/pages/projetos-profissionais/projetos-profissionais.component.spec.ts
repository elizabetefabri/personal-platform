import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { ProjetosProfissionaisComponent } from './projetos-profissionais.component';

describe('ProjetosProfissionaisComponent', () => {
  let component: ProjetosProfissionaisComponent;
  let fixture: ComponentFixture<ProjetosProfissionaisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProjetosProfissionaisComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(ProjetosProfissionaisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have at least one professional project', () => {
    expect(component.items.length).toBeGreaterThan(0);
  });

  it('should render a heading', () => {
    const h1 = fixture.nativeElement.querySelector('h1');
    expect(h1).toBeTruthy();
  });

  it('should have tags on all projects', () => {
    component.items.forEach((item) => {
      expect(Array.isArray(item.tags)).toBe(true);
    });
  });

  it('should include Rollout Service project', () => {
    const titles = component.items.map((i) => i.title);
    expect(titles).toContain('Rollout Service');
  });
});
