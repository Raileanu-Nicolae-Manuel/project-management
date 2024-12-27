import type { DefaultType } from "./index"

export type UserBody = {
    name: string
    email: string
    password: string
}

export type UserAdditionalParams = {
    type: number
}

export type User = DefaultType & UserBody & UserAdditionalParams
export type UserPatchBody = Partial<UserBody & UserAdditionalParams>
