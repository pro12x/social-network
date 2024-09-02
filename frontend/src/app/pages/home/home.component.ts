import {Component, OnInit} from '@angular/core';
import {Router, RouterLink} from "@angular/router";
import {AuthService} from "../../service/auth.service";
import {User} from "../../../entity/user";
import {NgForOf, NgIf} from "@angular/common";
import {MatIcon} from "@angular/material/icon";
import {FollowService} from "../../service/follow.service";
import {UtilsService} from "../../service/utils.service";
import {ToolbarComponent} from "../nav/toolbar/toolbar.component";

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
    userID: number | null = this.authService.getUserID()
    isFollowing: boolean = false

    constructor(
        private authService: AuthService,
        private followService: FollowService,
        private toolbar: ToolbarComponent,
        private router: Router,
        private utilsService: UtilsService
    ) {}

    usersList() {
        this.authService.getAll().subscribe((response: any) => {
            if (response.status === 'empty') {
                console.log(response.message)
                return
            } else {
                this.users = response.users.filter((user: User) => user.id !== this.authService.getUserID())
                this.users.forEach((user: User) => {
                    this.existFollow(user.id)
                })
            }
        })
    }

    existFollow(id: number) {
        const data = {
            follower_id: this.authService.getUserID(),
            followee_id: id
        }

        this.followService.checkFollow(data, "following").subscribe((response) => {
            if (response.status !== 200) {
                this.utilsService.onSnackBar(response.message, 'error')
                return
            }
            this.isFollowing = response.is_following
        })
    }

    onFollow(id: number) {
        const data = {
            follower_id: this.authService.getUserID(),
            followee_id: id
        }

        this.followService.follow(data, "follow").subscribe((response) => {
            if (response.status !== 200) {
                this.utilsService.onSnackBar(response.message, 'error')
                return
            }
            this.usersList()
        })
    }

    ngOnInit(): void {
        this.utilsService.setTitle(this.title)
        if (!this.authService.getToken()) {
            this.router.navigate(['/login']).then()
            return
        }
        this.toolbar.isOnline()

        this.toolbar.getUser(this.userID!)
        this.usersList()
    }
}
