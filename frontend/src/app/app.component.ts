import {Component, OnInit} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {NgIf} from "@angular/common";
import {ToolbarComponent} from "./pages/nav/toolbar/toolbar.component";
import {HomeComponent} from "./pages/home/home.component";
import {ErrorService} from "./service/error.service";
import {SidenavComponent} from "./pages/nav/sidenav/sidenav.component";

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [
        RouterOutlet,
        NgIf,
        ToolbarComponent,
        HomeComponent,
        SidenavComponent
    ],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
    title = 'frontend';
    error!: boolean

    constructor(private errorService: ErrorService) {}

    ngOnInit() {
        this.errorService.error$.subscribe(error => {
            this.error = error
            console.log('Error:', this.error)
        })
    }
}
