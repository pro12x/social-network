import {Component} from '@angular/core';
import {
    MatDrawer,
    MatDrawerContainer,
    MatDrawerContent,
    MatSidenav,
    MatSidenavContainer
} from "@angular/material/sidenav";
import {RouterLink, RouterOutlet} from "@angular/router";
import {MatListModule} from "@angular/material/list";
import {MatIcon} from "@angular/material/icon";
import {NgForOf, NgIf} from "@angular/common";
import {MatFabAnchor} from "@angular/material/button";

@Component({
    selector: 'app-sidenav',
    standalone: true,
    imports: [
        MatDrawer,
        MatDrawerContainer,
        MatDrawerContent,
        RouterOutlet,
        MatListModule,
        MatIcon,
        NgForOf,
        MatFabAnchor,
        RouterLink,
        NgIf,
        MatSidenavContainer,
        MatSidenav
    ],
    templateUrl: './sidenav.component.html',
    styleUrl: './sidenav.component.scss'
})
export class SidenavComponent {
    menuItems = [
        {name: 'Home', route: '/', icon: 'home'},
        {name: 'Profile', route: '/profile', icon: 'person'},
        {name: 'Friends', route: '/followers', icon: 'person_add'},
        {name: 'Groups', route: '/groups', icon: 'group'},
    ]
}
