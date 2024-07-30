import {User} from "./user";

export class Follow {
    id! : number
    followerID! : User
    followingID! : User
    State! : string
    createdAt! : string
    updatedAt! : string
}
