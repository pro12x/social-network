import {Component, OnInit} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {RegisterComponent} from "./pages/auth/register/register.component";
import {UtilsService} from "./service/utils.service";
import {Title} from "@angular/platform-browser";
import {ToolbarComponent} from "./pages/nav/toolbar/toolbar.component";
import {AuthService} from "./service/auth.service";

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [
        RouterOutlet,
        RegisterComponent,
        ToolbarComponent,
    ],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
    error!: boolean

    constructor(
        private utilsService: UtilsService,
        private title: Title,
    ) {}

    ngOnInit() {
        this.utilsService.getTitle().subscribe((title: string) => {
            this.title.setTitle(title)
        })
    }
}
