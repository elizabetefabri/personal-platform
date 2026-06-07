import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { ToggleSwitchModule } from 'primeng/toggleswitch';

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [FormsModule, ButtonModule, InputTextModule, ToggleSwitchModule],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.scss',
})
export class SettingsComponent {
  profile = {
    name: '',
    email: '',
  };

  preferences = {
    darkMode: false,
    notifications: true,
  };

  saveProfile(): void {
    console.log('Perfil salvo:', this.profile);
  }
}
