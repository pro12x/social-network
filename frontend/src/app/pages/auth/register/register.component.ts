import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthService} from "../../../service/auth.service";
import {Router, RouterLink} from "@angular/router";

@Component({
    selector: 'app-register',
    standalone: true,
    imports: [
        ReactiveFormsModule,
        RouterLink
    ],
    templateUrl: './register.component.html',
    styleUrl: './register.component.scss'
})
export class RegisterComponent implements OnInit{
    title: string = 'Register';
    btnText: string = 'Register';
    age!: number
    avatar: File | null = null

    user: FormGroup = this.fb.group({
        firstname: [null, [Validators.required, Validators.minLength(3)]],
        lastname: [null, [Validators.required, Validators.minLength(3)]],
        email: [null, [Validators.required, Validators.email]],
        password: [null, [Validators.required, Validators.minLength(6)]],
        password_confirmation: [null, [Validators.required, Validators.minLength(6)]],
        date_of_birth: [null, [Validators.required]],
        avatar: [null],
        nickname: [null],
        about_me: [null],
    })

    constructor(
        private fb: FormBuilder,
        private authService: AuthService,
        private router: Router
    ) {}

    onSubmit() {
        const data = {
            ...this.user.value,
            avatar: 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png'
        }

        this.age = this.checkAge(data.date_of_birth)

        if (this.age < 12 || this.age > 120) {
            alert('You must be between 12 and 120 years old to register')
            return
        }

        if (data.password !== data.password_confirmation) {
            alert('Password and password confirmation do not match');
            return;
        }

        if (this.user.invalid) {
            alert('Please fill all the required fields');
            return;
        } else {
            this.authService.register(data).subscribe(() => {
                console.log("User registered")
                this.router.navigateByUrl('/login').then();
            }, (error) => {
                console.log(error)
            })
        }
    }

    checkAge(data: Date) : number {
        return Math.floor(Math.abs(Date.now() - new Date(data).getTime()) / (1000 * 3600 * 24 * 365))
    }

    ngOnInit(): void {
        if (this.authService.getToken()) {
            this.router.navigate(['/home']).then();
            console.log('You are already logged in');
            return;
        }

        console.log('Register component is running');
    }
}
