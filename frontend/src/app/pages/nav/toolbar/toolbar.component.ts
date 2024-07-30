import {Component} from '@angular/core';
import {MatToolbar} from "@angular/material/toolbar";
import {MatIcon} from "@angular/material/icon";
import {MatIconButton} from "@angular/material/button";
import {MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {RouterLink} from "@angular/router";
import {MatBadge} from "@angular/material/badge";
import {MatMenu} from "@angular/material/menu";
import {MatCardAvatar} from "@angular/material/card";

@Component({
    selector: 'app-toolbar',
    standalone: true,
    imports: [
        MatToolbar,
        MatIcon,
        MatIconButton,
        MatFormField,
        MatInput,
        MatLabel,
        RouterLink,
        MatBadge,
        MatMenu,
        MatCardAvatar
    ],
    templateUrl: './toolbar.component.html',
    styleUrl: './toolbar.component.scss'
})
export class ToolbarComponent {
    title = 'Social Network';
    username = 'Janel Proverbes';
    hiddenNotif = false;
    hiddenMessage = false;

    messages = [
        {
            username: 'Janel Proverbes',
            message: 'Hey, how are you?'
        },
        {
            username: 'Janel Proverbes',
            message: 'Hey, how are you?'
        },
        {
            username: 'Janel Proverbes',
            message: 'Hey, how are you?'
        },
        {
            username: 'Janel Proverbes',
            message: 'Hey, how are you?'
        }
    ]

    visibilityNotif() {
        this.hiddenNotif = !this.hiddenNotif;
    }

    visibilityMessage() {
        this.hiddenMessage = !this.hiddenMessage;
    }
}
