import {Component, Injectable, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthService} from "../../../service/auth.service";
import {tap} from "rxjs";
import {Router, RouterLink} from "@angular/router";
import {MatSnackBar} from "@angular/material/snack-bar";
import {UtilsService} from "../../../service/utils.service";

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [
        ReactiveFormsModule,
        RouterLink
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss'
})

export class LoginComponent implements OnInit {
    title = 'Login';
    btnText = 'Login';

    user: FormGroup = this.fb.group({
        email: [null, [Validators.required, Validators.email]],
        password: [null, [Validators.required, Validators.minLength(6)]]
    })

    constructor(
        private fb: FormBuilder,
        private authService: AuthService,
        private router: Router,
        private utilsService: UtilsService
    ) {
    }

    onLogin() {
        if (this.user.invalid) {
            this.utilsService.onSnackBar('Please fill in the form correctly', 'warning');
            return;
        }

        this.login(this.user.value).subscribe((response: any) => {
            this.utilsService.onSnackBar('You are now logged in', 'success');
            this.router.navigate(['/home']).then();
        })
    }

    login(credentials: { email: string, password: string }) {
        return this.authService.login(credentials).pipe(
            tap((res: any) => {
                if (!res.status || res.status !== 200) {
                    this.utilsService.onSnackBar(res.message, 'error');
                    return;
                }
                this.authService.createSession(res.token, res.user.id);
            })
        )
    }

    ngOnInit(): void {
        this.utilsService.setTitle(this.title);
        if (this.authService.getToken() && this.authService.getUserID()) {
            this.utilsService.onSnackBar('You are already logged in', 'info');
            this.router.navigate(['/home']).then();
            return;
        }
    }
}
