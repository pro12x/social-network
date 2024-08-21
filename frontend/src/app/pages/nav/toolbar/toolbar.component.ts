import {Component} from '@angular/core';
import {NgIf} from "@angular/common";
import {RouterLink} from "@angular/router";
import {ErrorService} from "../../../service/error.service";

@Component({
    selector: 'app-toolbar',
    standalone: true,
    imports: [
        NgIf,
        RouterLink
    ],
    templateUrl: './toolbar.component.html',
    styleUrl: './toolbar.component.scss'
})
export class ToolbarComponent {
    error!: boolean

    constructor(private errorService: ErrorService) {}

    ngOnInit(): void {
        this.errorService.error$.subscribe(error => {
            this.error = error
            console.log('Error:', this.error)
        })
    }
}
