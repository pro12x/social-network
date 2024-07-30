import {Component} from '@angular/core';
import {MatFabAnchor, MatMiniFabButton} from "@angular/material/button";
import {RouterLink} from "@angular/router";
import {MatIconModule} from "@angular/material/icon";
import {NgOptimizedImage} from "@angular/common";

@Component({
    selector: 'app-error',
    standalone: true,
    imports: [
        MatIconModule,
        MatFabAnchor,
        RouterLink,
        MatMiniFabButton,
        NgOptimizedImage
    ],
    templateUrl: './error.component.html',
    styleUrl: './error.component.scss'
})
export class ErrorComponent {
    error = true
}
