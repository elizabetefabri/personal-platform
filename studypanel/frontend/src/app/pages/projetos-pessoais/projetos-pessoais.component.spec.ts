import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { ProjetosPessoaisComponent } from './projetos-pessoais.component';

describe('ProjetosPessoaisComponent', () => {
  let component: ProjetosPessoaisComponent;
  let fixture: ComponentFixture<ProjetosPessoaisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProjetosPessoaisComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(ProjetosPessoaisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have at least one personal project', () => {
    expect(component.items.length).toBeGreaterThan(0);
  });

  it('should render a heading', () => {
    const h1 = fixture.nativeElement.querySelector('h1');
    expect(h1).toBeTruthy();
  });

  it('should have tags on all projects', () => {
    component.items.forEach((item) => {
      expect(Array.isArray(item.tags)).toBe(true);
      expect(item.tags.length).toBeGreaterThan(0);
    });
  });

  it('should have repoUrl or deployUrl or detailRoute on each project', () => {
    component.items.forEach((item) => {
      const hasLink = !!(item.repoUrl || item.deployUrl || item.detailRoute);
      expect(hasLink).toBe(true);
    });
  });
});
