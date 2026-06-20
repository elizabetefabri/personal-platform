import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { ProjetosComponent } from './projetos.component';

describe('ProjetosComponent', () => {
  let component: ProjetosComponent;
  let fixture: ComponentFixture<ProjetosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProjetosComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(ProjetosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have 2 landing cards', () => {
    expect(component.items.length).toBe(2);
  });

  it('should include Projetos Pessoais card', () => {
    const titles = component.items.map((i) => i.title);
    expect(titles).toContain('Projetos Pessoais');
  });

  it('should include Projetos Profissionais card', () => {
    const titles = component.items.map((i) => i.title);
    expect(titles).toContain('Projetos Profissionais');
  });

  it('should have detailRoute pointing to /projetos/pessoais', () => {
    const pessoais = component.items.find((i) => i.title === 'Projetos Pessoais');
    expect(pessoais?.detailRoute).toBe('/projetos/pessoais');
  });

  it('should have detailRoute pointing to /projetos/profissionais', () => {
    const prof = component.items.find((i) => i.title === 'Projetos Profissionais');
    expect(prof?.detailRoute).toBe('/projetos/profissionais');
  });

  it('should render projetos heading', () => {
    const h1 = fixture.nativeElement.querySelector('h1');
    expect(h1?.textContent).toContain('Projetos');
  });
});
