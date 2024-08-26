import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router, RouterLink} from "@angular/router";
import {User} from "../../../../entity/user";
import {AuthService} from "../../../service/auth.service";

@Component({
    selector: 'app-profile',
    standalone: true,
    imports: [
        RouterLink
    ],
    templateUrl: './profile.component.html',
    styleUrl: './profile.component.scss'
})
export class ProfileComponent implements OnInit {
    title: string = 'Profile'
    id!: number
    user!: User

    constructor(
        private authService: AuthService,
        private activatedRoute: ActivatedRoute,
        private router: Router
    ) {}

    getUser() {
        this.id = this.activatedRoute.snapshot.params['id']
        console.log(this.id)
    }

    ngOnInit(): void {
        console.log('Profile component initialized')
        this.getUser()
    }
}
