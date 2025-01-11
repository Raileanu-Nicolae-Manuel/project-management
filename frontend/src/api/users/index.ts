import type { User } from "@/types/users"
import { RequestBuilder, type Transaction } from "../request"

export default class UsersAPI {
    private readonly request: RequestBuilder
    constructor(private readonly path: string, private readonly headers: Headers) {
        this.request = new RequestBuilder(this.path+"/users", this.headers)
    }

    async getUsers(): Promise<User[]> {
        const {status, data, message} = await this.request.get<User[]>({})
        if (status < 200 || status >= 300) {
            console.warn(message)
            return [];
        }
        return data?? [];
    }
    async getUser(id: string): Promise<User|null> {
        const {status, data, message} = await this.request.get<User>({path: `/${id}`})
        if (status < 200 || status >= 300) {
            console.warn(message)
            return null;
        }
        return data;
    }

    async patchUsers(transactions: Transaction[]): Promise<User|null> {
        const {status, data, message} = await this.request.patch<User>({transactions})
        if (status < 200 || status >= 300) {
            console.warn(message)
            return null;
        }
        return data;
    }
}