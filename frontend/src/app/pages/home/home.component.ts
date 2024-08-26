import {Component, OnInit} from '@angular/core';
import {Router, RouterLink} from "@angular/router";
import {AuthService} from "../../service/auth.service";
import {User} from "../../../entity/user";
import {NgForOf, NgIf} from "@angular/common";
import {FormBuilder} from "@angular/forms";
import {MatIcon} from "@angular/material/icon";

@Component({
    selector: 'app-home',
    standalone: true,
    imports: [
        RouterLink,
        NgForOf,
        NgIf,
        MatIcon
    ],
    templateUrl: './home.component.html',
    styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit {
    title: string = 'Home'
    users!: User[]
    token!: string
    activeUser: any = {}
    userID: number | null = this.authService.getUserID()

    constructor(
        private authService: AuthService,
        private router: Router,
        private fb: FormBuilder
    ) {}

    onLoggout() {
        const data = {
            token: this.token
        }
        this.authService.logout(data).subscribe((response) => {
            if (!response.status || response.status !== 'success') {
                console.log("Error logging out")
                return
            }

            console.log(response)
            localStorage.removeItem('token')
            localStorage.removeItem('userID')
            console.log('Logged out')
            this.router.navigate(['/login']).then()
        })
    }

    isOnline() {
        this.authService.isLoggedIn().subscribe(response => {
            if (response) {
                console.log('You are online')
                return
            } else {
                console.log('You are offline')
                localStorage.removeItem('token')
                localStorage.removeItem('userID')
                this.router.navigate(['/login']).then()
            }
        })
    }

    usersList() {
        this.authService.getAll().subscribe((response: any) => {
            if (response.status === 'empty') {
                console.log(response.message)
                return
            } else {
                this.users = response.users.filter((user: User) => user.id !== this.authService.getUserID())
            }
        })
    }

    getUser(id: number): any {
        return this.authService.getUser(id).subscribe((response: any) => {
            this.activeUser = response.user
        })
    }

    ngOnInit(): void {
        if (!this.authService.getToken()) {
            this.router.navigate(['/login']).then()
            console.log('You are not logged in')
            return
        }

        this.isOnline()

        this.getUser(this.authService.getUserID()!)

        this.token = this.authService.getToken()!
        this.usersList()
        console.log('Home component is running')
    }
}
