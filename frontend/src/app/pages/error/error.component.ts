import {Component, Injectable, OnInit} from '@angular/core';
import {RouterLink} from "@angular/router";
import {MatIconModule} from "@angular/material/icon";
import {MatMiniFabButton} from "@angular/material/button";
import {ErrorService} from "../../service/error.service";

@Component({
    selector: 'app-error',
    standalone: true,
    imports: [
        RouterLink,
        MatIconModule,
        MatMiniFabButton
    ],
    templateUrl: './error.component.html',
    styleUrl: './error.component.scss'
})
@Injectable({
    providedIn: 'root'
})
export class ErrorComponent implements OnInit {
    error!: boolean
    title = 'Opps!'
    message = 'Sorry, an error occurred.'

    constructor(private errorService: ErrorService) {}

    ngOnInit(): void {
        this.errorService.error$.subscribe(() => {
            this.error = !this.error
            console.log('Error:', this.error)
        })
    }
}
