import { Component } from '@angular/core';
import { BackButtonComponent } from '../../shared/components/back-button/back-button';

@Component({
  selector: 'app-culinaria',
  standalone: true,
  imports: [BackButtonComponent],
  templateUrl: './culinaria.html',
  styleUrl: './culinaria.scss',
})
export class Culinaria {}
