import ENV from "@/service/env"

export type RequestBuilderParams = {
    path: string
    params: { [key: string]: string|number|boolean }
    headers: Headers
}

export type RequestBuilderBody<T> = {
    body: T
}

export type Transaction = {
    path: string
    op: 'remove' | 'add' | 'replace' | 'move' | 'copy' | 'test'
    value?: any
    from?: string
}

export type RequestBuilderPatch = {
    transactions: Transaction[]
}

export class RequestBuilder {
    private readonly baseUrl: string
    private readonly headers: Headers
    constructor(baseUrl: string, headers: Headers) {
        this.baseUrl = baseUrl
        this.headers = headers
    }

    async get<T>(params: Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithOutBody<T>({...params, method: 'GET'})
    }

    async post<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithBody<T>({...params, method: 'POST'})
    }

    async put<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithBody<T>({...params, method: 'PUT'})
    }

    async delete(params: Partial<RequestBuilderParams>): Promise<{status: number, data: null, message: string}> {
        return this.requestWithOutBody<null>({...params, method: 'DELETE'})
    }

    async patch<T>(params: RequestBuilderPatch & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        if (ENV.DEV_MODE) {
            console.log(`url: ${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, `method: PATCH`, `Transactions: ${params.transactions}`);
        }
        const response = await fetch(`${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, {
            method: "PATCH",
            body: JSON.stringify(params.transactions),
            headers: computedHeaders(this.headers, params.headers??new Headers())
        })
        if (!response.ok) {
            return {status: response.status, data: null, message: await response.text()}
        }
        return {status: response.status, data: await response.json() as T, message: ""}
    }

    private async requestWithBody<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams> & {method: string}) {
        if (ENV.DEV_MODE) {
            console.log(`url: ${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, `method: ${params.method}`, `body: ${params.body}`);
        }
        const response = await fetch(`${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, {
            method: params.method,
            body: JSON.stringify(params.body),
            headers: computedHeaders(this.headers, params.headers??new Headers())
        })
        if (!response.ok) {
            return {status: response.status, data: null, message: await response.text()}
        }
        return {status: response.status, data: await response.json() as T, message: ""}
    }

    private async requestWithOutBody<T>(params: Partial<RequestBuilderParams> & {method: string}) {
        if (ENV.DEV_MODE) {
            console.log(`url: ${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, `method: ${params.method}`);
        }
        const response = await fetch(`${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, {
            method: params.method,
            headers: computedHeaders(this.headers, params.headers??new Headers())
        });
        if (!response.ok) {
            return {status: response.status, data: null, message: await response.text()}
        }
        return {status: response.status, data: await response.json() as T, message: ""}
    }
}

const computedPath = (path: string, params: { [key: string]: string|number|boolean }) => {
    return `${path}${Object.entries(params).length > 0 ? 
        `?${Object.entries(params).map(([key, value]) => `${key}=${value}`).join('&')}` : ''}`
}

const computedHeaders = (RootHeaders: Headers, headers: Headers) => {
    return new Headers([...RootHeaders, ...headers, ...new Headers({credentials: 'include'})]);
}
