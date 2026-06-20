import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { EstudosLabsComponent } from './estudos-labs.component';

describe('EstudosLabsComponent', () => {
  let component: EstudosLabsComponent;
  let fixture: ComponentFixture<EstudosLabsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EstudosLabsComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(EstudosLabsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have 9 study card items', () => {
    expect(component.items.length).toBe(9);
  });

  it('should include backend card', () => {
    const slugs = component.items.map((i) => i.detailRoute);
    expect(slugs).toContain('/backend');
  });

  it('should include devops card', () => {
    const slugs = component.items.map((i) => i.detailRoute);
    expect(slugs).toContain('/devops');
  });

  it('should have detailRoute set for all items', () => {
    component.items.forEach((item) => {
      expect(item.detailRoute).toBeTruthy();
    });
  });

  it('should render page heading', () => {
    const h1 = fixture.nativeElement.querySelector('h1');
    expect(h1).toBeTruthy();
  });
});
