import {Component} from '@angular/core';
import {MatSidenav, MatSidenavContainer} from "@angular/material/sidenav";
import {MatListModule} from "@angular/material/list";
import {RouterLink} from "@angular/router";
import {MatCardModule} from "@angular/material/card";
import {MatIconModule} from "@angular/material/icon";
import {MatButton, MatIconButton} from "@angular/material/button";
import {NgForOf, NgOptimizedImage} from "@angular/common";
import {ReactiveFormsModule} from "@angular/forms";
import {MatFormField} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";

@Component({
    selector: 'app-home',
    standalone: true,
    imports: [
        MatSidenavContainer,
        MatSidenav,
        MatListModule,
        RouterLink,
        MatCardModule,
        MatIconModule,
        MatButton,
        NgOptimizedImage,
        NgForOf,
        MatIconButton,
        ReactiveFormsModule,
        MatFormField,
        MatInput
    ],
    templateUrl: './home.component.html',
    styleUrl: './home.component.scss'
})
export class HomeComponent {
    id: number = 1

    users = [
        {id: 1, name: 'Janel Proverbes', username: 'proverbes'},
        {id: 2, name: 'John Doe', username: 'doe'},
        {id: 3, name: 'Jane Doe', username: 'jane'},
    ]

    posts = [
        {
            id: 1,
            userID: 1,
            title: 'My first post',
            content: 'Automobile company Wanderer was originally established in 1885, later becoming a branch of Audi AG. Another company,',
            categories: ['first', 'post'],
            image: 'posts/post001.jpg',
            comments: [
                {id: 1, owner: 'Janel Proverbes', content: 'Nice post!'},
                {id: 2, owner: 'John Doe', content: 'Great post!'},
            ],
            likes: 2,
            dislikes: 0,
            shares: 1
        },
        {
            id: 2,
            userID: 2,
            title: 'My second post',
            content: 'NSU, which also later merged into Audi, was founded during this time, and later supplied the chassis for Gottlieb Daimler\'s four-wheeler.',
            categories: ['second', 'post'],
            image: 'posts/post002.jpg',
            comments: [
                {id: 1, owner: 'Janel Proverbes', content: 'Nice post!'},
            ],
            likes: 1,
            dislikes: 1,
            shares: 0
        },
        {
            id: 3,
            userID: 1,
            title: 'My third post',
            content: 'The next major model change came in 1995 when the Audi A4 replaced the Audi 80. The new nomenclature scheme was applied to the Audi 100 to become the Audi A6 (with a minor facelift).',
            categories: ['third', 'post'],
            image: 'posts/post003.jpg',
            comments: [
                {id: 1, owner: 'Janel Proverbes', content: 'Nice post!'},
            ],
            likes: 1,
            dislikes: 0,
            shares: 0
        },
        {
            id: 4,
            userID: 3,
            title: 'My fourth post',
            content: 'The Shiba Inu is the smallest of the six original and distinct spitz breeds of dog from Japan. A small, agile dog that copes very well with mountainous terrain, the Shiba Inu was originally bred for hunting.',
            categories: ['fourth', 'post'],
            image: 'posts/post004.jpg',
            comments: [
                {id: 1, owner: 'Janel Proverbes', content: 'Nice post!'},
                {id: 2, owner: 'John Doe', content: 'Great post!'},
                {id: 3, owner: 'Jane Doe', content: 'Awesome post!'},
                {id: 4, owner: 'Janel Proverbes', content: 'Nice post!'},
            ],
            likes: 4,
            dislikes: 0,
            shares: 12
        }
    ]
}
