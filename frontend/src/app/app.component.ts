import {Component, OnInit} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {RegisterComponent} from "./pages/auth/register/register.component";

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [
        RouterOutlet,
        RegisterComponent,
    ],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
    title = 'frontend';
    error!: boolean

    constructor() {
    }

    ngOnInit() {
        console.log('Your app is running');
    }
}
