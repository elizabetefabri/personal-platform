import { Component } from '@angular/core';
import { BackButtonComponent } from '../../shared/components/back-button/back-button';

@Component({
  selector: 'app-painel-financeiro',
  standalone: true,
  imports: [BackButtonComponent],
  templateUrl: './painel-financeiro.component.html',
  styleUrl: './painel-financeiro.component.scss',
})
export class PainelFinanceiroComponent {}
