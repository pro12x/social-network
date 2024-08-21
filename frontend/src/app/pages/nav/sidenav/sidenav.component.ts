import {Component, OnInit} from '@angular/core';
import {RouterOutlet} from "@angular/router";
import {ErrorService} from "../../../service/error.service";
import {ToolbarComponent} from "../toolbar/toolbar.component";
import {NgIf} from "@angular/common";

@Component({
  selector: 'app-sidenav',
  standalone: true,
    imports: [
        RouterOutlet,
        ToolbarComponent,
        NgIf
    ],
  templateUrl: './sidenav.component.html',
  styleUrl: './sidenav.component.scss'
})
export class SidenavComponent implements OnInit{
    error!: boolean

    constructor(private errorService: ErrorService) {}

    ngOnInit(): void {
        this.errorService.error$.subscribe(error => {
            this.error = error
            console.log('Error:', this.error)
        })
    }
}
