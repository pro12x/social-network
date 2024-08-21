import {Component, OnInit} from '@angular/core';
import {ToolbarComponent} from "../nav/toolbar/toolbar.component";
import {ErrorService} from "../../service/error.service";
import {NgIf} from "@angular/common";

@Component({
    selector: 'app-home',
    standalone: true,
    imports: [
        ToolbarComponent,
        NgIf
    ],
    templateUrl: './home.component.html',
    styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit {
    error!: boolean

    constructor(private errorService: ErrorService) {}

    ngOnInit() {
        this.errorService.error$.subscribe(error => {
            this.error = error
            console.log('Error:', this.error)
        })
    }
}
