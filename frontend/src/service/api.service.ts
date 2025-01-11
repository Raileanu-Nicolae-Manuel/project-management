import ENV from "./env";
import UsersAPI from "@/api/users";

export default class ApiService {
    private baseUrl: string;
    users: UsersAPI;
    constructor() {
        this.baseUrl = ENV.API_URL;
        this.users = new UsersAPI(this.baseUrl, new Headers({
            "Content-Type": "application/json"
        }));
    }
}