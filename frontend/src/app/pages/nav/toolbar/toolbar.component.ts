import {Component, Injectable, OnInit} from '@angular/core';
import {MatToolbar} from "@angular/material/toolbar";
import {UtilsService} from "../../../service/utils.service";
import {NgIf} from "@angular/common";
import {Router, RouterLink} from "@angular/router";
import {AuthService} from "../../../service/auth.service";

@Component({
  selector: 'app-toolbar',
  standalone: true,
    imports: [
        MatToolbar,
        NgIf,
        RouterLink
    ],
  templateUrl: './toolbar.component.html',
  styleUrl: './toolbar.component.scss'
})
@Injectable({
    providedIn: "root"
})
export class ToolbarComponent implements OnInit {
    page!: string
    token!: string
    activeUser: any = {}
    userID: number | null = this.authService.getUserID()

    constructor(
        private utilsService: UtilsService,
        private authService: AuthService,
        private router: Router
    ) {}

    getPage() {
        this.utilsService.getTitle().source?.subscribe(res => {
            this.page = res
        })
    }

    getUser(id: number): any {
        if (!id) {
            return
        }
        return this.authService.getUser(id).subscribe((response: any) => {
            this.activeUser = response.user
        })
    }

    onLogout() {
        const data = {
            token: this.token
        }
        this.authService.logout(data).subscribe((response) => {
            if (!response.status || response.status !== 200) {
                this.utilsService.onSnackBar(response.message, 'error')
                return
            }
            this.utilsService.onSnackBar('You are now logged out', 'primary')
            this.authService.removeSession()
            this.router.navigate(['/login']).then()
        })
    }

    isOnline() {
        this.authService.isLoggedIn().subscribe(response => {
            if (response) {
                return
            } else {
                this.authService.removeSession()
                this.router.navigate(['/login']).then()
            }
        })
    }

    ngOnInit(): void {
        this.isOnline()
        this.getUser(this.userID!)
        this.token = this.authService.getToken()!

        this.getPage()
    }
}
