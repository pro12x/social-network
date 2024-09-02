import {Routes} from '@angular/router';
import {LoginComponent} from "./pages/auth/login/login.component";
import {HomeComponent} from "./pages/home/home.component";
import {RegisterComponent} from "./pages/auth/register/register.component";
import {ProfileComponent} from "./pages/auth/profile/profile.component";

export const routes: Routes = [
    {path: '', component: HomeComponent, data: {title: 'Home'}},
    {path: 'login', component: LoginComponent, data: {title: 'Login'}},
    {path: 'register', component: RegisterComponent, data: {title: 'Register'}},
    {path: 'profile/:id', component: ProfileComponent, data: {title: 'Profile'}},
    //{path: 'error', component: ErrorComponent, data: {title: 'Error'}},
    {path: '**', redirectTo: '', pathMatch: 'full'},
];
