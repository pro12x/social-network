import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {MatDrawer, MatDrawerContainer, MatSidenavModule} from "@angular/material/sidenav";
import {ToolbarComponent} from "./pages/nav/toolbar/toolbar.component";
import {SidenavComponent} from "./pages/nav/sidenav/sidenav.component";

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [
        RouterOutlet,
        MatSidenavModule,
        MatDrawerContainer,
        MatDrawer,
        ToolbarComponent,
        SidenavComponent
    ],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss'
})
export class AppComponent {
    title = 'frontend';
}
